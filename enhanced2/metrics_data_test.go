package enhanced2

import (
	"strings"
)

var (
	dataMySQL57 = []byte(strings.TrimSpace(`
	{
		"engine": "MYSQL",
		"instanceID": "rds-mysql57",
		"instanceResourceID": "db-EXAMPLE",
		"timestamp": "2018-01-08T14:35:21Z",
		"version": 1.00,
		"uptime": "39 days, 5:08:13",
		"numVCPUs": 1,
		"cpuUtilization": {
			"guest": 0.00,
			"irq": 1.16,
			"system": 4.71,
			"wait": 81.43,
			"idle": 0.00,
			"user": 1.18,
			"total": 99.99,
			"steal": 1.01,
			"nice": 10.50
		},
		"loadAverageMinute": {
			"fifteen": 7.64,
			"five": 14.08,
			"one": 32.83
		},
		"memory": {
			"writeback": 0,
			"hugePagesFree": 0,
			"hugePagesRsvd": 0,
			"hugePagesSurp": 0,
			"cached": 162452,
			"hugePagesSize": 2048,
			"free": 127048,
			"hugePagesTotal": 0,
			"inactive": 511232,
			"pageTables": 7916,
			"dirty": 360,
			"mapped": 14184,
			"active": 1299668,
			"total": 2051520,
			"slab": 45440,
			"buffers": 149412
		},
		"tasks": {
			"sleeping": 258,
			"zombie": 0,
			"running": 1,
			"stopped": 0,
			"total": 313,
			"blocked": 54
		},
		"swap": {
			"cached": 908,
			"total": 4095996,
			"out": 0.00,
			"free": 3803212,
			"in": 0.00
		},
		"network": [
			{
				"interface": "eth0",
				"rx": 93611.95,
				"tx": 801373.55
			}
		],
		"diskIO": [
			{
				"writeKbPS": 6157.27,
				"readIOsPS": 354.18,
				"await": 74.67,
				"readKbPS": 5664.93,
				"rrqmPS": 0.00,
				"util": 99.94,
				"avgQueueLen": 2783.31,
				"tps": 623.33,
				"readKb": 339896,
				"device": "rdsdev",
				"writeKb": 369436,
				"avgReqSz": 18.97,
				"wrqmPS": 0.00,
				"writeIOsPS": 269.15
			}
		],
		"fileSys": [
			{
				"used": 9659328,
				"name": "rdsfilesys",
				"usedFiles": 731,
				"usedFilePercent": 0.02,
				"maxFiles": 3932160,
				"mountPoint": "/rdsdbdata",
				"total": 61774768,
				"usedPercent": 15.64
			}
		],
		"processList": [
			{
				"vss": 2552316,
				"name": "mysqld",
				"tgid": 4516,
				"vmlimit": "unlimited",
				"parentID": 1,
				"memoryUsedPc": 62.51,
				"cpuUsedPc": 0.13,
				"id": 1594,
				"rss": 1282364
			},
			{
				"vss": 2552316,
				"name": "mysqld",
				"tgid": 4516,
				"vmlimit": "unlimited",
				"parentID": 1,
				"memoryUsedPc": 62.51,
				"cpuUsedPc": 0.12,
				"id": 1595,
				"rss": 1282364
			},
			{
				"vss": 2552316,
				"name": "mysqld",
				"tgid": 4516,
				"vmlimit": "unlimited",
				"parentID": 1,
				"memoryUsedPc": 62.51,
				"cpuUsedPc": 0.12,
				"id": 1596,
				"rss": 1282364
			},
			{
				"vss": 2552316,
				"name": "mysqld",
				"tgid": 4516,
				"vmlimit": "unlimited",
				"parentID": 1,
				"memoryUsedPc": 62.51,
				"cpuUsedPc": 0.13,
				"id": 1597,
				"rss": 1282364
			},
			{
				"vss": 2552316,
				"name": "mysqld",
				"tgid": 4516,
				"vmlimit": "unlimited",
				"parentID": 1,
				"memoryUsedPc": 62.51,
				"cpuUsedPc": 0.13,
				"id": 1599,
				"rss": 1282364
			},
			{
				"vss": 2552316,
				"name": "mysqld",
				"tgid": 4516,
				"vmlimit": "unlimited",
				"parentID": 1,
				"memoryUsedPc": 62.51,
				"cpuUsedPc": 0.12,
				"id": 1600,
				"rss": 1282364
			},
			{
				"vss": 2552316,
				"name": "mysqld",
				"tgid": 4516,
				"vmlimit": "unlimited",
				"parentID": 1,
				"memoryUsedPc": 62.51,
				"cpuUsedPc": 0.13,
				"id": 1601,
				"rss": 1282364
			},
			{
				"vss": 2552316,
				"name": "mysqld",
				"tgid": 4516,
				"vmlimit": "unlimited",
				"parentID": 1,
				"memoryUsedPc": 62.51,
				"cpuUsedPc": 0.13,
				"id": 1602,
				"rss": 1282364
			},
			{
				"vss": 2552316,
				"name": "mysqld",
				"tgid": 4516,
				"vmlimit": "unlimited",
				"parentID": 1,
				"memoryUsedPc": 62.51,
				"cpuUsedPc": 0.10,
				"id": 1603,
				"rss": 1282364
			},
			{
				"vss": 2552316,
				"name": "mysqld",
				"tgid": 4516,
				"vmlimit": "unlimited",
				"parentID": 1,
				"memoryUsedPc": 62.51,
				"cpuUsedPc": 0.12,
				"id": 1604,
				"rss": 1282364
			},
			{
				"vss": 2552316,
				"name": "mysqld",
				"tgid": 4516,
				"vmlimit": "unlimited",
				"parentID": 1,
				"memoryUsedPc": 62.51,
				"cpuUsedPc": 0.13,
				"id": 1605,
				"rss": 1282364
			},
			{
				"vss": 2552316,
				"name": "mysqld",
				"tgid": 4516,
				"vmlimit": "unlimited",
				"parentID": 1,
				"memoryUsedPc": 62.51,
				"cpuUsedPc": 0.13,
				"id": 1606,
				"rss": 1282364
			},
			{
				"vss": 2552316,
				"name": "mysqld",
				"tgid": 4516,
				"vmlimit": "unlimited",
				"parentID": 1,
				"memoryUsedPc": 62.51,
				"cpuUsedPc": 0.13,
				"id": 1607,
				"rss": 1282364
			},
			{
				"vss": 2552316,
				"name": "mysqld",
				"tgid": 4516,
				"vmlimit": "unlimited",
				"parentID": 1,
				"memoryUsedPc": 62.51,
				"cpuUsedPc": 0.13,
				"id": 1608,
				"rss": 1282364
			},
			{
				"vss": 2552316,
				"name": "mysqld",
				"tgid": 4516,
				"vmlimit": "unlimited",
				"parentID": 1,
				"memoryUsedPc": 62.51,
				"cpuUsedPc": 0.13,
				"id": 1609,
				"rss": 1282364
			},
			{
				"vss": 2552316,
				"name": "mysqld",
				"tgid": 4516,
				"vmlimit": "unlimited",
				"parentID": 1,
				"memoryUsedPc": 62.51,
				"cpuUsedPc": 0.12,
				"id": 1610,
				"rss": 1282364
			},
			{
				"vss": 2552316,
				"name": "mysqld",
				"tgid": 4516,
				"vmlimit": "unlimited",
				"parentID": 1,
				"memoryUsedPc": 62.51,
				"cpuUsedPc": 0.13,
				"id": 1611,
				"rss": 1282364
			},
			{
				"vss": 2552316,
				"name": "mysqld",
				"tgid": 4516,
				"vmlimit": "unlimited",
				"parentID": 1,
				"memoryUsedPc": 62.51,
				"cpuUsedPc": 0.13,
				"id": 1612,
				"rss": 1282364
			},
			{
				"vss": 2552316,
				"name": "mysqld",
				"tgid": 4516,
				"vmlimit": "unlimited",
				"parentID": 1,
				"memoryUsedPc": 62.51,
				"cpuUsedPc": 0.12,
				"id": 1613,
				"rss": 1282364
			},
			{
				"vss": 2552316,
				"name": "mysqld",
				"tgid": 4516,
				"vmlimit": "unlimited",
				"parentID": 1,
				"memoryUsedPc": 62.51,
				"cpuUsedPc": 0.13,
				"id": 1614,
				"rss": 1282364
			},
			{
				"vss": 2552316,
				"name": "mysqld",
				"tgid": 4516,
				"vmlimit": "unlimited",
				"parentID": 1,
				"memoryUsedPc": 62.51,
				"cpuUsedPc": 0.12,
				"id": 1615,
				"rss": 1282364
			},
			{
				"vss": 2552316,
				"name": "mysqld",
				"tgid": 4516,
				"vmlimit": "unlimited",
				"parentID": 1,
				"memoryUsedPc": 62.51,
				"cpuUsedPc": 0.13,
				"id": 1616,
				"rss": 1282364
			},
			{
				"vss": 2552316,
				"name": "mysqld",
				"tgid": 4516,
				"vmlimit": "unlimited",
				"parentID": 1,
				"memoryUsedPc": 62.51,
				"cpuUsedPc": 0.13,
				"id": 1617,
				"rss": 1282364
			},
			{
				"vss": 2552316,
				"name": "mysqld",
				"tgid": 4516,
				"vmlimit": "unlimited",
				"parentID": 1,
				"memoryUsedPc": 62.51,
				"cpuUsedPc": 0.13,
				"id": 1618,
				"rss": 1282364
			},
			{
				"vss": 2552316,
				"name": "mysqld",
				"tgid": 4516,
				"vmlimit": "unlimited",
				"parentID": 1,
				"memoryUsedPc": 62.51,
				"cpuUsedPc": 0.12,
				"id": 1619,
				"rss": 1282364
			},
			{
				"vss": 2552316,
				"name": "mysqld",
				"tgid": 4516,
				"vmlimit": "unlimited",
				"parentID": 1,
				"memoryUsedPc": 62.51,
				"cpuUsedPc": 0.13,
				"id": 1620,
				"rss": 1282364
			},
			{
				"vss": 2552316,
				"name": "mysqld",
				"tgid": 4516,
				"vmlimit": "unlimited",
				"parentID": 1,
				"memoryUsedPc": 62.51,
				"cpuUsedPc": 0.13,
				"id": 1621,
				"rss": 1282364
			},
			{
				"vss": 2552316,
				"name": "mysqld",
				"tgid": 4516,
				"vmlimit": "unlimited",
				"parentID": 1,
				"memoryUsedPc": 62.51,
				"cpuUsedPc": 0.12,
				"id": 1622,
				"rss": 1282364
			},
			{
				"vss": 2552316,
				"name": "mysqld",
				"tgid": 4516,
				"vmlimit": "unlimited",
				"parentID": 1,
				"memoryUsedPc": 62.51,
				"cpuUsedPc": 0.12,
				"id": 1623,
				"rss": 1282364
			},
			{
				"vss": 2552316,
				"name": "mysqld",
				"tgid": 4516,
				"vmlimit": "unlimited",
				"parentID": 1,
				"memoryUsedPc": 62.51,
				"cpuUsedPc": 0.12,
				"id": 1624,
				"rss": 1282364
			},
			{
				"vss": 2552316,
				"name": "mysqld",
				"tgid": 4516,
				"vmlimit": "unlimited",
				"parentID": 1,
				"memoryUsedPc": 62.51,
				"cpuUsedPc": 0.13,
				"id": 1625,
				"rss": 1282364
			},
			{
				"vss": 2552316,
				"name": "mysqld",
				"tgid": 4516,
				"vmlimit": "unlimited",
				"parentID": 1,
				"memoryUsedPc": 62.51,
				"cpuUsedPc": 0.15,
				"id": 1626,
				"rss": 1282364
			},
			{
				"vss": 2552316,
				"name": "mysqld",
				"tgid": 4516,
				"vmlimit": "unlimited",
				"parentID": 1,
				"memoryUsedPc": 62.51,
				"cpuUsedPc": 0.10,
				"id": 1627,
				"rss": 1282364
			},
			{
				"vss": 2552316,
				"name": "mysqld",
				"tgid": 4516,
				"vmlimit": "unlimited",
				"parentID": 1,
				"memoryUsedPc": 62.51,
				"cpuUsedPc": 0.13,
				"id": 1628,
				"rss": 1282364
			},
			{
				"vss": 2552316,
				"name": "mysqld",
				"tgid": 4516,
				"vmlimit": "unlimited",
				"parentID": 1,
				"memoryUsedPc": 62.51,
				"cpuUsedPc": 0.12,
				"id": 1629,
				"rss": 1282364
			},
			{
				"vss": 2552316,
				"name": "mysqld",
				"tgid": 4516,
				"vmlimit": "unlimited",
				"parentID": 1,
				"memoryUsedPc": 62.51,
				"cpuUsedPc": 0.12,
				"id": 1630,
				"rss": 1282364
			},
			{
				"vss": 2552316,
				"name": "mysqld",
				"tgid": 4516,
				"vmlimit": "unlimited",
				"parentID": 1,
				"memoryUsedPc": 62.51,
				"cpuUsedPc": 0.13,
				"id": 1631,
				"rss": 1282364
			},
			{
				"vss": 2552316,
				"name": "mysqld",
				"tgid": 4516,
				"vmlimit": "unlimited",
				"parentID": 1,
				"memoryUsedPc": 62.51,
				"cpuUsedPc": 0.13,
				"id": 1632,
				"rss": 1282364
			},
			{
				"vss": 2552316,
				"name": "mysqld",
				"tgid": 4516,
				"vmlimit": "unlimited",
				"parentID": 1,
				"memoryUsedPc": 62.51,
				"cpuUsedPc": 0.13,
				"id": 1633,
				"rss": 1282364
			},
			{
				"vss": 2552316,
				"name": "mysqld",
				"tgid": 4516,
				"vmlimit": "unlimited",
				"parentID": 1,
				"memoryUsedPc": 62.51,
				"cpuUsedPc": 0.13,
				"id": 1634,
				"rss": 1282364
			},
			{
				"vss": 2552316,
				"name": "mysqld",
				"tgid": 4516,
				"vmlimit": "unlimited",
				"parentID": 1,
				"memoryUsedPc": 62.51,
				"cpuUsedPc": 0.13,
				"id": 1635,
				"rss": 1282364
			},
			{
				"vss": 2552316,
				"name": "mysqld",
				"tgid": 4516,
				"vmlimit": "unlimited",
				"parentID": 1,
				"memoryUsedPc": 62.51,
				"cpuUsedPc": 0.13,
				"id": 1636,
				"rss": 1282364
			},
			{
				"vss": 2552316,
				"name": "mysqld",
				"tgid": 4516,
				"vmlimit": "unlimited",
				"parentID": 1,
				"memoryUsedPc": 62.51,
				"cpuUsedPc": 0.13,
				"id": 1637,
				"rss": 1282364
			},
			{
				"vss": 2552316,
				"name": "mysqld",
				"tgid": 4516,
				"vmlimit": "unlimited",
				"parentID": 1,
				"memoryUsedPc": 62.51,
				"cpuUsedPc": 0.12,
				"id": 1638,
				"rss": 1282364
			},
			{
				"vss": 2552316,
				"name": "mysqld",
				"tgid": 4516,
				"vmlimit": "unlimited",
				"parentID": 1,
				"memoryUsedPc": 62.51,
				"cpuUsedPc": 0.12,
				"id": 1639,
				"rss": 1282364
			},
			{
				"vss": 2552316,
				"name": "mysqld",
				"tgid": 4516,
				"vmlimit": "unlimited",
				"parentID": 1,
				"memoryUsedPc": 62.51,
				"cpuUsedPc": 0.15,
				"id": 1640,
				"rss": 1282364
			},
			{
				"vss": 2552316,
				"name": "mysqld",
				"tgid": 4516,
				"vmlimit": "unlimited",
				"parentID": 1,
				"memoryUsedPc": 62.51,
				"cpuUsedPc": 0.13,
				"id": 1641,
				"rss": 1282364
			},
			{
				"vss": 2552316,
				"name": "mysqld",
				"tgid": 4516,
				"vmlimit": "unlimited",
				"parentID": 1,
				"memoryUsedPc": 62.51,
				"cpuUsedPc": 0.13,
				"id": 1642,
				"rss": 1282364
			},
			{
				"vss": 2552316,
				"name": "mysqld",
				"tgid": 4516,
				"vmlimit": "unlimited",
				"parentID": 1,
				"memoryUsedPc": 62.51,
				"cpuUsedPc": 0.13,
				"id": 1643,
				"rss": 1282364
			},
			{
				"vss": 2552316,
				"name": "mysqld",
				"tgid": 4516,
				"vmlimit": "unlimited",
				"parentID": 1,
				"memoryUsedPc": 62.51,
				"cpuUsedPc": 0.00,
				"id": 4516,
				"rss": 1282364
			},
			{
				"vss": 2552316,
				"name": "mysqld",
				"tgid": 4516,
				"vmlimit": "unlimited",
				"parentID": 1,
				"memoryUsedPc": 62.51,
				"cpuUsedPc": 0.33,
				"id": 4551,
				"rss": 1282364
			},
			{
				"vss": 723152,
				"name": "OS processes",
				"tgid": 0,
				"vmlimit": "",
				"parentID": 0,
				"memoryUsedPc": 0.53,
				"cpuUsedPc": 0.02,
				"id": 0,
				"rss": 10968
			},
			{
				"vss": 1307432,
				"name": "RDS processes",
				"tgid": 0,
				"vmlimit": "",
				"parentID": 0,
				"memoryUsedPc": 10.58,
				"cpuUsedPc": 0.14,
				"id": 0,
				"rss": 217304
			}
		]
	}`))

	dataMySQL57Expected = strings.TrimSpace(`
# HELP node_cpu_average The percentage of CPU utilization. Units: Percent
# TYPE node_cpu_average gauge
node_cpu_average{cpu="All",instance="rds-mysql57",mode="guest",region="us-east-1"} 0
node_cpu_average{cpu="All",instance="rds-mysql57",mode="idle",region="us-east-1"} 0
node_cpu_average{cpu="All",instance="rds-mysql57",mode="irq",region="us-east-1"} 1.16
node_cpu_average{cpu="All",instance="rds-mysql57",mode="nice",region="us-east-1"} 10.5
node_cpu_average{cpu="All",instance="rds-mysql57",mode="steal",region="us-east-1"} 1.01
node_cpu_average{cpu="All",instance="rds-mysql57",mode="system",region="us-east-1"} 4.71
node_cpu_average{cpu="All",instance="rds-mysql57",mode="total",region="us-east-1"} 99.99
node_cpu_average{cpu="All",instance="rds-mysql57",mode="user",region="us-east-1"} 1.18
node_cpu_average{cpu="All",instance="rds-mysql57",mode="wait",region="us-east-1"} 81.43
# HELP rdsosmetrics_cpuUtilization_guest The percentage of CPU in use by guest programs.
# TYPE rdsosmetrics_cpuUtilization_guest gauge
rdsosmetrics_cpuUtilization_guest{instance="rds-mysql57",region="us-east-1"} 0
# HELP rdsosmetrics_cpuUtilization_idle The percentage of CPU that is idle.
# TYPE rdsosmetrics_cpuUtilization_idle gauge
rdsosmetrics_cpuUtilization_idle{instance="rds-mysql57",region="us-east-1"} 0
# HELP rdsosmetrics_cpuUtilization_irq The percentage of CPU in use by software interrupts.
# TYPE rdsosmetrics_cpuUtilization_irq gauge
rdsosmetrics_cpuUtilization_irq{instance="rds-mysql57",region="us-east-1"} 1.16
# HELP rdsosmetrics_cpuUtilization_nice The percentage of CPU in use by programs running at lowest priority.
# TYPE rdsosmetrics_cpuUtilization_nice gauge
rdsosmetrics_cpuUtilization_nice{instance="rds-mysql57",region="us-east-1"} 10.5
# HELP rdsosmetrics_cpuUtilization_steal The percentage of CPU in use by other virtual machines.
# TYPE rdsosmetrics_cpuUtilization_steal gauge
rdsosmetrics_cpuUtilization_steal{instance="rds-mysql57",region="us-east-1"} 1.01
# HELP rdsosmetrics_cpuUtilization_system The percentage of CPU in use by the kernel.
# TYPE rdsosmetrics_cpuUtilization_system gauge
rdsosmetrics_cpuUtilization_system{instance="rds-mysql57",region="us-east-1"} 4.71
# HELP rdsosmetrics_cpuUtilization_total The total percentage of the CPU in use. This value includes the nice value.
# TYPE rdsosmetrics_cpuUtilization_total gauge
rdsosmetrics_cpuUtilization_total{instance="rds-mysql57",region="us-east-1"} 99.99
# HELP rdsosmetrics_cpuUtilization_user The percentage of CPU in use by user programs.
# TYPE rdsosmetrics_cpuUtilization_user gauge
rdsosmetrics_cpuUtilization_user{instance="rds-mysql57",region="us-east-1"} 1.18
# HELP rdsosmetrics_cpuUtilization_wait The percentage of CPU unused while waiting for I/O access.
# TYPE rdsosmetrics_cpuUtilization_wait gauge
rdsosmetrics_cpuUtilization_wait{instance="rds-mysql57",region="us-east-1"} 81.43
# HELP rdsosmetrics_diskIO_avgQueueLen The number of requests waiting in the I/O device's queue.
# TYPE rdsosmetrics_diskIO_avgQueueLen gauge
rdsosmetrics_diskIO_avgQueueLen{device="rdsdev",instance="rds-mysql57",region="us-east-1"} 2783.31
# HELP rdsosmetrics_diskIO_avgReqSz The average request size, in kilobytes.
# TYPE rdsosmetrics_diskIO_avgReqSz gauge
rdsosmetrics_diskIO_avgReqSz{device="rdsdev",instance="rds-mysql57",region="us-east-1"} 18.97
# HELP rdsosmetrics_diskIO_await The number of milliseconds required to respond to requests, including queue time and service time.
# TYPE rdsosmetrics_diskIO_await gauge
rdsosmetrics_diskIO_await{device="rdsdev",instance="rds-mysql57",region="us-east-1"} 74.67
# HELP rdsosmetrics_diskIO_readIOsPS The number of read operations per second.
# TYPE rdsosmetrics_diskIO_readIOsPS gauge
rdsosmetrics_diskIO_readIOsPS{device="rdsdev",instance="rds-mysql57",region="us-east-1"} 354.18
# HELP rdsosmetrics_diskIO_readKb The total number of kilobytes read.
# TYPE rdsosmetrics_diskIO_readKb gauge
rdsosmetrics_diskIO_readKb{device="rdsdev",instance="rds-mysql57",region="us-east-1"} 339896
# HELP rdsosmetrics_diskIO_readKbPS The number of kilobytes read per second.
# TYPE rdsosmetrics_diskIO_readKbPS gauge
rdsosmetrics_diskIO_readKbPS{device="rdsdev",instance="rds-mysql57",region="us-east-1"} 5664.93
# HELP rdsosmetrics_diskIO_rrqmPS The number of merged read requests queued per second.
# TYPE rdsosmetrics_diskIO_rrqmPS gauge
rdsosmetrics_diskIO_rrqmPS{device="rdsdev",instance="rds-mysql57",region="us-east-1"} 0
# HELP rdsosmetrics_diskIO_tps The number of I/O transactions per second.
# TYPE rdsosmetrics_diskIO_tps gauge
rdsosmetrics_diskIO_tps{device="rdsdev",instance="rds-mysql57",region="us-east-1"} 623.33
# HELP rdsosmetrics_diskIO_util The percentage of CPU time during which requests were issued.
# TYPE rdsosmetrics_diskIO_util gauge
rdsosmetrics_diskIO_util{device="rdsdev",instance="rds-mysql57",region="us-east-1"} 99.94
# HELP rdsosmetrics_diskIO_writeIOsPS The number of write operations per second.
# TYPE rdsosmetrics_diskIO_writeIOsPS gauge
rdsosmetrics_diskIO_writeIOsPS{device="rdsdev",instance="rds-mysql57",region="us-east-1"} 269.15
# HELP rdsosmetrics_diskIO_writeKb The total number of kilobytes written.
# TYPE rdsosmetrics_diskIO_writeKb gauge
rdsosmetrics_diskIO_writeKb{device="rdsdev",instance="rds-mysql57",region="us-east-1"} 369436
# HELP rdsosmetrics_diskIO_writeKbPS The number of kilobytes written per second.
# TYPE rdsosmetrics_diskIO_writeKbPS gauge
rdsosmetrics_diskIO_writeKbPS{device="rdsdev",instance="rds-mysql57",region="us-east-1"} 6157.27
# HELP rdsosmetrics_diskIO_wrqmPS The number of merged write requests queued per second.
# TYPE rdsosmetrics_diskIO_wrqmPS gauge
rdsosmetrics_diskIO_wrqmPS{device="rdsdev",instance="rds-mysql57",region="us-east-1"} 0
# HELP rdsosmetrics_fileSys_maxFiles The maximum number of files that can be created for the file system.
# TYPE rdsosmetrics_fileSys_maxFiles gauge
rdsosmetrics_fileSys_maxFiles{instance="rds-mysql57",mount_point="/rdsdbdata",name="rdsfilesys",region="us-east-1"} 3.93216e+06
# HELP rdsosmetrics_fileSys_total The total number of disk space available for the file system, in kilobytes.
# TYPE rdsosmetrics_fileSys_total gauge
rdsosmetrics_fileSys_total{instance="rds-mysql57",mount_point="/rdsdbdata",name="rdsfilesys",region="us-east-1"} 6.1774768e+07
# HELP rdsosmetrics_fileSys_used The amount of disk space used by files in the file system, in kilobytes.
# TYPE rdsosmetrics_fileSys_used gauge
rdsosmetrics_fileSys_used{instance="rds-mysql57",mount_point="/rdsdbdata",name="rdsfilesys",region="us-east-1"} 9.659328e+06
# HELP rdsosmetrics_fileSys_usedFilePercent The percentage of available files in use.
# TYPE rdsosmetrics_fileSys_usedFilePercent gauge
rdsosmetrics_fileSys_usedFilePercent{instance="rds-mysql57",mount_point="/rdsdbdata",name="rdsfilesys",region="us-east-1"} 0.02
# HELP rdsosmetrics_fileSys_usedFiles The number of files in the file system.
# TYPE rdsosmetrics_fileSys_usedFiles gauge
rdsosmetrics_fileSys_usedFiles{instance="rds-mysql57",mount_point="/rdsdbdata",name="rdsfilesys",region="us-east-1"} 731
# HELP rdsosmetrics_fileSys_usedPercent The percentage of the file-system disk space in use.
# TYPE rdsosmetrics_fileSys_usedPercent gauge
rdsosmetrics_fileSys_usedPercent{instance="rds-mysql57",mount_point="/rdsdbdata",name="rdsfilesys",region="us-east-1"} 15.64
# HELP rdsosmetrics_loadAverageMinute_fifteen The number of processes requesting CPU time over the last 15 minutes.
# TYPE rdsosmetrics_loadAverageMinute_fifteen gauge
rdsosmetrics_loadAverageMinute_fifteen{instance="rds-mysql57",region="us-east-1"} 7.64
# HELP rdsosmetrics_loadAverageMinute_five The number of processes requesting CPU time over the last 5 minutes.
# TYPE rdsosmetrics_loadAverageMinute_five gauge
rdsosmetrics_loadAverageMinute_five{instance="rds-mysql57",region="us-east-1"} 14.08
# HELP rdsosmetrics_loadAverageMinute_one The number of processes requesting CPU time over the last minute.0
# TYPE rdsosmetrics_loadAverageMinute_one gauge
rdsosmetrics_loadAverageMinute_one{instance="rds-mysql57",region="us-east-1"} 32.83
# HELP rdsosmetrics_memory_active The amount of assigned memory, in kilobytes.
# TYPE rdsosmetrics_memory_active gauge
rdsosmetrics_memory_active{instance="rds-mysql57",region="us-east-1"} 1.299668e+06
# HELP rdsosmetrics_memory_buffers The amount of memory used for buffering I/O requests prior to writing to the storage device, in kilobytes.
# TYPE rdsosmetrics_memory_buffers gauge
rdsosmetrics_memory_buffers{instance="rds-mysql57",region="us-east-1"} 149412
# HELP rdsosmetrics_memory_cached The amount of memory used for caching file system–based I/O.
# TYPE rdsosmetrics_memory_cached gauge
rdsosmetrics_memory_cached{instance="rds-mysql57",region="us-east-1"} 162452
# HELP rdsosmetrics_memory_dirty The amount of memory pages in RAM that have been modified but not written to their related data block in storage, in kilobytes.
# TYPE rdsosmetrics_memory_dirty gauge
rdsosmetrics_memory_dirty{instance="rds-mysql57",region="us-east-1"} 360
# HELP rdsosmetrics_memory_free The amount of unassigned memory, in kilobytes.
# TYPE rdsosmetrics_memory_free gauge
rdsosmetrics_memory_free{instance="rds-mysql57",region="us-east-1"} 127048
# HELP rdsosmetrics_memory_hugePagesFree The number of free huge pages. Huge pages are a feature of the Linux kernel.
# TYPE rdsosmetrics_memory_hugePagesFree gauge
rdsosmetrics_memory_hugePagesFree{instance="rds-mysql57",region="us-east-1"} 0
# HELP rdsosmetrics_memory_hugePagesRsvd The number of committed huge pages.
# TYPE rdsosmetrics_memory_hugePagesRsvd gauge
rdsosmetrics_memory_hugePagesRsvd{instance="rds-mysql57",region="us-east-1"} 0
# HELP rdsosmetrics_memory_hugePagesSize The size for each huge pages unit, in kilobytes.
# TYPE rdsosmetrics_memory_hugePagesSize gauge
rdsosmetrics_memory_hugePagesSize{instance="rds-mysql57",region="us-east-1"} 2048
# HELP rdsosmetrics_memory_hugePagesSurp The number of available surplus huge pages over the total.
# TYPE rdsosmetrics_memory_hugePagesSurp gauge
rdsosmetrics_memory_hugePagesSurp{instance="rds-mysql57",region="us-east-1"} 0
# HELP rdsosmetrics_memory_hugePagesTotal The total number of huge pages for the system.
# TYPE rdsosmetrics_memory_hugePagesTotal gauge
rdsosmetrics_memory_hugePagesTotal{instance="rds-mysql57",region="us-east-1"} 0
# HELP rdsosmetrics_memory_inactive The amount of least-frequently used memory pages, in kilobytes.
# TYPE rdsosmetrics_memory_inactive gauge
rdsosmetrics_memory_inactive{instance="rds-mysql57",region="us-east-1"} 511232
# HELP rdsosmetrics_memory_mapped The total amount of file-system contents that is memory mapped inside a process address space, in kilobytes.
# TYPE rdsosmetrics_memory_mapped gauge
rdsosmetrics_memory_mapped{instance="rds-mysql57",region="us-east-1"} 14184
# HELP rdsosmetrics_memory_pageTables The amount of memory used by page tables, in kilobytes.
# TYPE rdsosmetrics_memory_pageTables gauge
rdsosmetrics_memory_pageTables{instance="rds-mysql57",region="us-east-1"} 7916
# HELP rdsosmetrics_memory_slab The amount of reusable kernel data structures, in kilobytes.
# TYPE rdsosmetrics_memory_slab gauge
rdsosmetrics_memory_slab{instance="rds-mysql57",region="us-east-1"} 45440
# HELP rdsosmetrics_memory_total The total amount of memory, in kilobytes.
# TYPE rdsosmetrics_memory_total gauge
rdsosmetrics_memory_total{instance="rds-mysql57",region="us-east-1"} 2.05152e+06
# HELP rdsosmetrics_memory_writeback The amount of dirty pages in RAM that are still being written to the backing storage, in kilobytes.
# TYPE rdsosmetrics_memory_writeback gauge
rdsosmetrics_memory_writeback{instance="rds-mysql57",region="us-east-1"} 0
# HELP rdsosmetrics_network_rx The number of bytes received per second.
# TYPE rdsosmetrics_network_rx gauge
rdsosmetrics_network_rx{instance="rds-mysql57",interface="eth0",region="us-east-1"} 93611.95
# HELP rdsosmetrics_network_tx The number of bytes uploaded per second.
# TYPE rdsosmetrics_network_tx gauge
rdsosmetrics_network_tx{instance="rds-mysql57",interface="eth0",region="us-east-1"} 801373.55
# HELP rdsosmetrics_processList_cpuUsedPc The percentage of CPU used by the process.
# TYPE rdsosmetrics_processList_cpuUsedPc gauge
rdsosmetrics_processList_cpuUsedPc{id="0",instance="rds-mysql57",name="OS processes",parentID="0",region="us-east-1",tgid="0"} 0.02
rdsosmetrics_processList_cpuUsedPc{id="0",instance="rds-mysql57",name="RDS processes",parentID="0",region="us-east-1",tgid="0"} 0.14
rdsosmetrics_processList_cpuUsedPc{id="1594",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="4516"} 0.13
rdsosmetrics_processList_cpuUsedPc{id="1595",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="4516"} 0.12
rdsosmetrics_processList_cpuUsedPc{id="1596",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="4516"} 0.12
rdsosmetrics_processList_cpuUsedPc{id="1597",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="4516"} 0.13
rdsosmetrics_processList_cpuUsedPc{id="1599",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="4516"} 0.13
rdsosmetrics_processList_cpuUsedPc{id="1600",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="4516"} 0.12
rdsosmetrics_processList_cpuUsedPc{id="1601",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="4516"} 0.13
rdsosmetrics_processList_cpuUsedPc{id="1602",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="4516"} 0.13
rdsosmetrics_processList_cpuUsedPc{id="1603",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="4516"} 0.1
rdsosmetrics_processList_cpuUsedPc{id="1604",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="4516"} 0.12
rdsosmetrics_processList_cpuUsedPc{id="1605",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="4516"} 0.13
rdsosmetrics_processList_cpuUsedPc{id="1606",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="4516"} 0.13
rdsosmetrics_processList_cpuUsedPc{id="1607",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="4516"} 0.13
rdsosmetrics_processList_cpuUsedPc{id="1608",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="4516"} 0.13
rdsosmetrics_processList_cpuUsedPc{id="1609",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="4516"} 0.13
rdsosmetrics_processList_cpuUsedPc{id="1610",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="4516"} 0.12
rdsosmetrics_processList_cpuUsedPc{id="1611",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="4516"} 0.13
rdsosmetrics_processList_cpuUsedPc{id="1612",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="4516"} 0.13
rdsosmetrics_processList_cpuUsedPc{id="1613",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="4516"} 0.12
rdsosmetrics_processList_cpuUsedPc{id="1614",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="4516"} 0.13
rdsosmetrics_processList_cpuUsedPc{id="1615",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="4516"} 0.12
rdsosmetrics_processList_cpuUsedPc{id="1616",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="4516"} 0.13
rdsosmetrics_processList_cpuUsedPc{id="1617",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="4516"} 0.13
rdsosmetrics_processList_cpuUsedPc{id="1618",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="4516"} 0.13
rdsosmetrics_processList_cpuUsedPc{id="1619",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="4516"} 0.12
rdsosmetrics_processList_cpuUsedPc{id="1620",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="4516"} 0.13
rdsosmetrics_processList_cpuUsedPc{id="1621",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="4516"} 0.13
rdsosmetrics_processList_cpuUsedPc{id="1622",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="4516"} 0.12
rdsosmetrics_processList_cpuUsedPc{id="1623",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="4516"} 0.12
rdsosmetrics_processList_cpuUsedPc{id="1624",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="4516"} 0.12
rdsosmetrics_processList_cpuUsedPc{id="1625",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="4516"} 0.13
rdsosmetrics_processList_cpuUsedPc{id="1626",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="4516"} 0.15
rdsosmetrics_processList_cpuUsedPc{id="1627",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="4516"} 0.1
rdsosmetrics_processList_cpuUsedPc{id="1628",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="4516"} 0.13
rdsosmetrics_processList_cpuUsedPc{id="1629",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="4516"} 0.12
rdsosmetrics_processList_cpuUsedPc{id="1630",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="4516"} 0.12
rdsosmetrics_processList_cpuUsedPc{id="1631",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="4516"} 0.13
rdsosmetrics_processList_cpuUsedPc{id="1632",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="4516"} 0.13
rdsosmetrics_processList_cpuUsedPc{id="1633",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="4516"} 0.13
rdsosmetrics_processList_cpuUsedPc{id="1634",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="4516"} 0.13
rdsosmetrics_processList_cpuUsedPc{id="1635",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="4516"} 0.13
rdsosmetrics_processList_cpuUsedPc{id="1636",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="4516"} 0.13
rdsosmetrics_processList_cpuUsedPc{id="1637",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="4516"} 0.13
rdsosmetrics_processList_cpuUsedPc{id="1638",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="4516"} 0.12
rdsosmetrics_processList_cpuUsedPc{id="1639",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="4516"} 0.12
rdsosmetrics_processList_cpuUsedPc{id="1640",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="4516"} 0.15
rdsosmetrics_processList_cpuUsedPc{id="1641",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="4516"} 0.13
rdsosmetrics_processList_cpuUsedPc{id="1642",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="4516"} 0.13
rdsosmetrics_processList_cpuUsedPc{id="1643",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="4516"} 0.13
rdsosmetrics_processList_cpuUsedPc{id="4516",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="4516"} 0
rdsosmetrics_processList_cpuUsedPc{id="4551",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="4516"} 0.33
# HELP rdsosmetrics_processList_memoryUsedPc The amount of memory used by the process, in kilobytes.
# TYPE rdsosmetrics_processList_memoryUsedPc gauge
rdsosmetrics_processList_memoryUsedPc{id="0",instance="rds-mysql57",name="OS processes",parentID="0",region="us-east-1",tgid="0"} 0.53
rdsosmetrics_processList_memoryUsedPc{id="0",instance="rds-mysql57",name="RDS processes",parentID="0",region="us-east-1",tgid="0"} 10.58
rdsosmetrics_processList_memoryUsedPc{id="1594",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="4516"} 62.51
rdsosmetrics_processList_memoryUsedPc{id="1595",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="4516"} 62.51
rdsosmetrics_processList_memoryUsedPc{id="1596",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="4516"} 62.51
rdsosmetrics_processList_memoryUsedPc{id="1597",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="4516"} 62.51
rdsosmetrics_processList_memoryUsedPc{id="1599",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="4516"} 62.51
rdsosmetrics_processList_memoryUsedPc{id="1600",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="4516"} 62.51
rdsosmetrics_processList_memoryUsedPc{id="1601",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="4516"} 62.51
rdsosmetrics_processList_memoryUsedPc{id="1602",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="4516"} 62.51
rdsosmetrics_processList_memoryUsedPc{id="1603",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="4516"} 62.51
rdsosmetrics_processList_memoryUsedPc{id="1604",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="4516"} 62.51
rdsosmetrics_processList_memoryUsedPc{id="1605",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="4516"} 62.51
rdsosmetrics_processList_memoryUsedPc{id="1606",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="4516"} 62.51
rdsosmetrics_processList_memoryUsedPc{id="1607",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="4516"} 62.51
rdsosmetrics_processList_memoryUsedPc{id="1608",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="4516"} 62.51
rdsosmetrics_processList_memoryUsedPc{id="1609",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="4516"} 62.51
rdsosmetrics_processList_memoryUsedPc{id="1610",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="4516"} 62.51
rdsosmetrics_processList_memoryUsedPc{id="1611",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="4516"} 62.51
rdsosmetrics_processList_memoryUsedPc{id="1612",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="4516"} 62.51
rdsosmetrics_processList_memoryUsedPc{id="1613",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="4516"} 62.51
rdsosmetrics_processList_memoryUsedPc{id="1614",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="4516"} 62.51
rdsosmetrics_processList_memoryUsedPc{id="1615",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="4516"} 62.51
rdsosmetrics_processList_memoryUsedPc{id="1616",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="4516"} 62.51
rdsosmetrics_processList_memoryUsedPc{id="1617",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="4516"} 62.51
rdsosmetrics_processList_memoryUsedPc{id="1618",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="4516"} 62.51
rdsosmetrics_processList_memoryUsedPc{id="1619",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="4516"} 62.51
rdsosmetrics_processList_memoryUsedPc{id="1620",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="4516"} 62.51
rdsosmetrics_processList_memoryUsedPc{id="1621",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="4516"} 62.51
rdsosmetrics_processList_memoryUsedPc{id="1622",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="4516"} 62.51
rdsosmetrics_processList_memoryUsedPc{id="1623",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="4516"} 62.51
rdsosmetrics_processList_memoryUsedPc{id="1624",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="4516"} 62.51
rdsosmetrics_processList_memoryUsedPc{id="1625",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="4516"} 62.51
rdsosmetrics_processList_memoryUsedPc{id="1626",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="4516"} 62.51
rdsosmetrics_processList_memoryUsedPc{id="1627",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="4516"} 62.51
rdsosmetrics_processList_memoryUsedPc{id="1628",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="4516"} 62.51
rdsosmetrics_processList_memoryUsedPc{id="1629",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="4516"} 62.51
rdsosmetrics_processList_memoryUsedPc{id="1630",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="4516"} 62.51
rdsosmetrics_processList_memoryUsedPc{id="1631",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="4516"} 62.51
rdsosmetrics_processList_memoryUsedPc{id="1632",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="4516"} 62.51
rdsosmetrics_processList_memoryUsedPc{id="1633",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="4516"} 62.51
rdsosmetrics_processList_memoryUsedPc{id="1634",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="4516"} 62.51
rdsosmetrics_processList_memoryUsedPc{id="1635",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="4516"} 62.51
rdsosmetrics_processList_memoryUsedPc{id="1636",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="4516"} 62.51
rdsosmetrics_processList_memoryUsedPc{id="1637",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="4516"} 62.51
rdsosmetrics_processList_memoryUsedPc{id="1638",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="4516"} 62.51
rdsosmetrics_processList_memoryUsedPc{id="1639",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="4516"} 62.51
rdsosmetrics_processList_memoryUsedPc{id="1640",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="4516"} 62.51
rdsosmetrics_processList_memoryUsedPc{id="1641",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="4516"} 62.51
rdsosmetrics_processList_memoryUsedPc{id="1642",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="4516"} 62.51
rdsosmetrics_processList_memoryUsedPc{id="1643",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="4516"} 62.51
rdsosmetrics_processList_memoryUsedPc{id="4516",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="4516"} 62.51
rdsosmetrics_processList_memoryUsedPc{id="4551",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="4516"} 62.51
# HELP rdsosmetrics_processList_rss The amount of RAM allocated to the process, in kilobytes.
# TYPE rdsosmetrics_processList_rss gauge
rdsosmetrics_processList_rss{id="0",instance="rds-mysql57",name="OS processes",parentID="0",region="us-east-1",tgid="0"} 10968
rdsosmetrics_processList_rss{id="0",instance="rds-mysql57",name="RDS processes",parentID="0",region="us-east-1",tgid="0"} 217304
rdsosmetrics_processList_rss{id="1594",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="4516"} 1.282364e+06
rdsosmetrics_processList_rss{id="1595",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="4516"} 1.282364e+06
rdsosmetrics_processList_rss{id="1596",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="4516"} 1.282364e+06
rdsosmetrics_processList_rss{id="1597",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="4516"} 1.282364e+06
rdsosmetrics_processList_rss{id="1599",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="4516"} 1.282364e+06
rdsosmetrics_processList_rss{id="1600",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="4516"} 1.282364e+06
rdsosmetrics_processList_rss{id="1601",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="4516"} 1.282364e+06
rdsosmetrics_processList_rss{id="1602",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="4516"} 1.282364e+06
rdsosmetrics_processList_rss{id="1603",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="4516"} 1.282364e+06
rdsosmetrics_processList_rss{id="1604",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="4516"} 1.282364e+06
rdsosmetrics_processList_rss{id="1605",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="4516"} 1.282364e+06
rdsosmetrics_processList_rss{id="1606",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="4516"} 1.282364e+06
rdsosmetrics_processList_rss{id="1607",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="4516"} 1.282364e+06
rdsosmetrics_processList_rss{id="1608",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="4516"} 1.282364e+06
rdsosmetrics_processList_rss{id="1609",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="4516"} 1.282364e+06
rdsosmetrics_processList_rss{id="1610",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="4516"} 1.282364e+06
rdsosmetrics_processList_rss{id="1611",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="4516"} 1.282364e+06
rdsosmetrics_processList_rss{id="1612",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="4516"} 1.282364e+06
rdsosmetrics_processList_rss{id="1613",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="4516"} 1.282364e+06
rdsosmetrics_processList_rss{id="1614",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="4516"} 1.282364e+06
rdsosmetrics_processList_rss{id="1615",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="4516"} 1.282364e+06
rdsosmetrics_processList_rss{id="1616",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="4516"} 1.282364e+06
rdsosmetrics_processList_rss{id="1617",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="4516"} 1.282364e+06
rdsosmetrics_processList_rss{id="1618",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="4516"} 1.282364e+06
rdsosmetrics_processList_rss{id="1619",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="4516"} 1.282364e+06
rdsosmetrics_processList_rss{id="1620",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="4516"} 1.282364e+06
rdsosmetrics_processList_rss{id="1621",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="4516"} 1.282364e+06
rdsosmetrics_processList_rss{id="1622",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="4516"} 1.282364e+06
rdsosmetrics_processList_rss{id="1623",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="4516"} 1.282364e+06
rdsosmetrics_processList_rss{id="1624",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="4516"} 1.282364e+06
rdsosmetrics_processList_rss{id="1625",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="4516"} 1.282364e+06
rdsosmetrics_processList_rss{id="1626",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="4516"} 1.282364e+06
rdsosmetrics_processList_rss{id="1627",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="4516"} 1.282364e+06
rdsosmetrics_processList_rss{id="1628",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="4516"} 1.282364e+06
rdsosmetrics_processList_rss{id="1629",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="4516"} 1.282364e+06
rdsosmetrics_processList_rss{id="1630",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="4516"} 1.282364e+06
rdsosmetrics_processList_rss{id="1631",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="4516"} 1.282364e+06
rdsosmetrics_processList_rss{id="1632",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="4516"} 1.282364e+06
rdsosmetrics_processList_rss{id="1633",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="4516"} 1.282364e+06
rdsosmetrics_processList_rss{id="1634",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="4516"} 1.282364e+06
rdsosmetrics_processList_rss{id="1635",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="4516"} 1.282364e+06
rdsosmetrics_processList_rss{id="1636",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="4516"} 1.282364e+06
rdsosmetrics_processList_rss{id="1637",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="4516"} 1.282364e+06
rdsosmetrics_processList_rss{id="1638",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="4516"} 1.282364e+06
rdsosmetrics_processList_rss{id="1639",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="4516"} 1.282364e+06
rdsosmetrics_processList_rss{id="1640",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="4516"} 1.282364e+06
rdsosmetrics_processList_rss{id="1641",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="4516"} 1.282364e+06
rdsosmetrics_processList_rss{id="1642",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="4516"} 1.282364e+06
rdsosmetrics_processList_rss{id="1643",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="4516"} 1.282364e+06
rdsosmetrics_processList_rss{id="4516",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="4516"} 1.282364e+06
rdsosmetrics_processList_rss{id="4551",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="4516"} 1.282364e+06
# HELP rdsosmetrics_processList_vss The amount of virtual memory allocated to the process, in kilobytes.
# TYPE rdsosmetrics_processList_vss gauge
rdsosmetrics_processList_vss{id="0",instance="rds-mysql57",name="OS processes",parentID="0",region="us-east-1",tgid="0"} 723152
rdsosmetrics_processList_vss{id="0",instance="rds-mysql57",name="RDS processes",parentID="0",region="us-east-1",tgid="0"} 1.307432e+06
rdsosmetrics_processList_vss{id="1594",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="4516"} 2.552316e+06
rdsosmetrics_processList_vss{id="1595",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="4516"} 2.552316e+06
rdsosmetrics_processList_vss{id="1596",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="4516"} 2.552316e+06
rdsosmetrics_processList_vss{id="1597",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="4516"} 2.552316e+06
rdsosmetrics_processList_vss{id="1599",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="4516"} 2.552316e+06
rdsosmetrics_processList_vss{id="1600",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="4516"} 2.552316e+06
rdsosmetrics_processList_vss{id="1601",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="4516"} 2.552316e+06
rdsosmetrics_processList_vss{id="1602",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="4516"} 2.552316e+06
rdsosmetrics_processList_vss{id="1603",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="4516"} 2.552316e+06
rdsosmetrics_processList_vss{id="1604",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="4516"} 2.552316e+06
rdsosmetrics_processList_vss{id="1605",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="4516"} 2.552316e+06
rdsosmetrics_processList_vss{id="1606",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="4516"} 2.552316e+06
rdsosmetrics_processList_vss{id="1607",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="4516"} 2.552316e+06
rdsosmetrics_processList_vss{id="1608",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="4516"} 2.552316e+06
rdsosmetrics_processList_vss{id="1609",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="4516"} 2.552316e+06
rdsosmetrics_processList_vss{id="1610",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="4516"} 2.552316e+06
rdsosmetrics_processList_vss{id="1611",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="4516"} 2.552316e+06
rdsosmetrics_processList_vss{id="1612",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="4516"} 2.552316e+06
rdsosmetrics_processList_vss{id="1613",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="4516"} 2.552316e+06
rdsosmetrics_processList_vss{id="1614",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="4516"} 2.552316e+06
rdsosmetrics_processList_vss{id="1615",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="4516"} 2.552316e+06
rdsosmetrics_processList_vss{id="1616",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="4516"} 2.552316e+06
rdsosmetrics_processList_vss{id="1617",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="4516"} 2.552316e+06
rdsosmetrics_processList_vss{id="1618",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="4516"} 2.552316e+06
rdsosmetrics_processList_vss{id="1619",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="4516"} 2.552316e+06
rdsosmetrics_processList_vss{id="1620",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="4516"} 2.552316e+06
rdsosmetrics_processList_vss{id="1621",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="4516"} 2.552316e+06
rdsosmetrics_processList_vss{id="1622",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="4516"} 2.552316e+06
rdsosmetrics_processList_vss{id="1623",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="4516"} 2.552316e+06
rdsosmetrics_processList_vss{id="1624",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="4516"} 2.552316e+06
rdsosmetrics_processList_vss{id="1625",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="4516"} 2.552316e+06
rdsosmetrics_processList_vss{id="1626",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="4516"} 2.552316e+06
rdsosmetrics_processList_vss{id="1627",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="4516"} 2.552316e+06
rdsosmetrics_processList_vss{id="1628",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="4516"} 2.552316e+06
rdsosmetrics_processList_vss{id="1629",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="4516"} 2.552316e+06
rdsosmetrics_processList_vss{id="1630",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="4516"} 2.552316e+06
rdsosmetrics_processList_vss{id="1631",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="4516"} 2.552316e+06
rdsosmetrics_processList_vss{id="1632",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="4516"} 2.552316e+06
rdsosmetrics_processList_vss{id="1633",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="4516"} 2.552316e+06
rdsosmetrics_processList_vss{id="1634",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="4516"} 2.552316e+06
rdsosmetrics_processList_vss{id="1635",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="4516"} 2.552316e+06
rdsosmetrics_processList_vss{id="1636",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="4516"} 2.552316e+06
rdsosmetrics_processList_vss{id="1637",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="4516"} 2.552316e+06
rdsosmetrics_processList_vss{id="1638",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="4516"} 2.552316e+06
rdsosmetrics_processList_vss{id="1639",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="4516"} 2.552316e+06
rdsosmetrics_processList_vss{id="1640",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="4516"} 2.552316e+06
rdsosmetrics_processList_vss{id="1641",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="4516"} 2.552316e+06
rdsosmetrics_processList_vss{id="1642",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="4516"} 2.552316e+06
rdsosmetrics_processList_vss{id="1643",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="4516"} 2.552316e+06
rdsosmetrics_processList_vss{id="4516",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="4516"} 2.552316e+06
rdsosmetrics_processList_vss{id="4551",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="4516"} 2.552316e+06
# HELP rdsosmetrics_swap_cached The amount of swap memory, in kilobytes, used as cache memory.
# TYPE rdsosmetrics_swap_cached gauge
rdsosmetrics_swap_cached{instance="rds-mysql57",region="us-east-1"} 908
# HELP rdsosmetrics_swap_free The total amount of swap memory free, in kilobytes.
# TYPE rdsosmetrics_swap_free gauge
rdsosmetrics_swap_free{instance="rds-mysql57",region="us-east-1"} 3.803212e+06
# HELP rdsosmetrics_swap_total The total amount of swap memory available, in kilobytes.
# TYPE rdsosmetrics_swap_total gauge
rdsosmetrics_swap_total{instance="rds-mysql57",region="us-east-1"} 4.095996e+06
# HELP rdsosmetrics_tasks_blocked The number of tasks that are blocked.
# TYPE rdsosmetrics_tasks_blocked gauge
rdsosmetrics_tasks_blocked{instance="rds-mysql57",region="us-east-1"} 54
# HELP rdsosmetrics_tasks_running The number of tasks that are running.
# TYPE rdsosmetrics_tasks_running gauge
rdsosmetrics_tasks_running{instance="rds-mysql57",region="us-east-1"} 1
# HELP rdsosmetrics_tasks_sleeping The number of tasks that are sleeping.
# TYPE rdsosmetrics_tasks_sleeping gauge
rdsosmetrics_tasks_sleeping{instance="rds-mysql57",region="us-east-1"} 258
# HELP rdsosmetrics_tasks_stopped The number of tasks that are stopped.
# TYPE rdsosmetrics_tasks_stopped gauge
rdsosmetrics_tasks_stopped{instance="rds-mysql57",region="us-east-1"} 0
# HELP rdsosmetrics_tasks_total The total number of tasks.
# TYPE rdsosmetrics_tasks_total gauge
rdsosmetrics_tasks_total{instance="rds-mysql57",region="us-east-1"} 313
# HELP rdsosmetrics_tasks_zombie The number of child tasks that are inactive with an active parent task.
# TYPE rdsosmetrics_tasks_zombie gauge
rdsosmetrics_tasks_zombie{instance="rds-mysql57",region="us-east-1"} 0
	`)
)
