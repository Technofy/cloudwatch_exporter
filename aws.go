package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatch"
	"github.com/prometheus/client_golang/prometheus"
	"regexp"
	"strings"
	"time"
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

// scrape makes the required calls to AWS CloudWatch by using the parameters in the cwCollector
// Once converted into Prometheus format, the metrics are pushed on the ch channel.
func scrape(collector *cwCollector, ch chan<- prometheus.Metric) {
	session := session.Must(session.NewSession(&aws.Config{
		Region: aws.String(collector.Region),
	}))

	svc := cloudwatch.New(session)
	for m := range collector.Template.Metrics {
		templMetric := &collector.Template.Metrics[m]

		now := time.Now()
		end := now.Add(time.Duration(-templMetric.ConfMetric.DelaySeconds) * time.Second)

		params := &cloudwatch.GetMetricStatisticsInput{
			EndTime:   aws.Time(end),
			StartTime: aws.Time(end.Add(time.Duration(-templMetric.ConfMetric.RangeSeconds) * time.Second)),

			Period:     aws.Int64(int64(templMetric.ConfMetric.PeriodSeconds)),
			MetricName: aws.String(templMetric.ConfMetric.Name),
			Namespace:  aws.String(templMetric.ConfMetric.Namespace),
			Dimensions: []*cloudwatch.Dimension{},
			Statistics: []*string{},
			Unit:       nil,
		}

		var dimensions []*cloudwatch.Dimension

		//This map will hold dimensions name which has been already collected
		valueCollected := map[string]bool{}

		if len(templMetric.ConfMetric.DimensionsSelectRegex) == 0 {
			templMetric.ConfMetric.DimensionsSelectRegex = map[string]string{}
		}

		//Check for dimensions who does not have either select or dimensions select_regex and make them select everything using regex
		for _, dimension := range templMetric.ConfMetric.Dimensions {
			_, found := templMetric.ConfMetric.DimensionsSelect[dimension]
			_, found2 := templMetric.ConfMetric.DimensionsSelectRegex[dimension]

			if !found && !found2 {
				templMetric.ConfMetric.DimensionsSelectRegex[dimension] = ".*"
			}
		}

		for _, stat := range templMetric.ConfMetric.Statistics {
			params.Statistics = append(params.Statistics, aws.String(stat))
		}

		labels := make([]string, 0, len(templMetric.LabelNames))

		// Loop through the dimensions selects to build the filters and the labels array
		for dim := range templMetric.ConfMetric.DimensionsSelect {
			for val := range templMetric.ConfMetric.DimensionsSelect[dim] {
				dimValue := templMetric.ConfMetric.DimensionsSelect[dim][val]

				// Replace $_target token by the actual URL target
				if dimValue == "$_target" {
					dimValue = collector.Target
				}

				dimensions = append(dimensions, &cloudwatch.Dimension{
					Name:  aws.String(dim),
					Value: aws.String(dimValue),
				})

				labels = append(labels, dimValue)
			}
		}

		if len(dimensions) > 0 {
			labels = append(labels, collector.Template.Task.Name)
			params.Dimensions = dimensions
			scrapeSingleDataPoint(collector, ch, params, templMetric, labels, svc)
		}

		//If no regex is specified, continue
		if len(templMetric.ConfMetric.DimensionsSelectRegex) == 0 {
			continue
		}

		var nextToken *string = nil
		var cwMetrics []*cloudwatch.Metric
		var cwDimensionsFilters []*cloudwatch.DimensionFilter

		for templDim := range templMetric.ConfMetric.DimensionsSelectRegex {
			templDimValue := templMetric.ConfMetric.DimensionsSelectRegex[templDim]

			// Replace $_target token by the actual URL target
			if templDimValue == "$_target" {
				templDimValue = collector.Target
			}

			cwDimensionsFilters = append(cwDimensionsFilters, &cloudwatch.DimensionFilter{
				Name: aws.String(templDim),
			})

			labels = append(labels, templDimValue)
		}

		for {
			result, err := svc.ListMetrics(&cloudwatch.ListMetricsInput{
				MetricName: aws.String(templMetric.ConfMetric.Name),
				Namespace:  aws.String(templMetric.ConfMetric.Namespace),
				Dimensions: cwDimensionsFilters,
				NextToken:  nextToken,
			})
			cwMetrics = append(cwMetrics, result.Metrics...)
			totalRequests.Inc()

			if err != nil {
				fmt.Println(err)
				continue
			}

			if result.NextToken == nil {
				break
			}
			nextToken = result.NextToken
		}

		//For each templMetric returned by aws
		for _, cwMet := range cwMetrics {
			labels := make([]string, 0, len(templMetric.LabelNames))
			dimensions = []*cloudwatch.Dimension{}


			//Try to match each dimensions to the regex
			for _, templDimName := range templMetric.ConfMetric.Dimensions {
				for _, cwDim := range cwMet.Dimensions {

					if templDimName != *cwDim.Name {

						continue

					}

					templDimRegex := templMetric.ConfMetric.DimensionsSelectRegex[templDimName]
					if templDimRegex == "" {
						templDimRegex = "\\b" + strings.Join(templMetric.ConfMetric.DimensionsSelect[*cwDim.Name], "\\b|\\b") + "\\b"
					}

					match, _ := regexp.MatchString(templDimRegex, *cwDim.Value)
					if match {

						dimensions = append(dimensions, &cloudwatch.Dimension{
							Name:  aws.String(*cwDim.Name),
							Value: aws.String(*cwDim.Value),
						})
						labels = append(labels, *cwDim.Value)
					}
				}
			}

			//Cheking if all dimensions matched
			if len(labels) == len(templMetric.ConfMetric.Dimensions) {

				//Checking if this couple of dimensions has already been scraped
				if _, ok := valueCollected[strings.Join(labels, ";")]; ok {
					continue
				}

				//If no, then scrape them
				valueCollected[strings.Join(labels, ";")] = true

				params.Dimensions = dimensions

				labels = append(labels, collector.Template.Task.Name)
				scrapeSingleDataPoint(collector, ch, params, templMetric, labels, svc)

			}
		}
	}
}

//Send a single dataPoint to the Prometheus lib
func scrapeSingleDataPoint(collector *cwCollector, ch chan<- prometheus.Metric, params *cloudwatch.GetMetricStatisticsInput, metric *cwMetric, labels []string, svc *cloudwatch.CloudWatch) error {

	resp, err := svc.GetMetricStatistics(params)
	totalRequests.Inc()

	if err != nil {
		collector.ErroneousRequests.Inc()
		fmt.Println(err)
		return err
	}

	// There's nothing in there, don't publish the metric
	if len(resp.Datapoints) == 0 {
		return nil
	}
	// Pick the latest datapoint
	dp := getLatestDatapoint(resp.Datapoints)

	if dp.Sum != nil {
		ch <- prometheus.MustNewConstMetric(metric.Desc, metric.ValType, float64(*dp.Sum), labels...)
	}

	if dp.Average != nil {
		ch <- prometheus.MustNewConstMetric(metric.Desc, metric.ValType, float64(*dp.Average), labels...)
	}

	if dp.Maximum != nil {
		ch <- prometheus.MustNewConstMetric(metric.Desc, metric.ValType, float64(*dp.Maximum), labels...)
	}

	if dp.Minimum != nil {
		ch <- prometheus.MustNewConstMetric(metric.Desc, metric.ValType, float64(*dp.Minimum), labels...)
	}

	if dp.SampleCount != nil {
		ch <- prometheus.MustNewConstMetric(metric.Desc, metric.ValType, float64(*dp.SampleCount), labels...)
	}
	return nil
}
