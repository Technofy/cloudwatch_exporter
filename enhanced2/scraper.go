package enhanced2

import (
	"context"
	"math/rand"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatchlogs"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/common/log"

	"github.com/percona/rds_exporter/sessions"
)

type scraper struct {
	instances      []sessions.Instance
	logStreamNames []string
	svc            *cloudwatchlogs.CloudWatchLogs
	startTime      time.Time
	logger         log.Logger
}

func newScraper(session *session.Session, instances []sessions.Instance) *scraper {
	logStreamNames := make([]string, 0, len(instances))
	for _, instance := range instances {
		logStreamNames = append(logStreamNames, instance.ResourceID)
	}

	return &scraper{
		instances:      instances,
		logStreamNames: logStreamNames,
		svc:            cloudwatchlogs.New(session),
		startTime:      time.Now().Add(-2 * time.Minute),
		logger:         log.With("component", "enhanced"),
	}
}

func (s *scraper) start(ctx context.Context, ch chan<- map[string][]prometheus.Metric) {
	interval := time.Minute
	for _, instance := range s.instances {
		if instance.EnhancedMonitoringInterval > 0 && instance.EnhancedMonitoringInterval < interval {
			interval = instance.EnhancedMonitoringInterval
		}
	}
	s.logger.Infof("Updating enhanced metrics every %s.", interval)

	ticker := time.NewTicker(interval)
	defer func() { ticker.Stop() }() // we can redefine ticker below, so use closure

	for {
		metrics := s.scrape(ctx)
		ch <- metrics

		if metrics == nil {
			ticker.Stop()
			time.Sleep(time.Second + time.Duration(rand.Intn(4*int(time.Second)))) // sleep 1-5 seconds
			ticker = time.NewTicker(interval)
			continue
		}

		select {
		case <-ticker.C:
			// nothing
		case <-ctx.Done():
			return
		}
	}
}

func (s *scraper) scrape(ctx context.Context) map[string][]prometheus.Metric {
	input := &cloudwatchlogs.FilterLogEventsInput{
		LogGroupName:   aws.String("RDSOSMetrics"),
		LogStreamNames: aws.StringSlice(s.logStreamNames),
		StartTime:      aws.Int64(aws.TimeUnixMilli(s.startTime)),
	}

	output, err := s.svc.FilterLogEventsWithContext(ctx, input)
	if err != nil {
		s.logger.Errorf("Failed to filter log events: %s.", err)
		return nil
	}

	if output.NextToken != nil {
		s.logger.Error("Pagination is not implemented yet, some data is lost.")
	}

	// FIXME find oldest timestamp
	s.startTime = s.startTime.Add(time.Minute)

	// FIXME find newest metrics
	res := make(map[string][]prometheus.Metric)
	for _, event := range output.Events {
		var instance *sessions.Instance
		for _, i := range s.instances {
			if i.ResourceID == *event.LogStreamName {
				instance = &i
				break
			}
		}
		if instance == nil {
			s.logger.Errorf("Failed to find instance for %s.", *event.LogStreamName)
			continue
		}

		// s.logger.Debugf("Message:\n%s", *event.Message)
		osMetrics, err := parseOSMetrics([]byte(*event.Message))
		if err != nil {
			s.logger.Errorf("Failed to parse metrics for %s/%s (%s): %s.", instance.Region, instance.Instance, instance.ResourceID, err)
			continue
		}
		res[instance.ResourceID] = osMetrics.originalMetrics(instance.Region)
	}
	return res
}
