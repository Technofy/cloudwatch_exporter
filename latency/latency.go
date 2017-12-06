package latency

import (
	"sync"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

var (
	Desc = prometheus.NewDesc(
		"rds_latency",
		"The difference between the current time and timestamp in the metric itself",
		[]string{"instance", "region"},
		nil,
	)
)

type Latency struct {
	t time.Time
	sync.RWMutex
}

func (l *Latency) TakeOldest(t time.Time) {
	l.Lock()
	defer l.Unlock()

	if l.t.IsZero() || t.Before(l.t) {
		l.t = t
	}
}

func (l *Latency) Duration() time.Duration {
	l.RLock()
	defer l.RUnlock()

	return time.Since(l.t)
}

func (l *Latency) IsZero() bool {
	l.RLock()
	defer l.RUnlock()

	return l.t.IsZero()
}
