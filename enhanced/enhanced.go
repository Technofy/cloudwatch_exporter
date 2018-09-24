package enhanced

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
	"sync"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/client"
	"github.com/aws/aws-sdk-go/service/cloudwatchlogs"
	"github.com/aws/aws-sdk-go/service/rds"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/common/log"

	"github.com/percona/rds_exporter/config"
	"github.com/percona/rds_exporter/enhanced/metrics"
	"github.com/percona/rds_exporter/latency"
	"github.com/percona/rds_exporter/sessions"
)

const (
	defaultNamespace = "rdsosmetrics"
	logGroupName     = "RDSOSMetrics"
)

var (
	ScrapeTimeDesc = prometheus.NewDesc(
		"rds_exporter_scrape_duration_seconds",
		"Time this RDS scrape took, in seconds.",
		[]string{},
		nil,
	)
)

type Exporter struct {
	Config   *config.Config
	Sessions *sessions.Sessions

	// Metrics
	ErroneousRequests prometheus.Counter
	TotalRequests     prometheus.Counter

	l log.Logger
}

func New(config *config.Config, sessions *sessions.Sessions) *Exporter {
	return &Exporter{
		Config:   config,
		Sessions: sessions,

		ErroneousRequests: prometheus.NewCounter(prometheus.CounterOpts{
			Name: "rds_exporter_erroneous_requests",
			Help: "The number of erroneous API request made to CloudWatch.",
		}),
		TotalRequests: prometheus.NewCounter(prometheus.CounterOpts{
			Name: "rds_exporter_requests_total",
			Help: "API requests made to CloudWatch",
		}),

		l: log.With("component", "enhanced"),
	}
}

func (e *Exporter) Describe(ch chan<- *prometheus.Desc) {
	// RDS metrics
	for _, m := range metrics.Default {
		ch <- m.Desc
	}

	// Latency metric
	ch <- latency.Desc

	// Scrape time
	ch <- ScrapeTimeDesc

	// Global number of requests, and global number of failed requests
	ch <- e.TotalRequests.Desc()
	ch <- e.ErroneousRequests.Desc()
}

func (e *Exporter) Collect(ch chan<- prometheus.Metric) {
	now := time.Now()
	e.collect(ch)

	// Collect scrape time
	ch <- prometheus.MustNewConstMetric(ScrapeTimeDesc, prometheus.GaugeValue, float64(time.Since(now).Seconds()))

	// Collect global number of requests, and global number of failed requests
	ch <- e.TotalRequests
	ch <- e.ErroneousRequests
}

func (e *Exporter) collect(ch chan<- prometheus.Metric) {
	wg := &sync.WaitGroup{}
	defer wg.Wait()

	instances := e.Config.Instances
	wg.Add(len(instances))
	for _, instance := range instances {
		instance := instance
		go func() {
			defer wg.Done()

			err := e.collectInstance(ch, &instance)
			if err != nil {
				e.l.With("region", instance.Region).With("instance", instance.Instance).Error(err)
			}
		}()
	}
}

func (e *Exporter) collectInstance(ch chan<- prometheus.Metric, instance *config.Instance) error {
	// Latency metric
	l := &latency.Latency{}

	// Collect all values
	err := e.collectValues(ch, instance, l)

	// Collect latency metric
	if !l.IsZero() {
		ch <- prometheus.MustNewConstMetric(latency.Desc, prometheus.GaugeValue, float64(l.Duration().Seconds()), instance.Instance, instance.Region)
	}
	return err
}

func (e *Exporter) collectValues(ch chan<- prometheus.Metric, instance *config.Instance, l *latency.Latency) error {
	sess := e.Sessions.GetSession(instance.Region, instance.Instance)

	logStreamName, err := dbiResourceId(sess, instance.Instance)
	if err != nil {
		return err
	}

	values, err := events(sess, logStreamName, instance.Instance, time.Minute)
	if err != nil {
		return err
	}
	var errs errs

	mv := metrics.Metric{
		Namespace:    defaultNamespace,
		Subsystem:    "General",
		Labels:       []string{"instance", "region"},
		LabelsValues: []string{instance.Instance, instance.Region},
	}

	for name, value := range values {
		mv.Name = name
		if err = e.collectValue(ch, l, mv, value); err != nil {
			errs = append(errs, err)
		}
	}
	if len(errs) > 0 {
		return errs
	}
	return nil
}

func (e *Exporter) collectValue(ch chan<- prometheus.Metric, l *latency.Latency, mv metrics.Metric, value interface{}) error {
	switch v := value.(type) {
	case float64:
		err := mv.Send(ch, v)
		if err != nil {
			return fmt.Errorf("unable to send value '%v': %s", value, err)
		}
		return nil
	case map[string]interface{}:
		mv.Subsystem = mv.Name
		return e.collectMapValue(ch, l, mv, v)
	case []interface{}:
		mv.Subsystem = mv.Name
		return e.collectSliceValue(ch, l, mv, v)
	case string:
		switch mv.Name {
		case
			"engine",
			"instanceID",
			"instanceResourceID",
			"name",
			"device",
			"interface",
			"mountPoint",
			"vmlimit",
			"uptime":
			// skipping those values
		case "timestamp":
			t, err := time.Parse(time.RFC3339, v)
			if err != nil {
				return err
			}
			l.TakeOldest(t)
		default:
			return fmt.Errorf("unsupported metric '%s': %s", mv.BuildFQName(), v)
		}
	default:
		return fmt.Errorf("unsupported type for metric '%s': %T", mv.BuildFQName(), value)
	}
	return nil
}

func (e *Exporter) collectSliceValue(ch chan<- prometheus.Metric, l *latency.Latency, mv metrics.Metric, value []interface{}) error {
	var errs errs
	for i, v := range value {
		m := mv
		m.LabelsValues = append(m.LabelsValues, strconv.Itoa(i))
		if err := e.collectValue(ch, l, m, v); err != nil {
			errs = append(errs, err)
		}
	}
	if len(errs) > 0 {
		return errs
	}
	return nil
}

func (e *Exporter) collectMapValue(ch chan<- prometheus.Metric, l *latency.Latency, mv metrics.Metric, value map[string]interface{}) error {
	var errs errs
	switch mv.Subsystem {
	case "fileSys":
		if _, ok := value["used"]; ok {
			if _, ok := value["total"]; ok {
				total, _ := value["total"].(float64)
				used, _ := value["used"].(float64)
				value["avail"] = total - used
			}
		}

		name := ""
		mountPoint := ""
		for metricName, v := range value {
			if label, ok := v.(string); ok {
				switch metricName {
				case "name":
					name = label
				case "mountPoint":
					mountPoint = label
				}
			}
			if name != "" && mountPoint != "" {
				break
			}
		}
		mv.LabelsValues = append(mv.LabelsValues, name, mountPoint)
	case "diskIO":
		device := ""
		for metricName, v := range value {
			if label, ok := v.(string); ok {
				if metricName == "device" {
					device = label
					break
				}
			}
		}
		mv.LabelsValues = append(mv.LabelsValues, device)
	case "processList", "network":
		for metricName, v := range value {
			if label, ok := v.(string); ok {
				switch {
				case mv.Subsystem == "processList" && metricName == "name",
					mv.Subsystem == "network" && metricName == "interface":
					mv.LabelsValues = append(mv.LabelsValues, label)
				}
			}
		}
	}

	for name, v := range value {
		mv.Name = name
		err := e.collectValue(ch, l, mv, v)
		if err != nil {
			errs = append(errs, err)
		}
	}
	if len(errs) > 0 {
		return errs
	}
	return nil
}

func dbiResourceId(p client.ConfigProvider, instance string) (string, error) {
	svc := rds.New(p)

	input := &rds.DescribeDBInstancesInput{
		DBInstanceIdentifier: aws.String(instance),
	}

	result, err := svc.DescribeDBInstances(input)
	if err != nil {
		if err, ok := err.(awserr.Error); ok {
			switch err.Code() {
			case rds.ErrCodeDBInstanceNotFoundFault:
				return "", fmt.Errorf("instance '%s' not found: %s", instance, err)
			}
		}
		return "", err
	}

	if len(result.DBInstances) != 1 {
		return "", fmt.Errorf("got '%d' instances but expected just one", len(result.DBInstances))
	}

	return aws.StringValue(result.DBInstances[0].DbiResourceId), nil
}

func events(p client.ConfigProvider, logStreamName, instance string, interval time.Duration) (events map[string]interface{}, err error) {
	svc := cloudwatchlogs.New(p)

	// PMM-2165
	//GetLogEventsInput := &cloudwatchlogs.GetLogEventsInput{
	//	StartTime:     aws.Int64(aws.TimeUnixMilli(time.Now().UTC().Add(-interval))),
	//	Limit:         aws.Int64(1),
	//	LogGroupName:  aws.String(logGroupName),
	//	LogStreamName: aws.String(logStreamName),
	//}
	FilterLogEventsInput := &cloudwatchlogs.FilterLogEventsInput{
		FilterPattern:  aws.String(fmt.Sprintf(`{ $.instanceID = "%s" }`, instance)),
		StartTime:      aws.Int64(aws.TimeUnixMilli(time.Now().UTC().Add(-interval))),
		Limit:          aws.Int64(1),
		LogGroupName:   aws.String(logGroupName),
		LogStreamNames: aws.StringSlice([]string{logStreamName}),
	}

	var errs errs
	events = map[string]interface{}{}
	// PMM-2165
	//fn := func(logs *cloudwatchlogs.GetLogEventsOutput, lastPage bool) (cont bool) {
	fn := func(logs *cloudwatchlogs.FilterLogEventsOutput, lastPage bool) (cont bool) {
		cont = !lastPage

		if len(logs.Events) == 0 {
			return
		}

		var message interface{}
		err := json.Unmarshal([]byte(*logs.Events[0].Message), &message)
		if err != nil {
			errs = append(errs, err)
			return
		}

		var ok bool
		events, ok = message.(map[string]interface{})
		if !ok {
			errs = append(errs, fmt.Errorf("expected type %T but got %T", events, message))
			return
		}

		return
	}
	// PMM-2165
	//err = svc.GetLogEventsPages(GetLogEventsInput, fn)
	err = svc.FilterLogEventsPages(FilterLogEventsInput, fn)
	if err != nil {
		return nil, fmt.Errorf("unable to get logs: %s", err)
	}

	if len(errs) > 0 {
		// Returning values and errors here is correct as we collect data on best effort.
		return events, errs
	}

	return events, nil
}

type errs []error

func (errs errs) Error() string {
	if len(errs) == 0 {
		return ""
	}
	buf := &bytes.Buffer{}
	for _, err := range errs {
		fmt.Fprintf(buf, "\n* %s", err)
	}
	return buf.String()
}
