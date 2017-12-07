package enhanced

func MapToNode(subsystem, name string, extraLabels ...string) (namespaceOUT, subsystemOUT, nameOUT string, labelsOUT []string) {
	switch subsystem {
	case "cpuUtilization":
		// map cpuUtilization to node_cpu{cpu="All"}

		// Turn metric name to 'mode' label e.g. node_cpu{cpu="All", mode="nice"}
		return "node", "", "cpu", append(extraLabels, name)
	case "loadAverageMinute":
		// map loadAverageMinute.one to node_load1
		switch name {
		case "one":
			return "node", "", "load1", nil
		}
	case "memory":
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
			"dirty":      "nr_dirty",
		}
		if nodeName, ok := names[name]; ok {
			return "node", "memory", nodeName, nil
		}
	case "swap":
		names := map[string]string{
			"free":  "SwapFree",
			"total": "SwapTotal",
		}
		if nodeName, ok := names[name]; ok {
			return "node", "memory", nodeName, nil
		}
	case "tasks":
		switch name {
		case "blocked",
			"running":
			return "node", "procs", name, nil
		}
	case "fileSys":
		names := map[string]string{
			"name":  "avail",
			"total": "size",
		}
		if nodeName, ok := names[name]; ok {
			return "node", "filesystem", nodeName, nil
		}
	case "diskIO":
		// Only first device is converted to node name
		if len(extraLabels) != 1 {
			break
		}
		if extraLabels[0] != "0" {
			break
		}

		names := map[string]string{
			"readKb":  "bytes_read",
			"writeKb": "bytes_written",
		}
		if nodeName, ok := names[name]; ok {
			return "node", "disk", nodeName, nil
		}
	}

	// If can't be mapped to node, then return original
	return defaultNamespace, subsystem, name, extraLabels
}
