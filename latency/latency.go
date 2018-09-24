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
	t  time.Time
	rw sync.RWMutex
}

func (l *Latency) TakeOldest(t time.Time) {
	l.rw.Lock()
	defer l.rw.Unlock()

	if l.t.IsZero() || t.Before(l.t) {
		l.t = t
	}
}

func (l *Latency) Duration() time.Duration {
	l.rw.RLock()
	defer l.rw.RUnlock()

	return time.Since(l.t)
}

func (l *Latency) IsZero() bool {
	l.rw.RLock()
	defer l.rw.RUnlock()

	return l.t.IsZero()
}
