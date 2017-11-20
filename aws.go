package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatch"
	"github.com/percona/rds_exporter/config"
	"github.com/prometheus/client_golang/prometheus"
)

var (
	Period = 60 * time.Second
	Delay  = 600 * time.Second
	Range  = 600 * time.Second
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
func scrape(instance config.Instance, collector *Collector, ch chan<- prometheus.Metric) {
	awsConfig := &aws.Config{
		Region: aws.String(instance.Region),
	}

	if instance.AwsAccessKey != "" || instance.AwsSecretKey != "" {
		awsConfig.Credentials = credentials.NewStaticCredentials(
			instance.AwsAccessKey,
			instance.AwsSecretKey,
			"",
		)
	}

	session := session.Must(session.NewSession(awsConfig))

	svc := cloudwatch.New(session)

	labels := []string{}
	labels = append(labels, instance.Instance, instance.Region)

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

func scrapeMetric(svc *cloudwatch.CloudWatch, metric Metric, instance config.Instance, collector *Collector, ch chan<- prometheus.Metric, labels []string) error {
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
		Value: aws.String(instance.Instance),
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
