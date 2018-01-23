package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatch"
	"github.com/prometheus/client_golang/prometheus"
	"time"
	"regexp"
	"strings"
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

		//This map will hold dimensions name which has been already collected
		valueCollected :=  map[string]bool{}


		if len(metric.ConfMetric.DimensionsSelectRegex) == 0 {
			metric.ConfMetric.DimensionsSelectRegex =  map[string]string{}
		}

		//Check for dimensions who does not have either select or dimensions select_regex and make them select everything using regex
		for _,dimension := range metric.ConfMetric.Dimensions {
			_, found := metric.ConfMetric.DimensionsSelect[dimension]
			_, found2 := metric.ConfMetric.DimensionsSelectRegex[dimension]
			if !found && !found2 {
				metric.ConfMetric.DimensionsSelectRegex[dimension]=".*"
			}
		}



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
		}

		//If no regex is specified, continue
		if (len(metric.ConfMetric.DimensionsSelectRegex)==0){
			continue
		}

		
		// Get all the metric to select the ones who'll match the regex
		result, err := svc.ListMetrics(&cloudwatch.ListMetricsInput{
			MetricName: aws.String(metric.ConfMetric.Name),
			Namespace:  aws.String(metric.ConfMetric.Namespace),
		})
		nextToken:=result.NextToken
		metrics:=result.Metrics
		totalRequests.Inc()

		if err != nil {
			fmt.Println(err)
			continue
		}

		for nextToken!=nil {
			result, err := svc.ListMetrics(&cloudwatch.ListMetricsInput{
				MetricName: aws.String(metric.ConfMetric.Name),
				Namespace:  aws.String(metric.ConfMetric.Namespace),
				NextToken: nextToken,
			})		
			if err != nil {
				fmt.Println(err)
				continue
			}
			nextToken=result.NextToken
			metrics=append(metrics,result.Metrics...)
		}
		
		//For each metric returned by aws
		for _,met := range result.Metrics {
			labels := make([]string, 0, len(metric.LabelNames))
			dimensions=[]*cloudwatch.Dimension{}

			//Try to match each dimensions to the regex
			for _,dim := range met.Dimensions {
				dimRegex:=metric.ConfMetric.DimensionsSelectRegex[*dim.Name]
				if(dimRegex==""){
					dimRegex="\\b"+strings.Join(metric.ConfMetric.DimensionsSelect[*dim.Name],"\\b|\\b")+"\\b"
				}

				match,_:=regexp.MatchString(dimRegex,*dim.Value)
				if match  {
					dimensions=append(dimensions, &cloudwatch.Dimension{
						Name:  aws.String(*dim.Name),
						Value: aws.String(*dim.Value),
					})
					labels = append(labels, *dim.Value)

					
				}
			}

			//Cheking if all dimensions matched
			if len(labels) ==  len(metric.ConfMetric.Dimensions) {	

				//Checking if this couple of dimensions has already been scraped
				if _, ok := valueCollected[strings.Join(labels,";")]; ok {
					continue
				}

				//If no, then scrape them
				valueCollected[strings.Join(labels,";")]=true
			
				params.Dimensions = dimensions

				labels = append(labels, collector.Template.Task.Name)	
				scrapeSingleDataPoint(collector,ch,params,metric,labels,svc)
			
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
