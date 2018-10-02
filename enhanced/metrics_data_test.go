package enhanced

import (
	"strings"
)

//nolint:lll
var (
	dataMySQL57 = []byte(strings.TrimSpace(`
	{
		"engine": "MYSQL",
		"instanceID": "rds-mysql57",
		"instanceResourceID": "db-FE4Y2GIJU6UADBOXKULV3DBATY",
		"timestamp": "2018-09-25T08:07:03Z",
		"version": 1.00,
		"uptime": "235 days, 17:17:40",
		"numVCPUs": 1,
		"cpuUtilization": {
			"guest": 0.00,
			"irq": 0.10,
			"system": 1.28,
			"wait": 56.35,
			"idle": 38.95,
			"user": 0.57,
			"total": 61.04,
			"steal": 0.33,
			"nice": 2.41
		},
		"loadAverageMinute": {
			"fifteen": 2.58,
			"five": 1.66,
			"one": 1.61
		},
		"memory": {
			"writeback": 0,
			"hugePagesFree": 0,
			"hugePagesRsvd": 0,
			"hugePagesSurp": 0,
			"cached": 81640,
			"hugePagesSize": 2048,
			"free": 127036,
			"hugePagesTotal": 0,
			"inactive": 476104,
			"pageTables": 8940,
			"dirty": 164,
			"mapped": 17696,
			"active": 1325340,
			"total": 2051520,
			"slab": 53924,
			"buffers": 133408
		},
		"tasks": {
			"sleeping": 278,
			"zombie": 0,
			"running": 3,
			"stopped": 0,
			"total": 281,
			"blocked": 0
		},
		"swap": {
			"cached": 4656,
			"total": 4095996,
			"out": 0.00,
			"free": 3755528,
			"in": 0.00
		},
		"network": [
			{
				"interface": "eth0",
				"rx": 50482.42,
				"tx": 792009.65
			}
		],
		"diskIO": [
			{
				"writeKbPS": 1667.47,
				"readIOsPS": 190.98,
				"await": 4.68,
				"readKbPS": 3055.73,
				"rrqmPS": 0.00,
				"util": 60.63,
				"avgQueueLen": 96.96,
				"tps": 345.52,
				"readKb": 183344,
				"device": "rdsdev",
				"writeKb": 100048,
				"avgReqSz": 13.67,
				"wrqmPS": 0.00,
				"writeIOsPS": 154.53
			}
		],
		"fileSys": [
			{
				"used": 25990072,
				"name": "rdsfilesys",
				"usedFiles": 728,
				"usedFilePercent": 0.02,
				"maxFiles": 3932160,
				"mountPoint": "/rdsdbdata",
				"total": 61774768,
				"usedPercent": 42.07
			}
		],
		"processList": [
			{
				"vss": 2965052,
				"name": "mysqld",
				"tgid": 3603,
				"vmlimit": "unlimited",
				"parentID": 1,
				"memoryUsedPc": 66.63,
				"cpuUsedPc": 0.00,
				"id": 3603,
				"rss": 1366964
			},
			{
				"vss": 2965052,
				"name": "mysqld",
				"tgid": 3603,
				"vmlimit": "unlimited",
				"parentID": 1,
				"memoryUsedPc": 66.63,
				"cpuUsedPc": 0.22,
				"id": 23950,
				"rss": 1366964
			},
			{
				"vss": 2965052,
				"name": "mysqld",
				"tgid": 3603,
				"vmlimit": "unlimited",
				"parentID": 1,
				"memoryUsedPc": 66.63,
				"cpuUsedPc": 0.20,
				"id": 23953,
				"rss": 1366964
			},
			{
				"vss": 2965052,
				"name": "mysqld",
				"tgid": 3603,
				"vmlimit": "unlimited",
				"parentID": 1,
				"memoryUsedPc": 66.63,
				"cpuUsedPc": 0.23,
				"id": 23958,
				"rss": 1366964
			},
			{
				"vss": 2965052,
				"name": "mysqld",
				"tgid": 3603,
				"vmlimit": "unlimited",
				"parentID": 1,
				"memoryUsedPc": 66.63,
				"cpuUsedPc": 0.22,
				"id": 23990,
				"rss": 1366964
			},
			{
				"vss": 2965052,
				"name": "mysqld",
				"tgid": 3603,
				"vmlimit": "unlimited",
				"parentID": 1,
				"memoryUsedPc": 66.63,
				"cpuUsedPc": 0.20,
				"id": 24000,
				"rss": 1366964
			},
			{
				"vss": 2965052,
				"name": "mysqld",
				"tgid": 3603,
				"vmlimit": "unlimited",
				"parentID": 1,
				"memoryUsedPc": 66.63,
				"cpuUsedPc": 0.17,
				"id": 24014,
				"rss": 1366964
			},
			{
				"vss": 628412,
				"name": "OS processes",
				"tgid": 0,
				"vmlimit": "",
				"parentID": 0,
				"memoryUsedPc": 0.30,
				"cpuUsedPc": 0.16,
				"id": 0,
				"rss": 6324
			},
			{
				"vss": 2419092,
				"name": "RDS processes",
				"tgid": 0,
				"vmlimit": "",
				"parentID": 0,
				"memoryUsedPc": 10.77,
				"cpuUsedPc": 0.06,
				"id": 0,
				"rss": 220940
			}
		]
	}`))

	dataAurora57 = []byte(strings.TrimSpace(`
	{
		"engine": "Aurora",
		"instanceID": "rds-aurora57",
		"instanceResourceID": "db-CDBSN4EK5SMBQCSXI4UPZVF3W4",
		"timestamp": "2018-09-25T08:16:20Z",
		"version": 1.00,
		"uptime": "186 days, 16:05:22",
		"numVCPUs": 1,
		"cpuUtilization": {
			"guest": 0.00,
			"irq": 0.03,
			"system": 2.22,
			"wait": 0.00,
			"idle": 0.00,
			"user": 3.38,
			"total": 100.00,
			"steal": 44.14,
			"nice": 50.23
		},
		"loadAverageMinute": {
			"fifteen": 3.99,
			"five": 3.84,
			"one": 3.73
		},
		"memory": {
			"writeback": 0,
			"hugePagesFree": 2048,
			"hugePagesRsvd": 0,
			"hugePagesSurp": 0,
			"cached": 139212,
			"hugePagesSize": 2048,
			"free": 110072,
			"hugePagesTotal": 737280,
			"inactive": 98936,
			"pageTables": 6360,
			"dirty": 308,
			"mapped": 39616,
			"active": 1033284,
			"total": 2051524,
			"slab": 38164,
			"buffers": 89008
		},
		"tasks": {
			"sleeping": 254,
			"zombie": 0,
			"running": 6,
			"stopped": 0,
			"total": 260,
			"blocked": 0
		},
		"swap": {
			"cached": 0,
			"total": 0,
			"out": 0.00,
			"free": 0,
			"in": 0.00
		},
		"network": [
			{
				"interface": "eth0",
				"rx": 736.93,
				"tx": 5464.83
			}
		],
		"diskIO": [
			{
				"readLatency": 0.00,
				"writeLatency": 24.59,
				"writeThroughput": 551.58,
				"readThroughput": 0.00,
				"readIOsPS": 0.00,
				"diskQueueDepth": 0,
				"writeIOsPS": 1.67
			}
		],
		"fileSys": [
			{
				"used": 3904152,
				"name": "rdsfilesys",
				"usedFiles": 1575,
				"usedFilePercent": 0.08,
				"maxFiles": 2097152,
				"mountPoint": "/rdsdbdata",
				"total": 32892784,
				"usedPercent": 11.87
			}
		],
		"processList": [
			{
				"vss": 1362404,
				"name": "aurora",
				"tgid": 31606,
				"vmlimit": "unlimited",
				"parentID": 1,
				"memoryUsedPc": 14.26,
				"cpuUsedPc": 0.00,
				"id": 31606,
				"rss": 292536
			},
			{
				"vss": 693404,
				"name": "OS processes",
				"tgid": 0,
				"vmlimit": "",
				"parentID": 0,
				"memoryUsedPc": 0.46,
				"cpuUsedPc": 0.12,
				"id": 0,
				"rss": 9952
			},
			{
				"vss": 5656964,
				"name": "RDS processes",
				"tgid": 0,
				"vmlimit": "",
				"parentID": 0,
				"memoryUsedPc": 30.51,
				"cpuUsedPc": 91.11,
				"id": 0,
				"rss": 625684
			}
		]
	}`))

	dataMySQL57Expected = strings.TrimSpace(`
# HELP node_cpu_average The percentage of CPU utilization.
# TYPE node_cpu_average gauge
node_cpu_average{cpu="All",instance="rds-mysql57",mode="guest",region="us-east-1"} 0
node_cpu_average{cpu="All",instance="rds-mysql57",mode="idle",region="us-east-1"} 38.95
node_cpu_average{cpu="All",instance="rds-mysql57",mode="irq",region="us-east-1"} 0.1
node_cpu_average{cpu="All",instance="rds-mysql57",mode="nice",region="us-east-1"} 2.41
node_cpu_average{cpu="All",instance="rds-mysql57",mode="steal",region="us-east-1"} 0.33
node_cpu_average{cpu="All",instance="rds-mysql57",mode="system",region="us-east-1"} 1.28
node_cpu_average{cpu="All",instance="rds-mysql57",mode="total",region="us-east-1"} 61.04
node_cpu_average{cpu="All",instance="rds-mysql57",mode="user",region="us-east-1"} 0.57
node_cpu_average{cpu="All",instance="rds-mysql57",mode="wait",region="us-east-1"} 56.35
# HELP node_disk_bytes_read The total number of bytes read successfully.
# TYPE node_disk_bytes_read counter
node_disk_bytes_read{device="rdsdev",instance="rds-mysql57",region="us-east-1"} 1.87744256e+08
# HELP node_disk_bytes_written The total number of bytes written successfully.
# TYPE node_disk_bytes_written counter
node_disk_bytes_written{device="rdsdev",instance="rds-mysql57",region="us-east-1"} 1.02449152e+08
# HELP node_load1 The number of processes requesting CPU time over the last minute.
# TYPE node_load1 gauge
node_load1{instance="rds-mysql57",region="us-east-1"} 1.61
# HELP node_memory_Active Memory information field Active.
# TYPE node_memory_Active gauge
node_memory_Active{instance="rds-mysql57",region="us-east-1"} 1.35714816e+09
# HELP node_memory_Buffers Memory information field Buffers.
# TYPE node_memory_Buffers gauge
node_memory_Buffers{instance="rds-mysql57",region="us-east-1"} 1.36609792e+08
# HELP node_memory_Cached Memory information field Cached.
# TYPE node_memory_Cached gauge
node_memory_Cached{instance="rds-mysql57",region="us-east-1"} 8.359936e+07
# HELP node_memory_Dirty Memory information field Dirty.
# TYPE node_memory_Dirty gauge
node_memory_Dirty{instance="rds-mysql57",region="us-east-1"} 167936
# HELP node_memory_HugePages_Free Memory information field HugePages_Free.
# TYPE node_memory_HugePages_Free gauge
node_memory_HugePages_Free{instance="rds-mysql57",region="us-east-1"} 0
# HELP node_memory_HugePages_Rsvd Memory information field HugePages_Rsvd.
# TYPE node_memory_HugePages_Rsvd gauge
node_memory_HugePages_Rsvd{instance="rds-mysql57",region="us-east-1"} 0
# HELP node_memory_HugePages_Surp Memory information field HugePages_Surp.
# TYPE node_memory_HugePages_Surp gauge
node_memory_HugePages_Surp{instance="rds-mysql57",region="us-east-1"} 0
# HELP node_memory_HugePages_Total Memory information field HugePages_Total.
# TYPE node_memory_HugePages_Total gauge
node_memory_HugePages_Total{instance="rds-mysql57",region="us-east-1"} 0
# HELP node_memory_Hugepagesize Memory information field Hugepagesize.
# TYPE node_memory_Hugepagesize gauge
node_memory_Hugepagesize{instance="rds-mysql57",region="us-east-1"} 2.097152e+06
# HELP node_memory_Inactive Memory information field Inactive.
# TYPE node_memory_Inactive gauge
node_memory_Inactive{instance="rds-mysql57",region="us-east-1"} 4.87530496e+08
# HELP node_memory_Mapped Memory information field Mapped.
# TYPE node_memory_Mapped gauge
node_memory_Mapped{instance="rds-mysql57",region="us-east-1"} 1.8120704e+07
# HELP node_memory_MemFree Memory information field MemFree.
# TYPE node_memory_MemFree gauge
node_memory_MemFree{instance="rds-mysql57",region="us-east-1"} 1.30084864e+08
# HELP node_memory_MemTotal Memory information field MemTotal.
# TYPE node_memory_MemTotal gauge
node_memory_MemTotal{instance="rds-mysql57",region="us-east-1"} 2.10075648e+09
# HELP node_memory_PageTables Memory information field PageTables.
# TYPE node_memory_PageTables gauge
node_memory_PageTables{instance="rds-mysql57",region="us-east-1"} 9.15456e+06
# HELP node_memory_Slab Memory information field Slab.
# TYPE node_memory_Slab gauge
node_memory_Slab{instance="rds-mysql57",region="us-east-1"} 5.5218176e+07
# HELP node_memory_Writeback Memory information field Writeback.
# TYPE node_memory_Writeback gauge
node_memory_Writeback{instance="rds-mysql57",region="us-east-1"} 0
# HELP rdsosmetrics_cpuUtilization_guest The percentage of CPU in use by guest programs.
# TYPE rdsosmetrics_cpuUtilization_guest gauge
rdsosmetrics_cpuUtilization_guest{instance="rds-mysql57",region="us-east-1"} 0
# HELP rdsosmetrics_cpuUtilization_idle The percentage of CPU that is idle.
# TYPE rdsosmetrics_cpuUtilization_idle gauge
rdsosmetrics_cpuUtilization_idle{instance="rds-mysql57",region="us-east-1"} 38.95
# HELP rdsosmetrics_cpuUtilization_irq The percentage of CPU in use by software interrupts.
# TYPE rdsosmetrics_cpuUtilization_irq gauge
rdsosmetrics_cpuUtilization_irq{instance="rds-mysql57",region="us-east-1"} 0.1
# HELP rdsosmetrics_cpuUtilization_nice The percentage of CPU in use by programs running at lowest priority.
# TYPE rdsosmetrics_cpuUtilization_nice gauge
rdsosmetrics_cpuUtilization_nice{instance="rds-mysql57",region="us-east-1"} 2.41
# HELP rdsosmetrics_cpuUtilization_steal The percentage of CPU in use by other virtual machines.
# TYPE rdsosmetrics_cpuUtilization_steal gauge
rdsosmetrics_cpuUtilization_steal{instance="rds-mysql57",region="us-east-1"} 0.33
# HELP rdsosmetrics_cpuUtilization_system The percentage of CPU in use by the kernel.
# TYPE rdsosmetrics_cpuUtilization_system gauge
rdsosmetrics_cpuUtilization_system{instance="rds-mysql57",region="us-east-1"} 1.28
# HELP rdsosmetrics_cpuUtilization_total The total percentage of the CPU in use. This value includes the nice value.
# TYPE rdsosmetrics_cpuUtilization_total gauge
rdsosmetrics_cpuUtilization_total{instance="rds-mysql57",region="us-east-1"} 61.04
# HELP rdsosmetrics_cpuUtilization_user The percentage of CPU in use by user programs.
# TYPE rdsosmetrics_cpuUtilization_user gauge
rdsosmetrics_cpuUtilization_user{instance="rds-mysql57",region="us-east-1"} 0.57
# HELP rdsosmetrics_cpuUtilization_wait The percentage of CPU unused while waiting for I/O access.
# TYPE rdsosmetrics_cpuUtilization_wait gauge
rdsosmetrics_cpuUtilization_wait{instance="rds-mysql57",region="us-east-1"} 56.35
# HELP rdsosmetrics_diskIO_avgQueueLen The number of requests waiting in the I/O device's queue.
# TYPE rdsosmetrics_diskIO_avgQueueLen gauge
rdsosmetrics_diskIO_avgQueueLen{device="rdsdev",instance="rds-mysql57",region="us-east-1"} 96.96
# HELP rdsosmetrics_diskIO_avgReqSz The average request size, in kilobytes.
# TYPE rdsosmetrics_diskIO_avgReqSz gauge
rdsosmetrics_diskIO_avgReqSz{device="rdsdev",instance="rds-mysql57",region="us-east-1"} 13.67
# HELP rdsosmetrics_diskIO_await The number of milliseconds required to respond to requests, including queue time and service time.
# TYPE rdsosmetrics_diskIO_await gauge
rdsosmetrics_diskIO_await{device="rdsdev",instance="rds-mysql57",region="us-east-1"} 4.68
# HELP rdsosmetrics_diskIO_readIOsPS The number of read operations per second.
# TYPE rdsosmetrics_diskIO_readIOsPS gauge
rdsosmetrics_diskIO_readIOsPS{device="rdsdev",instance="rds-mysql57",region="us-east-1"} 190.98
# HELP rdsosmetrics_diskIO_readKb The total number of kilobytes read.
# TYPE rdsosmetrics_diskIO_readKb gauge
rdsosmetrics_diskIO_readKb{device="rdsdev",instance="rds-mysql57",region="us-east-1"} 183344
# HELP rdsosmetrics_diskIO_readKbPS The number of kilobytes read per second.
# TYPE rdsosmetrics_diskIO_readKbPS gauge
rdsosmetrics_diskIO_readKbPS{device="rdsdev",instance="rds-mysql57",region="us-east-1"} 3055.73
# HELP rdsosmetrics_diskIO_rrqmPS The number of merged read requests queued per second.
# TYPE rdsosmetrics_diskIO_rrqmPS gauge
rdsosmetrics_diskIO_rrqmPS{device="rdsdev",instance="rds-mysql57",region="us-east-1"} 0
# HELP rdsosmetrics_diskIO_tps The number of I/O transactions per second.
# TYPE rdsosmetrics_diskIO_tps gauge
rdsosmetrics_diskIO_tps{device="rdsdev",instance="rds-mysql57",region="us-east-1"} 345.52
# HELP rdsosmetrics_diskIO_util The percentage of CPU time during which requests were issued.
# TYPE rdsosmetrics_diskIO_util gauge
rdsosmetrics_diskIO_util{device="rdsdev",instance="rds-mysql57",region="us-east-1"} 60.63
# HELP rdsosmetrics_diskIO_writeIOsPS The number of write operations per second.
# TYPE rdsosmetrics_diskIO_writeIOsPS gauge
rdsosmetrics_diskIO_writeIOsPS{device="rdsdev",instance="rds-mysql57",region="us-east-1"} 154.53
# HELP rdsosmetrics_diskIO_writeKb The total number of kilobytes written.
# TYPE rdsosmetrics_diskIO_writeKb gauge
rdsosmetrics_diskIO_writeKb{device="rdsdev",instance="rds-mysql57",region="us-east-1"} 100048
# HELP rdsosmetrics_diskIO_writeKbPS The number of kilobytes written per second.
# TYPE rdsosmetrics_diskIO_writeKbPS gauge
rdsosmetrics_diskIO_writeKbPS{device="rdsdev",instance="rds-mysql57",region="us-east-1"} 1667.47
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
rdsosmetrics_fileSys_used{instance="rds-mysql57",mount_point="/rdsdbdata",name="rdsfilesys",region="us-east-1"} 2.5990072e+07
# HELP rdsosmetrics_fileSys_usedFilePercent The percentage of available files in use.
# TYPE rdsosmetrics_fileSys_usedFilePercent gauge
rdsosmetrics_fileSys_usedFilePercent{instance="rds-mysql57",mount_point="/rdsdbdata",name="rdsfilesys",region="us-east-1"} 0.02
# HELP rdsosmetrics_fileSys_usedFiles The number of files in the file system.
# TYPE rdsosmetrics_fileSys_usedFiles gauge
rdsosmetrics_fileSys_usedFiles{instance="rds-mysql57",mount_point="/rdsdbdata",name="rdsfilesys",region="us-east-1"} 728
# HELP rdsosmetrics_fileSys_usedPercent The percentage of the file-system disk space in use.
# TYPE rdsosmetrics_fileSys_usedPercent gauge
rdsosmetrics_fileSys_usedPercent{instance="rds-mysql57",mount_point="/rdsdbdata",name="rdsfilesys",region="us-east-1"} 42.07
# HELP rdsosmetrics_loadAverageMinute_fifteen The number of processes requesting CPU time over the last 15 minutes.
# TYPE rdsosmetrics_loadAverageMinute_fifteen gauge
rdsosmetrics_loadAverageMinute_fifteen{instance="rds-mysql57",region="us-east-1"} 2.58
# HELP rdsosmetrics_loadAverageMinute_five The number of processes requesting CPU time over the last 5 minutes.
# TYPE rdsosmetrics_loadAverageMinute_five gauge
rdsosmetrics_loadAverageMinute_five{instance="rds-mysql57",region="us-east-1"} 1.66
# HELP rdsosmetrics_loadAverageMinute_one The number of processes requesting CPU time over the last minute.
# TYPE rdsosmetrics_loadAverageMinute_one gauge
rdsosmetrics_loadAverageMinute_one{instance="rds-mysql57",region="us-east-1"} 1.61
# HELP rdsosmetrics_memory_active The amount of assigned memory, in kilobytes.
# TYPE rdsosmetrics_memory_active gauge
rdsosmetrics_memory_active{instance="rds-mysql57",region="us-east-1"} 1.32534e+06
# HELP rdsosmetrics_memory_buffers The amount of memory used for buffering I/O requests prior to writing to the storage device, in kilobytes.
# TYPE rdsosmetrics_memory_buffers gauge
rdsosmetrics_memory_buffers{instance="rds-mysql57",region="us-east-1"} 133408
# HELP rdsosmetrics_memory_cached The amount of memory used for caching file system–based I/O.
# TYPE rdsosmetrics_memory_cached gauge
rdsosmetrics_memory_cached{instance="rds-mysql57",region="us-east-1"} 81640
# HELP rdsosmetrics_memory_dirty The amount of memory pages in RAM that have been modified but not written to their related data block in storage, in kilobytes.
# TYPE rdsosmetrics_memory_dirty gauge
rdsosmetrics_memory_dirty{instance="rds-mysql57",region="us-east-1"} 164
# HELP rdsosmetrics_memory_free The amount of unassigned memory, in kilobytes.
# TYPE rdsosmetrics_memory_free gauge
rdsosmetrics_memory_free{instance="rds-mysql57",region="us-east-1"} 127036
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
rdsosmetrics_memory_inactive{instance="rds-mysql57",region="us-east-1"} 476104
# HELP rdsosmetrics_memory_mapped The total amount of file-system contents that is memory mapped inside a process address space, in kilobytes.
# TYPE rdsosmetrics_memory_mapped gauge
rdsosmetrics_memory_mapped{instance="rds-mysql57",region="us-east-1"} 17696
# HELP rdsosmetrics_memory_pageTables The amount of memory used by page tables, in kilobytes.
# TYPE rdsosmetrics_memory_pageTables gauge
rdsosmetrics_memory_pageTables{instance="rds-mysql57",region="us-east-1"} 8940
# HELP rdsosmetrics_memory_slab The amount of reusable kernel data structures, in kilobytes.
# TYPE rdsosmetrics_memory_slab gauge
rdsosmetrics_memory_slab{instance="rds-mysql57",region="us-east-1"} 53924
# HELP rdsosmetrics_memory_total The total amount of memory, in kilobytes.
# TYPE rdsosmetrics_memory_total gauge
rdsosmetrics_memory_total{instance="rds-mysql57",region="us-east-1"} 2.05152e+06
# HELP rdsosmetrics_memory_writeback The amount of dirty pages in RAM that are still being written to the backing storage, in kilobytes.
# TYPE rdsosmetrics_memory_writeback gauge
rdsosmetrics_memory_writeback{instance="rds-mysql57",region="us-east-1"} 0
# HELP rdsosmetrics_network_rx The number of bytes received per second.
# TYPE rdsosmetrics_network_rx gauge
rdsosmetrics_network_rx{instance="rds-mysql57",interface="eth0",region="us-east-1"} 50482.42
# HELP rdsosmetrics_network_tx The number of bytes uploaded per second.
# TYPE rdsosmetrics_network_tx gauge
rdsosmetrics_network_tx{instance="rds-mysql57",interface="eth0",region="us-east-1"} 792009.65
# HELP rdsosmetrics_processList_cpuUsedPc The percentage of CPU used by the process.
# TYPE rdsosmetrics_processList_cpuUsedPc gauge
rdsosmetrics_processList_cpuUsedPc{id="0",instance="rds-mysql57",name="OS processes",parentID="0",region="us-east-1",tgid="0"} 0.16
rdsosmetrics_processList_cpuUsedPc{id="0",instance="rds-mysql57",name="RDS processes",parentID="0",region="us-east-1",tgid="0"} 0.06
rdsosmetrics_processList_cpuUsedPc{id="23950",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="3603"} 0.22
rdsosmetrics_processList_cpuUsedPc{id="23953",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="3603"} 0.2
rdsosmetrics_processList_cpuUsedPc{id="23958",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="3603"} 0.23
rdsosmetrics_processList_cpuUsedPc{id="23990",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="3603"} 0.22
rdsosmetrics_processList_cpuUsedPc{id="24000",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="3603"} 0.2
rdsosmetrics_processList_cpuUsedPc{id="24014",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="3603"} 0.17
rdsosmetrics_processList_cpuUsedPc{id="3603",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="3603"} 0
# HELP rdsosmetrics_processList_memoryUsedPc The amount of memory used by the process, in kilobytes.
# TYPE rdsosmetrics_processList_memoryUsedPc gauge
rdsosmetrics_processList_memoryUsedPc{id="0",instance="rds-mysql57",name="OS processes",parentID="0",region="us-east-1",tgid="0"} 0.3
rdsosmetrics_processList_memoryUsedPc{id="0",instance="rds-mysql57",name="RDS processes",parentID="0",region="us-east-1",tgid="0"} 10.77
rdsosmetrics_processList_memoryUsedPc{id="23950",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="3603"} 66.63
rdsosmetrics_processList_memoryUsedPc{id="23953",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="3603"} 66.63
rdsosmetrics_processList_memoryUsedPc{id="23958",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="3603"} 66.63
rdsosmetrics_processList_memoryUsedPc{id="23990",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="3603"} 66.63
rdsosmetrics_processList_memoryUsedPc{id="24000",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="3603"} 66.63
rdsosmetrics_processList_memoryUsedPc{id="24014",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="3603"} 66.63
rdsosmetrics_processList_memoryUsedPc{id="3603",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="3603"} 66.63
# HELP rdsosmetrics_processList_rss The amount of RAM allocated to the process, in kilobytes.
# TYPE rdsosmetrics_processList_rss gauge
rdsosmetrics_processList_rss{id="0",instance="rds-mysql57",name="OS processes",parentID="0",region="us-east-1",tgid="0"} 6324
rdsosmetrics_processList_rss{id="0",instance="rds-mysql57",name="RDS processes",parentID="0",region="us-east-1",tgid="0"} 220940
rdsosmetrics_processList_rss{id="23950",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="3603"} 1.366964e+06
rdsosmetrics_processList_rss{id="23953",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="3603"} 1.366964e+06
rdsosmetrics_processList_rss{id="23958",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="3603"} 1.366964e+06
rdsosmetrics_processList_rss{id="23990",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="3603"} 1.366964e+06
rdsosmetrics_processList_rss{id="24000",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="3603"} 1.366964e+06
rdsosmetrics_processList_rss{id="24014",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="3603"} 1.366964e+06
rdsosmetrics_processList_rss{id="3603",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="3603"} 1.366964e+06
# HELP rdsosmetrics_processList_vss The amount of virtual memory allocated to the process, in kilobytes.
# TYPE rdsosmetrics_processList_vss gauge
rdsosmetrics_processList_vss{id="0",instance="rds-mysql57",name="OS processes",parentID="0",region="us-east-1",tgid="0"} 628412
rdsosmetrics_processList_vss{id="0",instance="rds-mysql57",name="RDS processes",parentID="0",region="us-east-1",tgid="0"} 2.419092e+06
rdsosmetrics_processList_vss{id="23950",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="3603"} 2.965052e+06
rdsosmetrics_processList_vss{id="23953",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="3603"} 2.965052e+06
rdsosmetrics_processList_vss{id="23958",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="3603"} 2.965052e+06
rdsosmetrics_processList_vss{id="23990",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="3603"} 2.965052e+06
rdsosmetrics_processList_vss{id="24000",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="3603"} 2.965052e+06
rdsosmetrics_processList_vss{id="24014",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="3603"} 2.965052e+06
rdsosmetrics_processList_vss{id="3603",instance="rds-mysql57",name="mysqld",parentID="1",region="us-east-1",tgid="3603"} 2.965052e+06
# HELP rdsosmetrics_swap_cached The amount of swap memory, in kilobytes, used as cache memory.
# TYPE rdsosmetrics_swap_cached gauge
rdsosmetrics_swap_cached{instance="rds-mysql57",region="us-east-1"} 4656
# HELP rdsosmetrics_swap_free The total amount of swap memory free, in kilobytes.
# TYPE rdsosmetrics_swap_free gauge
rdsosmetrics_swap_free{instance="rds-mysql57",region="us-east-1"} 3.755528e+06
# HELP rdsosmetrics_swap_total The total amount of swap memory available, in kilobytes.
# TYPE rdsosmetrics_swap_total gauge
rdsosmetrics_swap_total{instance="rds-mysql57",region="us-east-1"} 4.095996e+06
# HELP rdsosmetrics_tasks_blocked The number of tasks that are blocked.
# TYPE rdsosmetrics_tasks_blocked gauge
rdsosmetrics_tasks_blocked{instance="rds-mysql57",region="us-east-1"} 0
# HELP rdsosmetrics_tasks_running The number of tasks that are running.
# TYPE rdsosmetrics_tasks_running gauge
rdsosmetrics_tasks_running{instance="rds-mysql57",region="us-east-1"} 3
# HELP rdsosmetrics_tasks_sleeping The number of tasks that are sleeping.
# TYPE rdsosmetrics_tasks_sleeping gauge
rdsosmetrics_tasks_sleeping{instance="rds-mysql57",region="us-east-1"} 278
# HELP rdsosmetrics_tasks_stopped The number of tasks that are stopped.
# TYPE rdsosmetrics_tasks_stopped gauge
rdsosmetrics_tasks_stopped{instance="rds-mysql57",region="us-east-1"} 0
# HELP rdsosmetrics_tasks_total The total number of tasks.
# TYPE rdsosmetrics_tasks_total gauge
rdsosmetrics_tasks_total{instance="rds-mysql57",region="us-east-1"} 281
# HELP rdsosmetrics_tasks_zombie The number of child tasks that are inactive with an active parent task.
# TYPE rdsosmetrics_tasks_zombie gauge
rdsosmetrics_tasks_zombie{instance="rds-mysql57",region="us-east-1"} 0
# HELP rdsosmetrics_timestamp Metrics timestamp (UNIX seconds).
# TYPE rdsosmetrics_timestamp counter
rdsosmetrics_timestamp{instance="rds-mysql57",region="us-east-1"} 1.537862823e+09
	`)

	dataAurora57Expected = strings.TrimSpace(`
# HELP node_cpu_average The percentage of CPU utilization.
# TYPE node_cpu_average gauge
node_cpu_average{cpu="All",instance="rds-aurora57",mode="guest",region="us-east-1"} 0
node_cpu_average{cpu="All",instance="rds-aurora57",mode="idle",region="us-east-1"} 0
node_cpu_average{cpu="All",instance="rds-aurora57",mode="irq",region="us-east-1"} 0.03
node_cpu_average{cpu="All",instance="rds-aurora57",mode="nice",region="us-east-1"} 50.23
node_cpu_average{cpu="All",instance="rds-aurora57",mode="steal",region="us-east-1"} 44.14
node_cpu_average{cpu="All",instance="rds-aurora57",mode="system",region="us-east-1"} 2.22
node_cpu_average{cpu="All",instance="rds-aurora57",mode="total",region="us-east-1"} 100
node_cpu_average{cpu="All",instance="rds-aurora57",mode="user",region="us-east-1"} 3.38
node_cpu_average{cpu="All",instance="rds-aurora57",mode="wait",region="us-east-1"} 0
# HELP node_load1 The number of processes requesting CPU time over the last minute.
# TYPE node_load1 gauge
node_load1{instance="rds-aurora57",region="us-east-1"} 3.73
# HELP node_memory_Active Memory information field Active.
# TYPE node_memory_Active gauge
node_memory_Active{instance="rds-aurora57",region="us-east-1"} 1.058082816e+09
# HELP node_memory_Buffers Memory information field Buffers.
# TYPE node_memory_Buffers gauge
node_memory_Buffers{instance="rds-aurora57",region="us-east-1"} 9.1144192e+07
# HELP node_memory_Cached Memory information field Cached.
# TYPE node_memory_Cached gauge
node_memory_Cached{instance="rds-aurora57",region="us-east-1"} 1.42553088e+08
# HELP node_memory_Dirty Memory information field Dirty.
# TYPE node_memory_Dirty gauge
node_memory_Dirty{instance="rds-aurora57",region="us-east-1"} 315392
# HELP node_memory_HugePages_Free Memory information field HugePages_Free.
# TYPE node_memory_HugePages_Free gauge
node_memory_HugePages_Free{instance="rds-aurora57",region="us-east-1"} 2048
# HELP node_memory_HugePages_Rsvd Memory information field HugePages_Rsvd.
# TYPE node_memory_HugePages_Rsvd gauge
node_memory_HugePages_Rsvd{instance="rds-aurora57",region="us-east-1"} 0
# HELP node_memory_HugePages_Surp Memory information field HugePages_Surp.
# TYPE node_memory_HugePages_Surp gauge
node_memory_HugePages_Surp{instance="rds-aurora57",region="us-east-1"} 0
# HELP node_memory_HugePages_Total Memory information field HugePages_Total.
# TYPE node_memory_HugePages_Total gauge
node_memory_HugePages_Total{instance="rds-aurora57",region="us-east-1"} 737280
# HELP node_memory_Hugepagesize Memory information field Hugepagesize.
# TYPE node_memory_Hugepagesize gauge
node_memory_Hugepagesize{instance="rds-aurora57",region="us-east-1"} 2.097152e+06
# HELP node_memory_Inactive Memory information field Inactive.
# TYPE node_memory_Inactive gauge
node_memory_Inactive{instance="rds-aurora57",region="us-east-1"} 1.01310464e+08
# HELP node_memory_Mapped Memory information field Mapped.
# TYPE node_memory_Mapped gauge
node_memory_Mapped{instance="rds-aurora57",region="us-east-1"} 4.0566784e+07
# HELP node_memory_MemFree Memory information field MemFree.
# TYPE node_memory_MemFree gauge
node_memory_MemFree{instance="rds-aurora57",region="us-east-1"} 1.12713728e+08
# HELP node_memory_MemTotal Memory information field MemTotal.
# TYPE node_memory_MemTotal gauge
node_memory_MemTotal{instance="rds-aurora57",region="us-east-1"} 2.100760576e+09
# HELP node_memory_PageTables Memory information field PageTables.
# TYPE node_memory_PageTables gauge
node_memory_PageTables{instance="rds-aurora57",region="us-east-1"} 6.51264e+06
# HELP node_memory_Slab Memory information field Slab.
# TYPE node_memory_Slab gauge
node_memory_Slab{instance="rds-aurora57",region="us-east-1"} 3.9079936e+07
# HELP node_memory_Writeback Memory information field Writeback.
# TYPE node_memory_Writeback gauge
node_memory_Writeback{instance="rds-aurora57",region="us-east-1"} 0
# HELP rdsosmetrics_cpuUtilization_guest The percentage of CPU in use by guest programs.
# TYPE rdsosmetrics_cpuUtilization_guest gauge
rdsosmetrics_cpuUtilization_guest{instance="rds-aurora57",region="us-east-1"} 0
# HELP rdsosmetrics_cpuUtilization_idle The percentage of CPU that is idle.
# TYPE rdsosmetrics_cpuUtilization_idle gauge
rdsosmetrics_cpuUtilization_idle{instance="rds-aurora57",region="us-east-1"} 0
# HELP rdsosmetrics_cpuUtilization_irq The percentage of CPU in use by software interrupts.
# TYPE rdsosmetrics_cpuUtilization_irq gauge
rdsosmetrics_cpuUtilization_irq{instance="rds-aurora57",region="us-east-1"} 0.03
# HELP rdsosmetrics_cpuUtilization_nice The percentage of CPU in use by programs running at lowest priority.
# TYPE rdsosmetrics_cpuUtilization_nice gauge
rdsosmetrics_cpuUtilization_nice{instance="rds-aurora57",region="us-east-1"} 50.23
# HELP rdsosmetrics_cpuUtilization_steal The percentage of CPU in use by other virtual machines.
# TYPE rdsosmetrics_cpuUtilization_steal gauge
rdsosmetrics_cpuUtilization_steal{instance="rds-aurora57",region="us-east-1"} 44.14
# HELP rdsosmetrics_cpuUtilization_system The percentage of CPU in use by the kernel.
# TYPE rdsosmetrics_cpuUtilization_system gauge
rdsosmetrics_cpuUtilization_system{instance="rds-aurora57",region="us-east-1"} 2.22
# HELP rdsosmetrics_cpuUtilization_total The total percentage of the CPU in use. This value includes the nice value.
# TYPE rdsosmetrics_cpuUtilization_total gauge
rdsosmetrics_cpuUtilization_total{instance="rds-aurora57",region="us-east-1"} 100
# HELP rdsosmetrics_cpuUtilization_user The percentage of CPU in use by user programs.
# TYPE rdsosmetrics_cpuUtilization_user gauge
rdsosmetrics_cpuUtilization_user{instance="rds-aurora57",region="us-east-1"} 3.38
# HELP rdsosmetrics_cpuUtilization_wait The percentage of CPU unused while waiting for I/O access.
# TYPE rdsosmetrics_cpuUtilization_wait gauge
rdsosmetrics_cpuUtilization_wait{instance="rds-aurora57",region="us-east-1"} 0
# HELP rdsosmetrics_diskIO_diskQueueDepth The number of outstanding IOs (read/write requests) waiting to access the disk.
# TYPE rdsosmetrics_diskIO_diskQueueDepth gauge
rdsosmetrics_diskIO_diskQueueDepth{device="",instance="rds-aurora57",region="us-east-1"} 0
# HELP rdsosmetrics_diskIO_readIOsPS The number of read operations per second.
# TYPE rdsosmetrics_diskIO_readIOsPS gauge
rdsosmetrics_diskIO_readIOsPS{device="",instance="rds-aurora57",region="us-east-1"} 0
# HELP rdsosmetrics_diskIO_readLatency The average amount of time taken per disk I/O operation.
# TYPE rdsosmetrics_diskIO_readLatency gauge
rdsosmetrics_diskIO_readLatency{device="",instance="rds-aurora57",region="us-east-1"} 0
# HELP rdsosmetrics_diskIO_readThroughput The average number of bytes read from disk per second.
# TYPE rdsosmetrics_diskIO_readThroughput gauge
rdsosmetrics_diskIO_readThroughput{device="",instance="rds-aurora57",region="us-east-1"} 0
# HELP rdsosmetrics_diskIO_writeIOsPS The number of write operations per second.
# TYPE rdsosmetrics_diskIO_writeIOsPS gauge
rdsosmetrics_diskIO_writeIOsPS{device="",instance="rds-aurora57",region="us-east-1"} 1.67
# HELP rdsosmetrics_diskIO_writeLatency The average amount of time taken per disk I/O operation.
# TYPE rdsosmetrics_diskIO_writeLatency gauge
rdsosmetrics_diskIO_writeLatency{device="",instance="rds-aurora57",region="us-east-1"} 24.59
# HELP rdsosmetrics_diskIO_writeThroughput The average number of bytes written to disk per second.
# TYPE rdsosmetrics_diskIO_writeThroughput gauge
rdsosmetrics_diskIO_writeThroughput{device="",instance="rds-aurora57",region="us-east-1"} 551.58
# HELP rdsosmetrics_fileSys_maxFiles The maximum number of files that can be created for the file system.
# TYPE rdsosmetrics_fileSys_maxFiles gauge
rdsosmetrics_fileSys_maxFiles{instance="rds-aurora57",mount_point="/rdsdbdata",name="rdsfilesys",region="us-east-1"} 2.097152e+06
# HELP rdsosmetrics_fileSys_total The total number of disk space available for the file system, in kilobytes.
# TYPE rdsosmetrics_fileSys_total gauge
rdsosmetrics_fileSys_total{instance="rds-aurora57",mount_point="/rdsdbdata",name="rdsfilesys",region="us-east-1"} 3.2892784e+07
# HELP rdsosmetrics_fileSys_used The amount of disk space used by files in the file system, in kilobytes.
# TYPE rdsosmetrics_fileSys_used gauge
rdsosmetrics_fileSys_used{instance="rds-aurora57",mount_point="/rdsdbdata",name="rdsfilesys",region="us-east-1"} 3.904152e+06
# HELP rdsosmetrics_fileSys_usedFilePercent The percentage of available files in use.
# TYPE rdsosmetrics_fileSys_usedFilePercent gauge
rdsosmetrics_fileSys_usedFilePercent{instance="rds-aurora57",mount_point="/rdsdbdata",name="rdsfilesys",region="us-east-1"} 0.08
# HELP rdsosmetrics_fileSys_usedFiles The number of files in the file system.
# TYPE rdsosmetrics_fileSys_usedFiles gauge
rdsosmetrics_fileSys_usedFiles{instance="rds-aurora57",mount_point="/rdsdbdata",name="rdsfilesys",region="us-east-1"} 1575
# HELP rdsosmetrics_fileSys_usedPercent The percentage of the file-system disk space in use.
# TYPE rdsosmetrics_fileSys_usedPercent gauge
rdsosmetrics_fileSys_usedPercent{instance="rds-aurora57",mount_point="/rdsdbdata",name="rdsfilesys",region="us-east-1"} 11.87
# HELP rdsosmetrics_loadAverageMinute_fifteen The number of processes requesting CPU time over the last 15 minutes.
# TYPE rdsosmetrics_loadAverageMinute_fifteen gauge
rdsosmetrics_loadAverageMinute_fifteen{instance="rds-aurora57",region="us-east-1"} 3.99
# HELP rdsosmetrics_loadAverageMinute_five The number of processes requesting CPU time over the last 5 minutes.
# TYPE rdsosmetrics_loadAverageMinute_five gauge
rdsosmetrics_loadAverageMinute_five{instance="rds-aurora57",region="us-east-1"} 3.84
# HELP rdsosmetrics_loadAverageMinute_one The number of processes requesting CPU time over the last minute.
# TYPE rdsosmetrics_loadAverageMinute_one gauge
rdsosmetrics_loadAverageMinute_one{instance="rds-aurora57",region="us-east-1"} 3.73
# HELP rdsosmetrics_memory_active The amount of assigned memory, in kilobytes.
# TYPE rdsosmetrics_memory_active gauge
rdsosmetrics_memory_active{instance="rds-aurora57",region="us-east-1"} 1.033284e+06
# HELP rdsosmetrics_memory_buffers The amount of memory used for buffering I/O requests prior to writing to the storage device, in kilobytes.
# TYPE rdsosmetrics_memory_buffers gauge
rdsosmetrics_memory_buffers{instance="rds-aurora57",region="us-east-1"} 89008
# HELP rdsosmetrics_memory_cached The amount of memory used for caching file system–based I/O.
# TYPE rdsosmetrics_memory_cached gauge
rdsosmetrics_memory_cached{instance="rds-aurora57",region="us-east-1"} 139212
# HELP rdsosmetrics_memory_dirty The amount of memory pages in RAM that have been modified but not written to their related data block in storage, in kilobytes.
# TYPE rdsosmetrics_memory_dirty gauge
rdsosmetrics_memory_dirty{instance="rds-aurora57",region="us-east-1"} 308
# HELP rdsosmetrics_memory_free The amount of unassigned memory, in kilobytes.
# TYPE rdsosmetrics_memory_free gauge
rdsosmetrics_memory_free{instance="rds-aurora57",region="us-east-1"} 110072
# HELP rdsosmetrics_memory_hugePagesFree The number of free huge pages. Huge pages are a feature of the Linux kernel.
# TYPE rdsosmetrics_memory_hugePagesFree gauge
rdsosmetrics_memory_hugePagesFree{instance="rds-aurora57",region="us-east-1"} 2048
# HELP rdsosmetrics_memory_hugePagesRsvd The number of committed huge pages.
# TYPE rdsosmetrics_memory_hugePagesRsvd gauge
rdsosmetrics_memory_hugePagesRsvd{instance="rds-aurora57",region="us-east-1"} 0
# HELP rdsosmetrics_memory_hugePagesSize The size for each huge pages unit, in kilobytes.
# TYPE rdsosmetrics_memory_hugePagesSize gauge
rdsosmetrics_memory_hugePagesSize{instance="rds-aurora57",region="us-east-1"} 2048
# HELP rdsosmetrics_memory_hugePagesSurp The number of available surplus huge pages over the total.
# TYPE rdsosmetrics_memory_hugePagesSurp gauge
rdsosmetrics_memory_hugePagesSurp{instance="rds-aurora57",region="us-east-1"} 0
# HELP rdsosmetrics_memory_hugePagesTotal The total number of huge pages for the system.
# TYPE rdsosmetrics_memory_hugePagesTotal gauge
rdsosmetrics_memory_hugePagesTotal{instance="rds-aurora57",region="us-east-1"} 737280
# HELP rdsosmetrics_memory_inactive The amount of least-frequently used memory pages, in kilobytes.
# TYPE rdsosmetrics_memory_inactive gauge
rdsosmetrics_memory_inactive{instance="rds-aurora57",region="us-east-1"} 98936
# HELP rdsosmetrics_memory_mapped The total amount of file-system contents that is memory mapped inside a process address space, in kilobytes.
# TYPE rdsosmetrics_memory_mapped gauge
rdsosmetrics_memory_mapped{instance="rds-aurora57",region="us-east-1"} 39616
# HELP rdsosmetrics_memory_pageTables The amount of memory used by page tables, in kilobytes.
# TYPE rdsosmetrics_memory_pageTables gauge
rdsosmetrics_memory_pageTables{instance="rds-aurora57",region="us-east-1"} 6360
# HELP rdsosmetrics_memory_slab The amount of reusable kernel data structures, in kilobytes.
# TYPE rdsosmetrics_memory_slab gauge
rdsosmetrics_memory_slab{instance="rds-aurora57",region="us-east-1"} 38164
# HELP rdsosmetrics_memory_total The total amount of memory, in kilobytes.
# TYPE rdsosmetrics_memory_total gauge
rdsosmetrics_memory_total{instance="rds-aurora57",region="us-east-1"} 2.051524e+06
# HELP rdsosmetrics_memory_writeback The amount of dirty pages in RAM that are still being written to the backing storage, in kilobytes.
# TYPE rdsosmetrics_memory_writeback gauge
rdsosmetrics_memory_writeback{instance="rds-aurora57",region="us-east-1"} 0
# HELP rdsosmetrics_network_rx The number of bytes received per second.
# TYPE rdsosmetrics_network_rx gauge
rdsosmetrics_network_rx{instance="rds-aurora57",interface="eth0",region="us-east-1"} 736.93
# HELP rdsosmetrics_network_tx The number of bytes uploaded per second.
# TYPE rdsosmetrics_network_tx gauge
rdsosmetrics_network_tx{instance="rds-aurora57",interface="eth0",region="us-east-1"} 5464.83
# HELP rdsosmetrics_processList_cpuUsedPc The percentage of CPU used by the process.
# TYPE rdsosmetrics_processList_cpuUsedPc gauge
rdsosmetrics_processList_cpuUsedPc{id="0",instance="rds-aurora57",name="OS processes",parentID="0",region="us-east-1",tgid="0"} 0.12
rdsosmetrics_processList_cpuUsedPc{id="0",instance="rds-aurora57",name="RDS processes",parentID="0",region="us-east-1",tgid="0"} 91.11
rdsosmetrics_processList_cpuUsedPc{id="31606",instance="rds-aurora57",name="aurora",parentID="1",region="us-east-1",tgid="31606"} 0
# HELP rdsosmetrics_processList_memoryUsedPc The amount of memory used by the process, in kilobytes.
# TYPE rdsosmetrics_processList_memoryUsedPc gauge
rdsosmetrics_processList_memoryUsedPc{id="0",instance="rds-aurora57",name="OS processes",parentID="0",region="us-east-1",tgid="0"} 0.46
rdsosmetrics_processList_memoryUsedPc{id="0",instance="rds-aurora57",name="RDS processes",parentID="0",region="us-east-1",tgid="0"} 30.51
rdsosmetrics_processList_memoryUsedPc{id="31606",instance="rds-aurora57",name="aurora",parentID="1",region="us-east-1",tgid="31606"} 14.26
# HELP rdsosmetrics_processList_rss The amount of RAM allocated to the process, in kilobytes.
# TYPE rdsosmetrics_processList_rss gauge
rdsosmetrics_processList_rss{id="0",instance="rds-aurora57",name="OS processes",parentID="0",region="us-east-1",tgid="0"} 9952
rdsosmetrics_processList_rss{id="0",instance="rds-aurora57",name="RDS processes",parentID="0",region="us-east-1",tgid="0"} 625684
rdsosmetrics_processList_rss{id="31606",instance="rds-aurora57",name="aurora",parentID="1",region="us-east-1",tgid="31606"} 292536
# HELP rdsosmetrics_processList_vss The amount of virtual memory allocated to the process, in kilobytes.
# TYPE rdsosmetrics_processList_vss gauge
rdsosmetrics_processList_vss{id="0",instance="rds-aurora57",name="OS processes",parentID="0",region="us-east-1",tgid="0"} 693404
rdsosmetrics_processList_vss{id="0",instance="rds-aurora57",name="RDS processes",parentID="0",region="us-east-1",tgid="0"} 5.656964e+06
rdsosmetrics_processList_vss{id="31606",instance="rds-aurora57",name="aurora",parentID="1",region="us-east-1",tgid="31606"} 1.362404e+06
# HELP rdsosmetrics_swap_cached The amount of swap memory, in kilobytes, used as cache memory.
# TYPE rdsosmetrics_swap_cached gauge
rdsosmetrics_swap_cached{instance="rds-aurora57",region="us-east-1"} 0
# HELP rdsosmetrics_swap_free The total amount of swap memory free, in kilobytes.
# TYPE rdsosmetrics_swap_free gauge
rdsosmetrics_swap_free{instance="rds-aurora57",region="us-east-1"} 0
# HELP rdsosmetrics_swap_total The total amount of swap memory available, in kilobytes.
# TYPE rdsosmetrics_swap_total gauge
rdsosmetrics_swap_total{instance="rds-aurora57",region="us-east-1"} 0
# HELP rdsosmetrics_tasks_blocked The number of tasks that are blocked.
# TYPE rdsosmetrics_tasks_blocked gauge
rdsosmetrics_tasks_blocked{instance="rds-aurora57",region="us-east-1"} 0
# HELP rdsosmetrics_tasks_running The number of tasks that are running.
# TYPE rdsosmetrics_tasks_running gauge
rdsosmetrics_tasks_running{instance="rds-aurora57",region="us-east-1"} 6
# HELP rdsosmetrics_tasks_sleeping The number of tasks that are sleeping.
# TYPE rdsosmetrics_tasks_sleeping gauge
rdsosmetrics_tasks_sleeping{instance="rds-aurora57",region="us-east-1"} 254
# HELP rdsosmetrics_tasks_stopped The number of tasks that are stopped.
# TYPE rdsosmetrics_tasks_stopped gauge
rdsosmetrics_tasks_stopped{instance="rds-aurora57",region="us-east-1"} 0
# HELP rdsosmetrics_tasks_total The total number of tasks.
# TYPE rdsosmetrics_tasks_total gauge
rdsosmetrics_tasks_total{instance="rds-aurora57",region="us-east-1"} 260
# HELP rdsosmetrics_tasks_zombie The number of child tasks that are inactive with an active parent task.
# TYPE rdsosmetrics_tasks_zombie gauge
rdsosmetrics_tasks_zombie{instance="rds-aurora57",region="us-east-1"} 0
# HELP rdsosmetrics_timestamp Metrics timestamp (UNIX seconds).
# TYPE rdsosmetrics_timestamp counter
rdsosmetrics_timestamp{instance="rds-aurora57",region="us-east-1"} 1.53786338e+09
	`)
)
