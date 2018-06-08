package main

import "github.com/prometheus/client_golang/prometheus"

var (
	normal = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "test_normal",
		Help: "Should go to all writers",
	})
	filter = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "test_filter",
		Help: "Only go to filter endpoint",
	})
)

func init() {
	normal.Set(123)
	filter.Set(456)

	prometheus.MustRegister(normal)
	prometheus.MustRegister(filter)
}
