package metrics

//go:generate go run generate/main.go

import (
	"fmt"

	"github.com/prometheus/client_golang/prometheus"
)

// Metric describes prometheus.Metric.
type Metric struct {
	Desc         *prometheus.Desc
	Unit         float64
	Namespace    string
	Subsystem    string
	Name         string
	Labels       []string
	LabelsValues []string
}

// Send value for metric m over chan<- prometheus.Metric.
func (m Metric) Send(ch chan<- prometheus.Metric, value float64) error {
	m = m.Node()
	FQName := m.BuildFQName()

	if len(m.Labels) != len(m.LabelsValues) {
		return fmt.Errorf("len(Labels) != len(LabelsValues) for metric '%s': len(%value) != len(%value): ", FQName, m.Labels, m.LabelsValues)
	}

	metric, ok := Default[FQName]
	if !ok {
		return fmt.Errorf("unknown metric %s", FQName)
	}

	ch <- prometheus.MustNewConstMetric(
		metric.Desc,
		prometheus.GaugeValue,
		value*metric.Unit,
		m.LabelsValues...,
	)
	return nil
}

// BuildFQName uses prometheus.BuildFQName to create metric name.
func (m Metric) BuildFQName() string {
	return prometheus.BuildFQName(m.Namespace, m.Subsystem, m.Name)
}

// Node transforms m to Metric compatible with node_exporter.
func (m Metric) Node() Metric {
	switch m.Subsystem {
	case "cpuUtilization":
		// map cpuUtilization to node_cpu_average{cpu="All"}
		// we can't map it to node_exporter's node_cpu since it uses seconds, not percents

		// Turn metric name to 'node' label e.g. node_cpu_average{cpu="All", mode="nice"}
		mode := m.Name
		m.Namespace = "node"
		m.Subsystem = ""
		m.Name = "cpu_average"
		m.Labels = append(m.Labels, "mode")
		m.LabelsValues = append(m.LabelsValues, mode)
	case "loadAverageMinute":
		// map loadAverageMinute.one to node_load1
		switch m.Name {
		case "one":
			m.Namespace = "node"
			m.Subsystem = ""
			m.Name = "load1"
		}
	case "memory":
		switch m.Name {
		case "dirty":
			m.Namespace = "node"
			m.Subsystem = "vmstat"
			m.Name = "nr_dirty"
		default:
			names := map[string]string{
				"buffers":    "Buffers",
				"cached":     "Cached",
				"free":       "MemFree",
				"total":      "MemTotal",
				"active":     "Active",
				"inactive":   "Inactive",
				"slab":       "Slab",
				"mapped":     "Mapped",
				"pageTables": "PageTables",
			}
			if nodeName, ok := names[m.Name]; ok {
				m.Namespace = "node"
				m.Subsystem = "memory"
				m.Name = nodeName
			}
		}
	case "swap":
		names := map[string]string{
			"free":  "SwapFree",
			"total": "SwapTotal",
		}
		if nodeName, ok := names[m.Name]; ok {
			m.Namespace = "node"
			m.Subsystem = "memory"
			m.Name = nodeName
			return m
		}
		names = map[string]string{
			"in":  "pswpin",
			"out": "pswpout",
		}
		if nodeName, ok := names[m.Name]; ok {
			m.Namespace = "node"
			m.Subsystem = "vmstat"
			m.Name = nodeName
			return m
		}
	case "tasks":
		switch m.Name {
		case "blocked",
			"running":
			m.Namespace = "node"
			m.Subsystem = "procs"
		}
	case "fileSys":
		m.Labels = append(m.Labels, "id", "name", "mountPoint")
		names := map[string]string{
			"avail": "avail",
			"total": "size",
		}
		if nodeName, ok := names[m.Name]; ok {
			m.Namespace = "node"
			m.Subsystem = "filesystem"
			m.Name = nodeName
		}
	case "diskIO":
		m.Labels = append(m.Labels, "id", "device")
		names := map[string]string{
			"readKb":  "bytes_read",
			"writeKb": "bytes_written",
		}
		if nodeName, ok := names[m.Name]; ok {
			m.Namespace = "node"
			m.Subsystem = "disk"
			m.Name = nodeName
		}
	case "processList":
		m.Labels = append(m.Labels, "id", "name")
	case "network":
		m.Labels = append(m.Labels, "id", "interface")
	}

	return m
}
