package basic

import (
	"sync"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudwatch"
	"github.com/prometheus/client_golang/prometheus"

	"github.com/percona/rds_exporter/config"
	"github.com/percona/rds_exporter/latency"
)

var (
	Period = 60 * time.Second
	Delay  = 600 * time.Second
	Range  = 600 * time.Second
)

type Scraper struct {
	// params
	instance *config.Instance
	exporter *Exporter
	ch       chan<- prometheus.Metric

	// internal
	svc     *cloudwatch.CloudWatch
	labels  []string
	latency *latency.Latency
}

func NewScraper(instance *config.Instance, exporter *Exporter, ch chan<- prometheus.Metric) *Scraper {
	// Create CloudWatch client
	sess, _ := exporter.Sessions.GetSession(instance.Region, instance.Instance)
	svc := cloudwatch.New(sess)

	// Create labels for all metrics
	labels := []string{
		instance.Instance,
		instance.Region,
	}

	return &Scraper{
		// params
		instance: instance,
		exporter: exporter,
		ch:       ch,

		// internal
		svc:     svc,
		labels:  labels,
		latency: &latency.Latency{},
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

// Scrape makes the required calls to AWS CloudWatch by using the parameters in the Exporter.
// Once converted into Prometheus format, the metrics are pushed on the ch channel.
func (s *Scraper) Scrape() {
	s.scrape()

	// Generate latency metric
	if !s.latency.IsZero() {
		s.ch <- prometheus.MustNewConstMetric(latency.Desc, prometheus.GaugeValue, float64(s.latency.Duration().Seconds()), s.labels...)
	}
}

func (s *Scraper) scrape() {
	wg := &sync.WaitGroup{}
	defer wg.Wait()

	wg.Add(len(s.exporter.Metrics))
	for _, metric := range s.exporter.Metrics {
		go func(metric Metric) {
			err := s.scrapeMetric(metric)
			if err != nil {
				s.exporter.l.With("metric", metric.Name).Error(err)
			}

			wg.Done()
		}(metric)
	}
}
func (s *Scraper) scrapeMetric(metric Metric) error {
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
	s.exporter.TotalRequests.Inc()

	if err != nil {
		s.exporter.ErroneousRequests.Inc()
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
