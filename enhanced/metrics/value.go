package metrics

//go:generate go run generate/main.go

import (
	"fmt"

	"github.com/prometheus/client_golang/prometheus"
)

type Value struct {
	Namespace    string
	Subsystem    string
	Name         string
	Labels       []string
	LabelsValues []string
}

func (v Value) Send(ch chan<- prometheus.Metric, value float64) error {
	v = v.Node()
	FQName := v.BuildFQName()

	if len(v.Labels) != len(v.LabelsValues) {
		return fmt.Errorf("len(Labels) != len(LabelsValues) for metric '%s': len(%value) != len(%value): ", FQName, v.Labels, v.LabelsValues)
	}

	metric, ok := Default[FQName]
	if !ok {
		return fmt.Errorf("unknown metric %s", FQName)
	}

	ch <- prometheus.MustNewConstMetric(
		metric.Desc,
		prometheus.GaugeValue,
		value*metric.Unit,
		v.LabelsValues...,
	)
	return nil
}

func (v Value) BuildFQName() string {
	return prometheus.BuildFQName(v.Namespace, v.Subsystem, v.Name)
}

func (v Value) Node() Value {
	switch v.Subsystem {
	case "cpuUtilization":
		// map cpuUtilization to node_cpu_average{cpu="All"}
		// we can't map it to node_exporter's node_cpu since it uses seconds, not percents

		// Turn metric name to 'node' label e.g. node_cpu_average{cpu="All", mode="nice"}
		return Value{"node", "", "cpu_average", append(v.Labels, "mode"), append(v.LabelsValues, v.Name)}
	case "loadAverageMinute":
		// map loadAverageMinute.one to node_load1
		switch v.Name {
		case "one":
			return Value{"node", "", "load1", v.Labels, v.LabelsValues}
		}
	case "memory":
		if v.Name == "dirty" {
			return Value{"node", "vmstat", "nr_dirty", v.Labels, v.LabelsValues}
		}
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
		if nodeName, ok := names[v.Name]; ok {
			return Value{"node", "memory", nodeName, v.Labels, v.LabelsValues}
		}
	case "swap":
		names := map[string]string{
			"free":  "SwapFree",
			"total": "SwapTotal",
		}
		if nodeName, ok := names[v.Name]; ok {
			return Value{"node", "memory", nodeName, v.Labels, v.LabelsValues}
		}
		names = map[string]string{
			"in":  "pswpin",
			"out": "pswpout",
		}
		if nodeName, ok := names[v.Name]; ok {
			return Value{"node", "vmstat", nodeName, v.Labels, v.LabelsValues}
		}
	case "tasks":
		switch v.Name {
		case "blocked",
			"running":
			return Value{"node", "procs", v.Name, v.Labels, v.LabelsValues}
		}
	case "fileSys":
		labels := append(v.Labels, "id", "name", "mountPoint")
		names := map[string]string{
			"avail": "avail",
			"total": "size",
		}
		if nodeName, ok := names[v.Name]; ok {
			return Value{"node", "filesystem", nodeName, labels, v.LabelsValues}
		}
		return Value{v.Namespace, v.Subsystem, v.Name, labels, v.LabelsValues}
	case "diskIO":
		names := map[string]string{
			"readKb":  "bytes_read",
			"writeKb": "bytes_written",
		}
		if nodeName, ok := names[v.Name]; ok {
			return Value{"node", "disk", nodeName, append(v.Labels, "id", "device"), v.LabelsValues}
		}
		return Value{v.Namespace, v.Subsystem, v.Name, append(v.Labels, "id", "device"), v.LabelsValues}
	case "processList":
		return Value{v.Namespace, v.Subsystem, v.Name, append(v.Labels, "id", "name"), v.LabelsValues}
	case "network":
		return Value{v.Namespace, v.Subsystem, v.Name, append(v.Labels, "id", "interface"), v.LabelsValues}
	}

	// If can't be mapped to node, then return original
	return v
}
