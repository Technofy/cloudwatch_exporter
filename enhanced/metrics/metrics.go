package metrics

import "github.com/prometheus/client_golang/prometheus"

type Metrics map[string]Metric

type Metric struct {
	Desc *prometheus.Desc
	Unit float64
}
