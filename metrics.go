package main

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

var (
	age12hours = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "kube_pod_start_time",
		Help: "Should not alert",
		ConstLabels: prometheus.Labels{
			"region":      "eastus",
			"underlay":    "c2",
			"shouldAlert": "false",
		},
	})
)

func init() {
	twelveHoursAgoEpoch := time.Now().Add(-12 * time.Hour).Unix()

	age12hours.Set(float64(twelveHoursAgoEpoch))

	prometheus.MustRegister(age12hours)
}
