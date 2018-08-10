package basic

import (
	"sync"
	"time"

	"github.com/prometheus/client_golang/prometheus"

	"github.com/percona/rds_exporter/config"
	"github.com/percona/rds_exporter/latency"
	"github.com/percona/rds_exporter/sessions"
)

//go:generate go run generate/main.go generate/utils.go

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
}

type Exporter struct {
	Settings *config.Settings
	Sessions *sessions.Sessions

	// Metrics
	Metrics           []Metric
	ErroneousRequests prometheus.Counter
	TotalRequests     prometheus.Counter
}

// New creates a new instance of a Exporter.
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
			NewScraper(instance, e, ch).Scrape()
			wg.Done()
		}(instance)
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
