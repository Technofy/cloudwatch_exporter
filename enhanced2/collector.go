package enhanced2

import (
	"context"
	"sync"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/common/log"

	"github.com/percona/rds_exporter/sessions"
)

type Collector struct {
	sessions *sessions.Sessions
	logger   log.Logger

	rw      sync.RWMutex
	metrics map[string][]prometheus.Metric
}

func NewCollector(sessions *sessions.Sessions) *Collector {
	c := &Collector{
		sessions: sessions,
		logger:   log.With("component", "enhanced"),
		metrics:  make(map[string][]prometheus.Metric),
	}

	for session, instances := range sessions.AllSessions() {
		s := newScraper(session, instances)
		c.setMetrics(s.scrape())

		ch := make(chan map[string][]prometheus.Metric)
		go func() {
			for m := range ch {
				c.setMetrics(m)
			}
		}()
		go s.start(context.TODO(), ch)
	}

	return c
}

func (c *Collector) setMetrics(m map[string][]prometheus.Metric) {
	c.rw.Lock()
	for id, metrics := range m {
		c.metrics[id] = metrics
	}
	c.rw.Unlock()
}

func (c *Collector) Describe(ch chan<- *prometheus.Desc) {
	c.rw.RLock()
	defer c.rw.RUnlock()

	for _, metrics := range c.metrics {
		for _, m := range metrics {
			ch <- m.Desc()
		}
	}
}

func (c *Collector) Collect(ch chan<- prometheus.Metric) {
	c.rw.RLock()
	defer c.rw.RUnlock()

	for _, metrics := range c.metrics {
		for _, m := range metrics {
			ch <- m
		}
	}
}

// check interfaces
var (
	_ prometheus.Collector = (*Collector)(nil)
)
