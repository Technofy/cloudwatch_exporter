package basic

import (
	"testing"

	"github.com/prometheus/client_golang/prometheus"

	"github.com/percona/rds_exporter/config"
	"github.com/percona/rds_exporter/sessions"
)

func TestNew(t *testing.T) {
	t.Parallel()

	settings := &config.Settings{}
	sessions := &sessions.Sessions{}
	c := New(settings, sessions)

	if c == nil {
		t.Fatal("exporter should not be nil")
	}
}

func TestCollector_Collect(t *testing.T) {
	t.Parallel()

	settings := &config.Settings{}
	sessions := &sessions.Sessions{}
	c := New(settings, sessions)

	if c == nil {
		t.Fatal("exporter should not be nil")
	}
	ch := make(chan prometheus.Metric, 1000)
	c.Collect(ch)

	metrics := []prometheus.Metric{}

	for {
		select {
		case m := <-ch:
			metrics = append(metrics, m)
			continue
		default:
		}
		break
	}

	if len(metrics) == 0 {
		t.Fatal("Collect() didn't collect any metrics")
	}
}

func TestCollector_Describe(t *testing.T) {
	t.Parallel()

	settings := &config.Settings{}
	sessions := &sessions.Sessions{}
	c := New(settings, sessions)

	if c == nil {
		t.Fatal("exporter should not be nil")
	}
	ch := make(chan *prometheus.Desc, 1000)
	c.Describe(ch)

	desc := []*prometheus.Desc{}

	for {
		select {
		case m := <-ch:
			desc = append(desc, m)
			continue
		default:
		}
		break
	}

	if len(desc) == 0 {
		t.Fatal("Describe() didn't describe any metrics")
	}
}
