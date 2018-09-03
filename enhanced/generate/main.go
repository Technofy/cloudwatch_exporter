// The following directive is necessary to make the package coherent:
// This program generates metrics.go. It can be invoked by running
// go generate
package main

import (
	"log"
	"os"
	"sort"
	"strings"
	"text/template"

	"github.com/prometheus/client_golang/prometheus"

	"github.com/percona/rds_exporter/enhanced"
)

type Group struct {
	Name    string
	metrics []Metric
}

type Metric struct {
	Group string
	Name  string
	Help  string
}

func (m Metric) FqName() string {
	namespace, subsystem, name, _, _ := enhanced.MapToNode("rdsosmetrics", m.Group, m.Name)
	return prometheus.BuildFQName(namespace, subsystem, name)
}

// ParsedName returns RDS Enhanced Monitoring metric name mapped to rds_exporter name.
func (m Metric) ParsedName() string {
	_, _, name, _, _ := enhanced.MapToNode("rdsosmetrics", m.Group, m.Name)
	return name
}

func (m Metric) Labels() []string {
	labels := []string{
		"instance",
		"region",
	}
	_, _, _, extraLabels, _ := enhanced.MapToNode("rdsosmetrics", m.Group, m.Name)
	return append(labels, extraLabels...)
}

func (m Metric) ConstLabels() map[string]string {
	switch m.Group {
	case "cpuUtilization":
		return map[string]string{
			"cpu": "All",
		}
	}

	return nil
}

func (m Metric) Unit() float64 {
	if strings.Contains(m.Help, "kilobytes") {
		if strings.HasPrefix(m.FqName(), "node_") {
			excluded := map[string]struct{}{
				"in":    {},
				"out":   {},
				"dirty": {},
			}
			if _, ok := excluded[m.Name]; !ok {
				return 1024
			}
		}
	}
	return 1
}

func (m Metric) ParsedHelp() string {
	if m.Unit() == 1024 {
		return strings.Replace(m.Help, "kilobytes", "bytes", 1)
	}
	return m.Help
}

func (g Group) Metrics() []Metric {
	switch g.Name {
	case "cpuUtilization":
		return []Metric{
			{
				Group: g.Name,
				Name:  "cpu_average",
				Help:  "The percentage of CPU utilization. Units: Percent",
			},
		}
	}
	return g.metrics
}

var (
	// http://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_Monitoring.OS.html
	docs = map[string]map[string]string{
		"General": {
			"engine":             "The database engine for the DB instance.",
			"instanceID":         "The DB instance identifier.",
			"instanceResourceID": "A region-unique, immutable identifier for the DB instance, also used as the log stream identifier.",
			"numVCPUs":           "The number of virtual CPUs for the DB instance.",
			"timestamp":          "The time at which the metrics were taken.",
			"uptime":             "The amount of time that the DB instance has been active.",
			"version":            "The version of the OS metrics' stream JSON format.",
		},
		"cpuUtilization": {
			"guest":  "The percentage of CPU in use by guest programs.",
			"idle":   "The percentage of CPU that is idle.",
			"irq":    "The percentage of CPU in use by software interrupts.",
			"nice":   "The percentage of CPU in use by programs running at lowest priority.",
			"steal":  "The percentage of CPU in use by other virtual machines.",
			"system": "The percentage of CPU in use by the kernel.",
			"total":  "The total percentage of the CPU in use. This value includes the nice value.",
			"user":   "The percentage of CPU in use by user programs.",
			"wait":   "The percentage of CPU unused while waiting for I/O access.",
		},
		"diskIO": {
			"avgQueueLen": "The number of requests waiting in the I/O device's queue. This metric is not available for Amazon Aurora.",
			"avgReqSz":    "The average request size, in kilobytes. This metric is not available for Amazon Aurora.",
			"await":       "The number of milliseconds required to respond to requests, including queue time and service time. This metric is not available for Amazon Aurora.",
			// "device" is mapped to label
			// "device":          "The identifier of the disk device in use. This metric is not available for Amazon Aurora.",
			"readIOsPS":       "The number of read operations per second. This metric is not available for Amazon Aurora.",
			"readKb":          "The total number of kilobytes read. This metric is not available for Amazon Aurora.",
			"readKbPS":        "The number of kilobytes read per second. This metric is not available for Amazon Aurora.",
			"rrqmPS":          "The number of merged read requests queued per second. This metric is not available for Amazon Aurora.",
			"tps":             "The number of I/O transactions per second. This metric is not available for Amazon Aurora.",
			"util":            "The percentage of CPU time during which requests were issued. This metric is not available for Amazon Aurora.",
			"writeIOsPS":      "The number of write operations per second. This metric is not available for Amazon Aurora.",
			"writeKb":         "The total number of kilobytes written. This metric is not available for Amazon Aurora.",
			"writeKbPS":       "The number of kilobytes written per second. This metric is not available for Amazon Aurora.",
			"wrqmPS":          "The number of merged write requests queued per second. This metric is not available for Amazon Aurora.",
			"readLatency":     "The average amount of time taken per disk I/O operation.",
			"writeLatency":    "The average amount of time taken per disk I/O operation.",
			"writeThroughput": "The average number of bytes written to disk per second.",
			"readThroughput":  "The average number of bytes read from disk per second.",
			"diskQueueDepth":  "The number of outstanding read/write requests waiting to access the disk.",
		},
		"fileSys": {
			"maxFiles": "The maximum number of files that can be created for the file system.",
			// "mountPoint" is mapped to label
			// "mountPoint":      "The path to the file system.",
			// "name" is mapped to label
			// "name":            "The name of the file system.",
			"total": "The total number of disk space available for the file system, in kilobytes.",
			"used":  "The amount of disk space used by files in the file system, in kilobytes.",
			// avail = total - used
			"avail":           "The amount of disk space available in the file system, in kilobytes.",
			"usedFilePercent": "The percentage of available files in use.",
			"usedFiles":       "The number of files in the file system.",
			"usedPercent":     "The percentage of the file-system disk space in use.",
		},
		"loadAverageMinute": {
			"fifteen": "The number of processes requesting CPU time over the last 15 minutes.",
			"five":    "The number of processes requesting CPU time over the last 5 minutes.",
			"one":     "The number of processes requesting CPU time over the last minute.",
		},
		"memory": {
			"active":         "The amount of assigned memory, in kilobytes.",
			"buffers":        "The amount of memory used for buffering I/O requests prior to writing to the storage device, in kilobytes.",
			"cached":         "The amount of memory used for caching file system–based I/O, in kilobytes.",
			"dirty":          "The amount of memory pages in RAM that have been modified but not written to their related data block in storage, in kilobytes.",
			"free":           "The amount of unassigned memory, in kilobytes.",
			"hugePagesFree":  "The number of free huge pages.Huge pages are a feature of the Linux kernel.",
			"hugePagesRsvd":  "The number of committed huge pages.",
			"hugePagesSize":  "The size for each huge pages unit, in kilobytes.",
			"hugePagesSurp":  "The number of available surplus huge pages over the total.",
			"hugePagesTotal": "The total number of huge pages for the system.",
			"inactive":       "The amount of least-frequently used memory pages, in kilobytes.",
			"mapped":         "The total amount of file-system contents that is memory mapped inside a process address space, in kilobytes.",
			"pageTables":     "The amount of memory used by page tables, in kilobytes.",
			"slab":           "The amount of reusable kernel data structures, in kilobytes.",
			"total":          "The total amount of memory, in kilobytes.",
			"writeback":      "The amount of dirty pages in RAM that are still being written to the backing storage, in kilobytes.",
		},
		"network": {
			// "interface" is mapped to label
			// "interface": "The identifier for the network interface being used for the DB instance.",
			"rx": "The number of bytes received per second.",
			"tx": "The number of bytes uploaded per second.",
		},
		"processList": {
			"cpuUsedPc":    "The percentage of CPU used by the process.",
			"id":           "The identifier of the process.",
			"memoryUsedPc": "The amount of memory used by the process, in kilobytes.",
			// "name" is mapped to label
			// "name":         "The name of the process.",
			"parentID": "The process identifier for the parent process of the process.",
			"rss":      "The amount of RAM allocated to the process, in kilobytes.",
			"tgid":     "The thread group identifier, which is a number representing the process ID to which a thread belongs.This identifier is used to group threads from the same process.",
			"vss":      "The amount of virtual memory allocated to the process, in kilobytes.",
			"vmlimit":  "TODO",
		},
		"swap": {
			"cached": "The amount of swap memory, in kilobytes, used as cache memory.",
			"free":   "The total amount of swap memory free, in kilobytes.",
			"total":  "The total amount of swap memory available, in kilobytes.",
			"in":     "Number of kilobytes the system has swapped in from disk per second (disk reads).",
			"out":    "Number of kilobytes the system has swapped out to disk per second (disk writes).",
		},
		"tasks": {
			"blocked":  "The number of tasks that are blocked.",
			"running":  "The number of tasks that are running.",
			"sleeping": "The number of tasks that are sleeping.",
			"stopped":  "The number of tasks that are stopped.",
			"total":    "The total number of tasks.",
			"zombie":   "The number of child tasks that are inactive with an active parent task.",
		},
	}
)

func main() {
	f, err := os.Create("metrics.go")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	groups := []Group{}
	for groupName, doc := range docs {
		metrics := []Metric{}
		for metricName, metricHelp := range doc {
			metric := Metric{
				Group: groupName,
				Name:  metricName,
				Help:  metricHelp,
			}
			metrics = append(metrics, metric)
		}
		sort.SliceStable(metrics, func(i, j int) bool {
			return metrics[i].Name < metrics[j].Name
		})
		group := Group{
			Name:    groupName,
			metrics: metrics,
		}
		groups = append(groups, group)
	}
	sort.SliceStable(groups, func(i, j int) bool {
		return groups[i].Name < groups[j].Name
	})
	packageTemplate.Execute(f, struct {
		Groups []Group
	}{
		Groups: groups,
	})
}

var packageTemplate = template.Must(template.New("").Parse(`// Code generated by go generate; DO NOT EDIT.
package enhanced

import (
	"github.com/prometheus/client_golang/prometheus"
)

var Metrics = map[string]Metric{
{{- range .Groups }}
{{- range .Metrics }}
	"{{.FqName}}" : {
		Name: "{{.ParsedName}}",
		Desc: prometheus.NewDesc(
			"{{.FqName}}",
			"{{.ParsedHelp}}",
			{{printf "%#v" .Labels}},
			{{printf "%#v" .ConstLabels}},
		),
        Unit: {{.Unit}},
	},
{{- end }}
{{- end }}
}
`))
