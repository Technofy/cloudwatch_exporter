package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatch"
	"github.com/prometheus/client_golang/prometheus"
)

var (
	PeriodSeconds = 60
	DelaySeconds  = 600
	RangeSeconds  = 600
)

func getLatestDatapoint(datapoints []*cloudwatch.Datapoint) *cloudwatch.Datapoint {
	var latest *cloudwatch.Datapoint = nil

	for dp := range datapoints {
		if latest == nil || latest.Timestamp.Before(*datapoints[dp].Timestamp) {
			latest = datapoints[dp]
		}
	}

	return latest
}

// scrape makes the required calls to AWS CloudWatch by using the parameters in the Collector
// Once converted into Prometheus format, the metrics are pushed on the ch channel.
func scrape(instance, region string, collector *Collector, ch chan<- prometheus.Metric) {
	session := session.Must(session.NewSession(&aws.Config{
		Region: aws.String(region),
	}))

	svc := cloudwatch.New(session)

	labels := []string{}
	labels = append(labels, instance, region)

	wg := &sync.WaitGroup{}
	wg.Add(len(collector.Metrics))
	for _, metric := range collector.Metrics {
		go func(metric Metric) {
			err := scrapeMetric(svc, metric, instance, collector, ch, labels)
			if err != nil {
				fmt.Println(err)
			}
			wg.Done()
		}(metric)
	}
	wg.Wait()
}

func scrapeMetric(svc *cloudwatch.CloudWatch, metric Metric, instance string, collector *Collector, ch chan<- prometheus.Metric, labels []string) error {
	now := time.Now()
	end := now.Add(time.Duration(-DelaySeconds) * time.Second)

	params := &cloudwatch.GetMetricStatisticsInput{
		EndTime:   aws.Time(end),
		StartTime: aws.Time(end.Add(time.Duration(-RangeSeconds) * time.Second)),

		Period:     aws.Int64(int64(PeriodSeconds)),
		MetricName: aws.String(metric.Name),
		Namespace:  aws.String("AWS/RDS"),
		Dimensions: []*cloudwatch.Dimension{},
		Statistics: aws.StringSlice([]string{"Average"}),
		Unit:       nil,
	}

	params.Dimensions = append(params.Dimensions, &cloudwatch.Dimension{
		Name:  aws.String("DBInstanceIdentifier"),
		Value: aws.String(instance),
	})

	// Call CloudWatch to gather the datapoints
	resp, err := svc.GetMetricStatistics(params)
	collector.TotalRequests.Inc()

	if err != nil {
		collector.ErroneousRequests.Inc()
		return err
	}

	// There's nothing in there, don't publish the metric
	if len(resp.Datapoints) == 0 {
		return nil
	}

	// Pick the latest datapoint
	dp := getLatestDatapoint(resp.Datapoints)
	ch <- prometheus.MustNewConstMetric(metric.Desc, prometheus.GaugeValue, aws.Float64Value(dp.Average), labels...)

	return nil
}
