package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudwatch"
	"github.com/percona/rds_exporter/config"
	"github.com/prometheus/client_golang/prometheus"
)

var (
	Period = 60 * time.Second
	Delay  = 600 * time.Second
	Range  = 600 * time.Second

	latencyDesc = prometheus.NewDesc(
		"rds_latency",
		"The difference between the current time and timestamp in the metric itself",
		[]string{"instance", "region"},
		map[string]string(nil),
	)
)

type Scrape struct {
	// params
	instance  config.Instance
	collector *Collector
	ch        chan<- prometheus.Metric

	// internal
	svc     *cloudwatch.CloudWatch
	labels  []string
	latency *Latency
}

func NewScrape(instance config.Instance, collector *Collector, ch chan<- prometheus.Metric) *Scrape {
	// Create CloudWatch client
	sess := collector.Sessions.Get(instance)
	svc := cloudwatch.New(sess)

	// Create labels for all metrics
	labels := []string{}
	labels = append(labels, instance.Instance, instance.Region)

	return &Scrape{
		// params
		instance:  instance,
		collector: collector,
		ch:        ch,

		// internal
		svc:     svc,
		labels:  labels,
		latency: &Latency{},
	}
}

func getLatestDatapoint(datapoints []*cloudwatch.Datapoint) *cloudwatch.Datapoint {
	var latest *cloudwatch.Datapoint = nil

	for dp := range datapoints {
		if latest == nil || latest.Timestamp.Before(*datapoints[dp].Timestamp) {
			latest = datapoints[dp]
		}
	}

	return latest
}

// Scrape makes the required calls to AWS CloudWatch by using the parameters in the Collector.
// Once converted into Prometheus format, the metrics are pushed on the ch channel.
func (s *Scrape) Scrape() {
	wg := &sync.WaitGroup{}
	wg.Add(len(s.collector.Metrics))
	for _, metric := range s.collector.Metrics {
		go func(metric Metric) {
			err := s.scrapeMetric(metric)
			if err != nil {
				fmt.Println(err)
			}

			wg.Done()
		}(metric)
	}
	wg.Wait()

	// Generate latency metric
	if !s.latency.Timestamp.IsZero() {
		latency := time.Since(s.latency.Timestamp).Seconds()
		s.ch <- prometheus.MustNewConstMetric(latencyDesc, prometheus.GaugeValue, float64(latency), s.labels...)
	}
}

func (s *Scrape) scrapeMetric(metric Metric) error {
	now := time.Now()
	end := now.Add(-Delay)

	params := &cloudwatch.GetMetricStatisticsInput{
		EndTime:   aws.Time(end),
		StartTime: aws.Time(end.Add(-Range)),

		Period:     aws.Int64(int64(Period.Seconds())),
		MetricName: aws.String(metric.Name),
		Namespace:  aws.String("AWS/RDS"),
		Dimensions: []*cloudwatch.Dimension{},
		Statistics: aws.StringSlice([]string{"Average"}),
		Unit:       nil,
	}

	params.Dimensions = append(params.Dimensions, &cloudwatch.Dimension{
		Name:  aws.String("DBInstanceIdentifier"),
		Value: aws.String(s.instance.Instance),
	})

	// Call CloudWatch to gather the datapoints
	resp, err := s.svc.GetMetricStatistics(params)
	s.collector.TotalRequests.Inc()

	if err != nil {
		s.collector.ErroneousRequests.Inc()
		return err
	}

	// There's nothing in there, don't publish the metric
	if len(resp.Datapoints) == 0 {
		return nil
	}

	// Pick the latest datapoint
	dp := getLatestDatapoint(resp.Datapoints)

	// Take the oldest timestamp for latency metric
	s.latency.TakeOldest(aws.TimeValue(dp.Timestamp))

	// Get the metric.
	v := aws.Float64Value(dp.Average)
	switch metric.Name {
	case "EngineUptime":
		// "Fake EngineUptime -> node_boot_time with time.Now().Unix() - EngineUptime."
		v = float64(time.Now().Unix() - int64(v))
	}

	// Send metric.
	s.ch <- prometheus.MustNewConstMetric(metric.Desc, prometheus.GaugeValue, v, s.labels...)

	return nil
}

type Latency struct {
	Timestamp time.Time
	sync.RWMutex
}

func (l *Latency) TakeOldest(t time.Time) {
	l.Lock()
	defer l.Unlock()
	if l.Timestamp.IsZero() || t.Before(l.Timestamp) {
		l.Timestamp = t
	}
}
