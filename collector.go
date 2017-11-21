package main

import (
	"sync"
	"time"

	"github.com/percona/rds_exporter/config"
	"github.com/percona/rds_exporter/sessions"
	"github.com/prometheus/client_golang/prometheus"
)

//go:generate go run generate/main.go generate/utils.go

type Metric struct {
	Name string
	Desc *prometheus.Desc
}

type Collector struct {
	Settings *config.Settings
	Sessions *sessions.Sessions
	Metrics  []Metric

	ScrapeTime        prometheus.Gauge
	ErroneousRequests prometheus.Counter
	TotalRequests     prometheus.Counter
}

// New creates a new instance of a Collector.
func New(settings *config.Settings, sessions *sessions.Sessions) *Collector {
	return &Collector{
		Settings: settings,
		Sessions: sessions,
		Metrics:  Metrics,
		ScrapeTime: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "rds_exporter_scrape_duration_seconds",
			Help: "Time this RDS scrape took, in seconds.",
		}),
		ErroneousRequests: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "rds_exporter_erroneous_requests",
			Help: "The number of erroneous request made by this scrape.",
		}),
		TotalRequests: prometheus.NewCounter(prometheus.CounterOpts{
			Name: "rds_requests_total",
			Help: "API requests made to CloudWatch",
		}),
	}
}

func (c *Collector) Collect(ch chan<- prometheus.Metric) {
	now := time.Now()
	wg := &sync.WaitGroup{}
	for _, instance := range c.Settings.Config().Instances {
		wg.Add(1)
		go func(instance config.Instance) {
			scrape(instance, c, ch)
			wg.Done()
		}(instance)
	}
	wg.Wait()
	c.ScrapeTime.Set(time.Since(now).Seconds())

	ch <- c.ScrapeTime
	ch <- c.ErroneousRequests
	ch <- c.TotalRequests
}

func (c *Collector) Describe(ch chan<- *prometheus.Desc) {
	ch <- c.ScrapeTime.Desc()
	ch <- c.ErroneousRequests.Desc()
	ch <- c.TotalRequests.Desc()

	for _, m := range c.Metrics {
		ch <- m.Desc
	}
}
