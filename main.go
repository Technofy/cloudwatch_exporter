package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/common/log"
	"github.com/prometheus/common/version"

	"github.com/percona/rds_exporter/basic"
	"github.com/percona/rds_exporter/config"
	"github.com/percona/rds_exporter/enhanced"
	"github.com/percona/rds_exporter/sessions"
)

var (
	listenAddress       = flag.String("web.listen-address", ":9042", "Address on which to expose metrics and web interface.")
	basicMetricsPath    = flag.String("web.basic-telemetry-path", "/basic", "Path under which to expose exporter's basic metrics.")
	enhancedMetricsPath = flag.String("web.enhanced-telemetry-path", "/enhanced", "Path under which to expose exporter's enhanced metrics.")
	configFile          = flag.String("config.file", "config.yml", "Path to configuration file.")

	settings     *config.Settings
	sessionsPool *sessions.Sessions
)

// handleReload handles a full reload of the configuration file and regenerates the collector templates.
func handleReload(w http.ResponseWriter, req *http.Request) {
	err := settings.Load(*configFile)
	if err != nil {
		str := fmt.Sprintf("Can't read configuration file: %s", err.Error())
		fmt.Fprintln(w, str)
		log.Errorln(str)
	}
	fmt.Fprintln(w, "Reload complete")
}

func main() {
	log.Infoln("Starting RDS exporter", version.Info())
	log.Infoln("Build context", version.BuildContext())

	flag.Parse()

	// Create settings.
	settings = &config.Settings{}

	// Create sessions pool.
	sessionsPool = &sessions.Sessions{}

	// Reset sessions pool after every settings reload.
	settings.AfterLoad = func(config config.Config) error {
		return sessionsPool.Load(config.Instances)
	}

	// Read configuration from file.
	err := settings.Load(*configFile)
	if err != nil {
		log.Fatalf("Can't read configuration file: %s\n", err.Error())
	}

	// Basic Metrics
	{
		// Create new Exporter with provided settings and session pool.
		exporter := basic.New(settings, sessionsPool)
		registry := prometheus.NewRegistry()
		registry.MustRegister(exporter)
		handler := promhttp.HandlerFor(registry, promhttp.HandlerOpts{})
		// Expose the exporter's own metrics on given path
		http.Handle(*basicMetricsPath, handler)
	}

	// Enhanced Metrics
	{
		// Create new Exporter with provided settings and session pool.
		exporter := enhanced.New(settings, sessionsPool)
		registry := prometheus.NewRegistry()
		registry.MustRegister(exporter)
		handler := promhttp.HandlerFor(registry, promhttp.HandlerOpts{})
		// Expose the exporter's own metrics on given path
		http.Handle(*enhancedMetricsPath, handler)
	}

	// Allows manual reload of the configuration
	http.HandleFunc("/reload", handleReload)

	// Inform user we are ready.
	log.Infoln("RDS exporter started")

	// Start serving for clients
	log.Fatal(http.ListenAndServe(*listenAddress, nil))
}
