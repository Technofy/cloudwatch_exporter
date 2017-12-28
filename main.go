package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/percona/rds_exporter/config"
	"github.com/percona/rds_exporter/sessions"
)

var (
	listenAddress = flag.String("web.listen-address", ":9042", "Address on which to expose metrics and web interface.")
	metricsPath   = flag.String("web.telemetry-path", "/metrics", "Path under which to expose exporter's metrics.")
	configFile    = flag.String("config.file", "config.yml", "Path to configuration file.")

	settings     *config.Settings
	sessionsPool *sessions.Sessions
)

// handleReload handles a full reload of the configuration file and regenerates the collector templates.
func handleReload(w http.ResponseWriter, req *http.Request) {
	err := settings.Load(*configFile)
	if err != nil {
		str := fmt.Sprintf("Can't read configuration file: %s", err.Error())
		fmt.Fprintln(w, str)
		fmt.Println(str)
	}
	fmt.Fprintln(w, "Reload complete")
}

func main() {
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
		fmt.Printf("Can't read configuration file: %s\n", err.Error())
		os.Exit(-1)
	}

	// Create new Exporter with provided settings and session pool.
	exporter := New(settings, sessionsPool)
	prometheus.MustRegister(exporter)

	// Expose the exporter's own metrics on /metrics
	http.Handle(*metricsPath, promhttp.Handler())

	// Allows manual reload of the configuration
	http.HandleFunc("/reload", handleReload)

	// Inform user we are ready.
	fmt.Println("RDS exporter started...")

	// Start serving for clients
	log.Fatal(http.ListenAndServe(*listenAddress, nil))
}
