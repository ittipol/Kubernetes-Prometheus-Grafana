package main

import "github.com/prometheus/client_golang/prometheus"

// metrics represents Prometheus metrics.
type metrics struct {
	devices prometheus.Gauge

	info *prometheus.GaugeVec

	upgrades *prometheus.CounterVec

	duration *prometheus.HistogramVec
}

// Create new metrics and register them with the Prometheus registry.
func NewMetrics(reg prometheus.Registerer) *metrics {
	// Create Prometheus metrics.
	m := &metrics{
		devices: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: "myapp",
			Name:      "connected_devices",
			Help:      "Number of currently connected devices.",
		}),
		info: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: "myapp",
			Name:      "info",
			Help:      "Information about the My App environment.",
		}, []string{"version"}),
		upgrades: prometheus.NewCounterVec(prometheus.CounterOpts{
			Namespace: "myapp",
			Name:      "device_upgrade_total",
			Help:      "Number of upgraded devices.",
		}, []string{"type"}),
		duration: prometheus.NewHistogramVec(prometheus.HistogramOpts{
			Namespace: "myapp",
			Name:      "request_duration_seconds",
			Help:      "Duration of the request.",
			// 4 times larger for apdex score
			// Buckets: prometheus.ExponentialBuckets(0.1, 1.5, 5),
			// Buckets: prometheus.LinearBuckets(0.1, 5, 5),
			Buckets: []float64{0.1, 0.15, 0.2, 0.25, 0.3},
		}, []string{"status", "method"}),
	}
	// Register metrics with Prometheus registry.
	reg.MustRegister(m.duration)

	return m
}
