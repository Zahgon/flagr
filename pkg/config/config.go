package config

import (
	"github.com/DataDog/datadog-go/statsd"
	"github.com/caarlos0/env"
	newrelic "github.com/newrelic/go-agent"
	"github.com/prometheus/client_golang/prometheus"
)

// EvalOnlyModeDBDrivers is a list of DBDrivers that we should only run in EvalOnlyMode.
var EvalOnlyModeDBDrivers = map[string]struct{}{
	"json_file": {},
	"json_http": {},
}

// Global is the global dependency we can use, such as the new relic app instance
var Global = struct {
	NewrelicApp  newrelic.Application
	StatsdClient *statsd.Client
	Prometheus   prometheusMetrics
}{}

func init() {
	env.Parse(&Config)

	setupEvalOnlyMode()
	setupSentry()
	setupLogrus()
	setupStatsd()
	setupNewrelic()
	setupPrometheus()
}

func setupEvalOnlyMode() { _ = "STUB: not implemented"; return }

func setupLogrus() { _ = "STUB: not implemented"; return }

func setupSentry() { _ = "STUB: not implemented"; return }

func setupStatsd() { _ = "STUB: not implemented"; return }

func setupNewrelic() { _ = "STUB: not implemented"; return }

// These two cannot be enabled at the same time and cross application is enabled by default

type prometheusMetrics struct {
	ScrapePath       string
	EvalCounter      *prometheus.CounterVec
	RequestCounter   *prometheus.CounterVec
	RequestHistogram *prometheus.HistogramVec
}

func setupPrometheus() { _ = "STUB: not implemented"; return }
