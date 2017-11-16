package main

import (
	"sync"
	"time"

	"github.com/prometheus/client_golang/prometheus"

	"github.com/percona/rds_exporter/config"
)

var (
	metrics = []string{
		"ActiveTransactions",
		"AuroraBinlogReplicaLag",
		"AuroraReplicaLag",
		"AuroraReplicaLagMaximum",
		"AuroraReplicaLagMinimum",
		"BinLogDiskUsage",
		"BlockedTransactions",
		"BufferCacheHitRatio",
		"BurstBalance",
		"CommitLatency",
		"CommitThroughput",
		"CPUCreditBalance",
		"CPUCreditUsage",
		"CPUUtilization",
		"DatabaseConnections",
		"DDLLatency",
		"DDLThroughput",
		"Deadlocks",
		"DeleteLatency",
		"DeleteThroughput",
		"DiskQueueDepth",
		"DMLLatency",
		"DMLThroughput",
		"EngineUptime",
		"FreeableMemory",
		"FreeLocalStorage",
		"FreeStorageSpace",
		"InsertLatency",
		"InsertThroughput",
		"LoginFailures",
		"NetworkReceiveThroughput",
		"NetworkThroughput",
		"NetworkTransmitThroughput",
		"Queries",
		"ReadIOPS",
		"ReadLatency",
		"ReadThroughput",
		"ResultSetCacheHitRatio",
		"SelectLatency",
		"SelectThroughput",
		"SwapUsage",
		"UpdateLatency",
		"UpdateThroughput",
		"VolumeBytesUsed",
		"VolumeReadIOPs",
		"VolumeWriteIOPs",
		"WriteIOPS",
		"WriteLatency",
		"WriteThroughput",
	}
)

type Metric struct {
	Name string
	Desc *prometheus.Desc
}

type Collector struct {
	Settings *config.Settings
	Metrics  []Metric

	ScrapeTime        prometheus.Gauge
	ErroneousRequests prometheus.Counter
	TotalRequests     prometheus.Counter
}

// New creates a new instance of a Collector.
func New(settings *config.Settings) *Collector {
	return &Collector{
		Settings: settings,
		Metrics:  generateMetrics(),
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

func generateMetrics() []Metric {
	ms := make([]Metric, len(metrics))
	for i, name := range metrics {
		ms[i].Name = name
		ms[i].Desc = prometheus.NewDesc(
			safeName("AWS/RDS_"+toSnakeCase(name)+"_average"),
			name,
			[]string{
				"instance",
				"region",
			},
			nil,
		)
	}

	return ms
}

func (c *Collector) Collect(ch chan<- prometheus.Metric) {
	now := time.Now()
	wg := &sync.WaitGroup{}
	for _, instance := range c.Settings.Config().Instances {
		wg.Add(1)
		go func(instance, region string) {
			scrape(instance, region, c, ch)
			wg.Done()
		}(instance.Instance, instance.Region)
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
