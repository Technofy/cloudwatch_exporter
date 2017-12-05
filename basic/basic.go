package basic

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

type Exporter struct {
	Settings *config.Settings
	Sessions *sessions.Sessions
	Metrics  []Metric

	ScrapeTime        prometheus.Gauge
	ErroneousRequests prometheus.Counter
	TotalRequests     prometheus.Counter
}

// New creates a new instance of a Exporter.
func New(settings *config.Settings, sessions *sessions.Sessions) *Exporter {
	return &Exporter{
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

func (e *Exporter) Collect(ch chan<- prometheus.Metric) {
	now := time.Now()
	wg := &sync.WaitGroup{}
	instances := e.Settings.Config().Instances
	wg.Add(len(instances))
	for _, instance := range instances {
		go func(instance config.Instance) {
			NewScrape(instance, e, ch).Scrape()
			wg.Done()
		}(instance)
	}
	wg.Wait()
	e.ScrapeTime.Set(time.Since(now).Seconds())

	ch <- e.ScrapeTime
	ch <- e.ErroneousRequests
	ch <- e.TotalRequests
}

func (e *Exporter) Describe(ch chan<- *prometheus.Desc) {
	ch <- e.ScrapeTime.Desc()
	ch <- e.ErroneousRequests.Desc()
	ch <- e.TotalRequests.Desc()

	for _, m := range e.Metrics {
		ch <- m.Desc
	}
}
