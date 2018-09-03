package enhanced

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
	"sync"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudwatchlogs"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/common/log"

	"github.com/percona/rds_exporter/config"
	"github.com/percona/rds_exporter/latency"
	"github.com/percona/rds_exporter/sessions"
)

//go:generate go run generate/main.go

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

type Metric struct {
	Name string
	Desc *prometheus.Desc
	Unit float64
}

type Exporter struct {
	Settings *config.Settings
	Sessions *sessions.Sessions

	// Metrics
	Metrics           map[string]Metric
	ErroneousRequests prometheus.Counter
	TotalRequests     prometheus.Counter
}

func New(settings *config.Settings, sessions *sessions.Sessions) *Exporter {
	return &Exporter{
		Settings: settings,
		Sessions: sessions,
		Metrics:  Metrics,
		ErroneousRequests: prometheus.NewCounter(prometheus.CounterOpts{
			Name: "rds_exporter_erroneous_requests",
			Help: "The number of erroneous API request made to CloudWatch.",
		}),
		TotalRequests: prometheus.NewCounter(prometheus.CounterOpts{
			Name: "rds_exporter_requests_total",
			Help: "API requests made to CloudWatch",
		}),
	}
}

func (e *Exporter) Describe(ch chan<- *prometheus.Desc) {
	// RDS metrics
	for _, m := range e.Metrics {
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

	instances := e.Settings.Config().Instances
	wg.Add(len(instances))
	for _, instance := range instances {
		go func(instance config.Instance) {
			defer wg.Done()

			err := e.collectInstance(ch, instance)
			if err != nil {
				log.Error(err)
			}
		}(instance)
	}
}

func (e *Exporter) collectInstance(ch chan<- prometheus.Metric, instance config.Instance) error {
	// Latency metric
	l := &latency.Latency{}

	// Collect all values
	err := e.collectValues(ch, instance, l)

	// Collect latency metric
	if !l.IsZero() {
		ch <- prometheus.MustNewConstMetric(latency.Desc, prometheus.GaugeValue, float64(l.Duration().Seconds()), instanceLabels(instance)...)
	}
	return err
}

func (e *Exporter) collectValues(ch chan<- prometheus.Metric, instance config.Instance, l *latency.Latency) error {
	sess := e.Sessions.Get(instance)
	svc := cloudwatchlogs.New(sess)

	values := map[string]interface{}{}
	FilterLogEventsInput := &cloudwatchlogs.FilterLogEventsInput{
		StartTime:     aws.Int64(aws.TimeUnixMilli(time.Now().UTC().Add(-instance.Interval))),
		Limit:         aws.Int64(1),
		LogGroupName:  aws.String(logGroupName),
		FilterPattern: aws.String(fmt.Sprintf(`{ $.instanceID = "%s" }`, instance.Instance)),
	}
	var err error
	fn := func(logs *cloudwatchlogs.FilterLogEventsOutput, lastPage bool) (cont bool) {
		cont = !lastPage

		if len(logs.Events) == 0 {
			return
		}

		var message interface{}
		err = json.Unmarshal([]byte(*logs.Events[0].Message), &message)
		if err != nil {
			return
		}

		var ok bool
		values, ok = message.(map[string]interface{})
		if !ok {
			return
		}
		return
	}
	err = svc.FilterLogEventsPages(FilterLogEventsInput, fn)
	if err != nil {
		return fmt.Errorf("unable to get logs for instance %s: %s", instance.Instance, err)
	}

	var errs errs
	for metricName, value := range values {
		if err = e.collectValue(ch, instance, l, defaultNamespace, "General", metricName, value); err != nil {
			errs = append(errs, err)
		}
	}
	if len(errs) > 0 {
		return errs
	}
	return nil
}

func (e *Exporter) collectValue(ch chan<- prometheus.Metric, instance config.Instance, l *latency.Latency, namespace string, subsystem string, metricName string, value interface{}, extraLabelsValues ...string) error {
	switch v := value.(type) {
	case float64:
		err := e.sendMetric(ch, instance, namespace, subsystem, metricName, v, extraLabelsValues...)
		if err != nil {
			return fmt.Errorf("unable to parse value '%v': %s", value, err)
		}
		return nil
	case map[string]interface{}:
		return e.collectMapValue(ch, instance, l, namespace, metricName, v, extraLabelsValues...)
	case []interface{}:
		return e.collectSliceValue(ch, instance, l, namespace, metricName, v, extraLabelsValues...)
	case string:
		switch metricName {
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
			return fmt.Errorf("unsupported metric '%s': %s", prometheus.BuildFQName(namespace, subsystem, metricName), v)
		}
	default:
		return fmt.Errorf("unsupported type for metric '%s': %T", prometheus.BuildFQName(namespace, subsystem, metricName), value)
	}
	return nil
}

func (e *Exporter) collectSliceValue(ch chan<- prometheus.Metric, instance config.Instance, l *latency.Latency, namespace string, subsystem string, value []interface{}, extraLabelsValuesIN ...string) error {
	var errs errs
	for i, u := range value {
		extraLabelsValues := append(extraLabelsValuesIN, strconv.Itoa(i))
		metrics, ok := u.(map[string]interface{})
		if !ok {
			errs = append(errs, fmt.Errorf("%s: wrong value type for metrics: %T", subsystem, u))
			continue
		}
		err := e.collectMapValue(ch, instance, l, namespace, subsystem, metrics, extraLabelsValues...)
		if err != nil {
			errs = append(errs, err)
			continue
		}
	}
	if len(errs) > 0 {
		return errs
	}
	return nil
}

func (e *Exporter) collectMapValue(ch chan<- prometheus.Metric, instance config.Instance, l *latency.Latency, namespace string, subsystem string, value map[string]interface{}, extraLabelsValuesIN ...string) error {
	var errs errs
	extraLabelsValues := extraLabelsValuesIN

	if subsystem == "fileSys" {
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
		extraLabelsValues = append(extraLabelsValuesIN, name, mountPoint)
	}

	if subsystem == "diskIO" {
		device := ""
		for metricName, v := range value {
			if label, ok := v.(string); ok {
				if metricName == "device" {
					device = label
					break
				}
			}
		}
		extraLabelsValues = append(extraLabelsValuesIN, device)
	}

	switch subsystem {
	case "processList", "network", "diskIO":
		for metricName, v := range value {
			if label, ok := v.(string); ok {
				switch {
				case subsystem == "processList" && metricName == "name",
					subsystem == "network" && metricName == "interface":
					extraLabelsValues = append(extraLabelsValuesIN, label)
				}
			}
		}
	}

	for metricName, v := range value {
		err := e.collectValue(ch, instance, l, namespace, subsystem, metricName, v, extraLabelsValues...)
		if err != nil {
			errs = append(errs, err)
		}
	}
	if len(errs) > 0 {
		return errs
	}
	return nil
}

func (e *Exporter) sendMetric(ch chan<- prometheus.Metric, instance config.Instance, namespaceIN, subsystem, metricName string, value float64, extraLabelsValuesIN ...string) error {
	namespace, subsystem, name, extraLabels, extraLabelsValues := MapToNode(namespaceIN, subsystem, metricName, extraLabelsValuesIN...)
	FQName := prometheus.BuildFQName(namespace, subsystem, name)

	if len(extraLabels) != len(extraLabelsValues) {
		return fmt.Errorf("len(labels) != len(labelsValues) for metric '%s': len(%v) != len(%v): ", FQName, extraLabels, extraLabelsValues)
	}

	metric, ok := e.Metrics[FQName]
	if !ok {
		return fmt.Errorf("unknown metric %s", FQName)
	}

	labels := instanceLabels(instance)
	if len(extraLabelsValues) > 0 {
		labels = append(labels, extraLabelsValues...)
	}

	ch <- prometheus.MustNewConstMetric(
		metric.Desc,
		prometheus.GaugeValue,
		value*metric.Unit,
		labels...,
	)
	return nil
}

func instanceLabels(instance config.Instance) []string {
	return []string{
		instance.Instance,
		instance.Region,
	}
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
