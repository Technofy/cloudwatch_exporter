package client

import (
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/common/log"
)

type transport struct {
	t *http.Transport
	l log.Logger

	mRequests prometheus.Counter
}

func newTransport() *transport {
	return &transport{
		t: &http.Transport{
			MaxIdleConnsPerHost: 5,
			IdleConnTimeout:     2 * time.Minute,
		},
		l: log.With("component", "transport"),

		mRequests: prometheus.NewCounter(prometheus.CounterOpts{
			Name: "rds_exporter_requests_total",
			Help: "Number of API requests to AWS.",
		}),
	}
}

func (t *transport) RoundTrip(req *http.Request) (*http.Response, error) {
	// We could use "net/http/httptrace" package if we ever need more metrics.

	start := time.Now()
	resp, err := t.t.RoundTrip(req)
	t.mRequests.Inc()
	if t.l != nil {
		t.l.Debugf("%s %s -> %d %v (%s)", req.Method, req.URL.String(), resp.StatusCode, err, time.Since(start))
	}
	return resp, err
}

// check interface
var _ http.RoundTripper = (*transport)(nil)
