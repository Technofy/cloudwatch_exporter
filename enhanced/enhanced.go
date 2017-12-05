package enhanced

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"sync"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatchlogs"
	"github.com/percona/rds_exporter/config"
	"github.com/percona/rds_exporter/sessions"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/common/log"
)

const (
	namespace    = "rdsosmetrics"
	logGroupName = "RDSOSMetrics"
)

type Exporter struct {
	Settings *config.Settings
	Sessions *sessions.Sessions

	// For simplification assume this sync.Map always store values of type *prometheus.Desc.
	// If for any reason this leads to problems feel free to add type checks to increase safety.
	desc sync.Map
}

func New(settings *config.Settings, sessions *sessions.Sessions) *Exporter {
	return &Exporter{
		Settings: settings,
		Sessions: sessions,
	}
}

func (e *Exporter) Describe(ch chan<- *prometheus.Desc) {
	wg := sync.WaitGroup{}
	defer wg.Wait()

	instances := e.Settings.Config().Instances
	wg.Add(len(instances))
	for _, instance := range instances {
		go func(region string) {
			err := e.describeRegion(region)
			if err != nil {
				log.Error(err)
			}
			wg.Done()
		}(instance.Region)
	}
	wg.Wait()

	e.desc.Range(
		func(key, value interface{}) bool {
			ch <- value.(*prometheus.Desc)
			return true
		},
	)
}

func (e *Exporter) describeRegion(regionID string) error {
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String(regionID),
	}))
	svc := cloudwatchlogs.New(sess)
	DescribeLogStreamsOutput, err := svc.DescribeLogStreams(&cloudwatchlogs.DescribeLogStreamsInput{
		LogGroupName: aws.String(logGroupName),
	})
	if err != nil {
		return fmt.Errorf("region '%s' is without RDS Enhanced Monitoring: %s", regionID, err)
	}

	wg := sync.WaitGroup{}
	defer wg.Wait()
	logStreams := DescribeLogStreamsOutput.LogStreams
	wg.Add(len(logStreams))
	for _, logStream := range logStreams {
		go func(logStreamName *string) {
			err := e.describeLogStream(logStreamName, svc)
			if err != nil {
				log.Error(err)
			}
			wg.Done()
		}(logStream.LogStreamName)
	}

	return nil
}

func (e *Exporter) collectRegion(ch chan<- prometheus.Metric, regionID string) error {
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String(regionID),
	}))
	svc := cloudwatchlogs.New(sess)
	DescribeLogStreamsOutput, err := svc.DescribeLogStreams(&cloudwatchlogs.DescribeLogStreamsInput{
		LogGroupName: aws.String(logGroupName),
	})
	if err != nil {
		// region without RDS Enhanced Monitoring
		return fmt.Errorf("region '%s' is without RDS Enhanced Monitoring: %s", regionID, err)
	}

	wg := sync.WaitGroup{}
	defer wg.Wait()
	logStreams := DescribeLogStreamsOutput.LogStreams
	wg.Add(len(logStreams))
	for _, logStream := range logStreams {
		go func(logStreamName *string) {
			err := e.collectLogStream(ch, regionID, logStreamName, svc)
			if err != nil {
				log.Error(err)
			}
			wg.Done()
		}(logStream.LogStreamName)
	}

	return nil
}

func (e *Exporter) describeLogStream(logStreamName *string, svc *cloudwatchlogs.CloudWatchLogs) error {
	GetLogEventsOutput, err := svc.GetLogEvents(&cloudwatchlogs.GetLogEventsInput{
		Limit:         aws.Int64(1),
		LogGroupName:  aws.String(logGroupName),
		LogStreamName: logStreamName,
	})
	if err != nil {
		return err
	}

	var message interface{}
	err = json.Unmarshal([]byte(*GetLogEventsOutput.Events[0].Message), &message)
	if err != nil {
		return err
	}

	wg := sync.WaitGroup{}
	defer wg.Wait()

	messages, ok := message.(map[string]interface{})
	if !ok {
		return fmt.Errorf("unsupported message type: %T", message)
	}
	wg.Add(len(messages))
	for key, value := range messages {
		go func(key string, value interface{}) {
			err := e.describeValue(key, value)
			if err != nil {
				log.Error(err)
			}
			wg.Done()
		}(key, value)
	}

	return nil
}

func (e *Exporter) describeValue(key string, value interface{}) error {
	switch v := value.(type) {
	case float64:
		e.addDesc("", key)
	case map[string]interface{}:
		for kkey, vvvalue := range v {
			switch vvvalue.(type) {
			case float64:
				e.addDesc(key, kkey)
			}
		}
	case []interface{}:
		for i, u := range v {
			switch vvvalue := u.(type) {
			case map[string]interface{}:
				for kkey, vvvvalue := range vvvalue {
					switch vvvvalue.(type) {
					case float64:
						e.addDesc(key, strconv.Itoa(i)+"_"+kkey)
					}
				}
			}
		}
	case string:
		switch key {
		case
			"engine",
			"instanceID",
			"instanceResourceID",
			"timestamp",
			"uptime":
			// skipping those values
		default:
			return fmt.Errorf("unsupported key '%s' when describing value", key)
		}
	default:
		return fmt.Errorf("unsupported value type '%T' for key '%s' when describing value", value, key)
	}

	return nil
}

func (e *Exporter) addDesc(subsystem string, name string) {
	FQName := prometheus.BuildFQName(namespace, subsystem, name)

	if _, ok := e.desc.Load(FQName); ok {
		return
	}

	e.desc.Store(
		FQName,
		prometheus.NewDesc(
			strings.ToLower(FQName),
			"Automatically parsed metric from "+logGroupName+" Log Group",
			[]string{
				"region",
				"instanceID",
			},
			nil,
		),
	)
}

func (e *Exporter) updateMetric(ch chan<- prometheus.Metric, region string, instanceID string, subsystem string, name string, value float64) {
	FQName := prometheus.BuildFQName(namespace, subsystem, name)
	if v, ok := e.desc.Load(FQName); ok {
		ch <- prometheus.MustNewConstMetric(
			v.(*prometheus.Desc),
			prometheus.UntypedValue,
			value,
			region,
			instanceID,
		)
	}
}

func (e *Exporter) Collect(ch chan<- prometheus.Metric) {
	wg := sync.WaitGroup{}
	defer wg.Wait()

	instances := e.Settings.Config().Instances
	wg.Add(len(instances))
	for _, instance := range instances {
		go func(region string) {
			err := e.collectRegion(ch, region)
			if err != nil {
				log.Error(err)
			}
			wg.Done()
		}(instance.Region)
	}
}

func (e *Exporter) collectLogStream(ch chan<- prometheus.Metric, regionID string, logStreamName *string, svc *cloudwatchlogs.CloudWatchLogs) error {
	GetLogEventsOutput, err := svc.GetLogEvents(&cloudwatchlogs.GetLogEventsInput{
		Limit:         aws.Int64(1),
		LogGroupName:  aws.String(logGroupName),
		LogStreamName: logStreamName,
	})
	if err != nil {
		return err
	}

	var message interface{}
	err = json.Unmarshal([]byte(*GetLogEventsOutput.Events[0].Message), &message)
	if err != nil {
		return err
	}

	wg := sync.WaitGroup{}
	defer wg.Wait()

	messages, ok := message.(map[string]interface{})
	if !ok {
		return fmt.Errorf("unsupported message type: %T", message)
	}
	wg.Add(len(messages))
	instanceID := messages["instanceID"].(string)
	for key, value := range messages {
		go func(key string, value interface{}) {
			err := e.collectValue(ch, regionID, instanceID, key, value)
			if err != nil {
				log.Error(err)
			}
			wg.Done()
		}(key, value)
	}

	return nil
}

func (e *Exporter) collectValue(ch chan<- prometheus.Metric, regionID string, instanceID string, key string, value interface{}) error {
	switch v := value.(type) {
	case float64:
		e.updateMetric(ch, regionID, instanceID, "", key, v)
	case map[string]interface{}:
		for kkey, vvvalue := range v {
			switch vvvvalue := vvvalue.(type) {
			case float64:
				e.updateMetric(ch, regionID, instanceID, key, kkey, vvvvalue)
			}
		}
	case []interface{}:
		for i, u := range v {
			switch vvvalue := u.(type) {
			case map[string]interface{}:
				for kkey, vvvvalue := range vvvalue {
					switch vvvvvalue := vvvvalue.(type) {
					case float64:
						e.updateMetric(ch, regionID, instanceID, key, strconv.Itoa(i)+"_"+kkey, vvvvvalue)
					}
				}
			}
		}
	case string:
		switch key {
		case
			"engine",
			"instanceID",
			"instanceResourceID",
			"timestamp",
			"uptime":
			// skipping those values
		default:
			return fmt.Errorf("unsupported key '%s' when collecting value", key)
		}
	default:
		return fmt.Errorf("unsupported value type '%T' for key '%s' when collecting value", value, key)
	}

	return nil
}
