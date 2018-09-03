package enhanced

// MapToNode maps RDS Enhanced Monitoring metrics to node_exporter metrics.
func MapToNode(namespace string, subsystem, name string, extraLabelsValues ...string) (namespaceOUT, subsystemOUT, nameOUT string, extraLabelsOUT []string, extraLabelsValuesOUT []string) {
	switch subsystem {
	case "cpuUtilization":
		// map cpuUtilization to node_cpu_average{cpu="All"}
		// we can't map it to node_exporter's node_cpu since it uses seconds, not percents

		// Turn metric name to 'mode' label e.g. node_cpu_average{cpu="All", mode="nice"}
		return "node", "", "cpu_average", []string{"mode"}, []string{name}
	case "loadAverageMinute":
		// map loadAverageMinute.one to node_load1
		switch name {
		case "one":
			return "node", "", "load1", nil, nil
		}
	case "memory":
		if name == "dirty" {
			return "node", "vmstat", "nr_dirty", nil, nil
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
		if nodeName, ok := names[name]; ok {
			return "node", "memory", nodeName, nil, nil
		}
	case "swap":
		names := map[string]string{
			"free":  "SwapFree",
			"total": "SwapTotal",
		}
		if nodeName, ok := names[name]; ok {
			return "node", "memory", nodeName, nil, nil
		}
		names = map[string]string{
			"in":  "pswpin",
			"out": "pswpout",
		}
		if nodeName, ok := names[name]; ok {
			return "node", "vmstat", nodeName, nil, nil
		}
	case "tasks":
		switch name {
		case "blocked",
			"running":
			return "node", "procs", name, nil, nil
		}
	case "fileSys":
		names := map[string]string{
			"avail": "avail",
			"total": "size",
		}
		if nodeName, ok := names[name]; ok {
			return "node", "filesystem", nodeName, []string{"id", "name", "mountPoint"}, extraLabelsValues
		}
		return namespace, subsystem, name, []string{"id", "name", "mountPoint"}, extraLabelsValues
	case "diskIO":
		names := map[string]string{
			"readKb":  "bytes_read",
			"writeKb": "bytes_written",
		}
		if nodeName, ok := names[name]; ok {
			return "node", "disk", nodeName, []string{"id", "device"}, extraLabelsValues
		}
		return namespace, subsystem, name, []string{"id", "device"}, extraLabelsValues
	case "processList":
		return namespace, subsystem, name, []string{"id", "name"}, extraLabelsValues
	case "network":
		return namespace, subsystem, name, []string{"id", "interface"}, extraLabelsValues
	}

	// If can't be mapped to node, then return original
	return namespace, subsystem, name, nil, nil
}
