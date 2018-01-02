package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatch"
	"github.com/prometheus/client_golang/prometheus"
	"time"
	"regexp"
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
		metric := &collector.Template.Metrics[m]

		now := time.Now()
		end := now.Add(time.Duration(-metric.ConfMetric.DelaySeconds) * time.Second)

		params := &cloudwatch.GetMetricStatisticsInput{
			EndTime:   aws.Time(end),
			StartTime: aws.Time(end.Add(time.Duration(-metric.ConfMetric.RangeSeconds) * time.Second)),

			Period:     aws.Int64(int64(metric.ConfMetric.PeriodSeconds)),
			MetricName: aws.String(metric.ConfMetric.Name),
			Namespace:  aws.String(metric.ConfMetric.Namespace),
			Dimensions: []*cloudwatch.Dimension{},
			Statistics: []*string{},
			Unit:       nil,
		}

		dimensions:=[]*cloudwatch.Dimension{}


		for _, stat := range metric.ConfMetric.Statistics {
			params.Statistics = append(params.Statistics, aws.String(stat))
		}

		labels := make([]string, 0, len(metric.LabelNames))

		// Loop through the dimensions selects to build the filters and the labels array
		for dim := range metric.ConfMetric.DimensionsSelect {
			for val := range metric.ConfMetric.DimensionsSelect[dim] {
				dimValue := metric.ConfMetric.DimensionsSelect[dim][val]

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

		if (len(dimensions)>0){
			labels = append(labels, collector.Template.Task.Name)
			params.Dimensions=dimensions
			scrapeSingleDataPoint(collector,ch,params,metric,labels,svc)
			continue
		}

		// Get all the metric to select the ones who'll match the regex
		result, err := svc.ListMetrics(&cloudwatch.ListMetricsInput{
			MetricName: aws.String(metric.ConfMetric.Name),
			Namespace:  aws.String(metric.ConfMetric.Namespace),
		})

		if err != nil {
			fmt.Println(err)
			continue
		}
		
		//This map will hold dimensions name which has been already collected
		valueCollected :=  map[string]bool{}

		for dimensions := range metric.ConfMetric.DimensionsSelectRegex {
			dimRegex := metric.ConfMetric.DimensionsSelectRegex[dimensions]

			// Replace $_target token by the actual URL target
			if dimRegex == "$_target" {
				dimRegex = collector.Target
			}

			//Loop through all the dimensions for the metric given
			for _,met := range result.Metrics {
					for _,dim := range met.Dimensions {

						//Select the one which match the regex
						match,_:=regexp.MatchString(dimRegex,*dim.Value)
						if _, ok := valueCollected[*dim.Value];  match && !ok  {
							//Create the request and send it to the prometheus lib
							labels := make([]string, 0, len(metric.LabelNames))
							
							params.Dimensions = []*cloudwatch.Dimension{}
							params.Dimensions = append(params.Dimensions, &cloudwatch.Dimension{
								Name:  aws.String(*dim.Name),
								Value: aws.String(*dim.Value),
							})
			
							labels = append(labels, *dim.Value)
	
							labels = append(labels, collector.Template.Task.Name)	
							scrapeSingleDataPoint(collector,ch,params,metric,labels,svc)

							valueCollected[*dim.Value]=true
						}
					}
				
				
			}
		}


	
	}
}

//Send a single dataPoint to the Prometheus lib
func scrapeSingleDataPoint(collector *cwCollector, ch chan<- prometheus.Metric,params *cloudwatch.GetMetricStatisticsInput,metric *cwMetric,labels []string,svc *cloudwatch.CloudWatch) error {
	
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