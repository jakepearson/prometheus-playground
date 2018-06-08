package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/prometheus/prompb"
)

func getLabels(metrics *prompb.WriteRequest) []string {
	labels := make([]string, 0)
	for _, timeSeries := range metrics.Timeseries {
		for _, label := range timeSeries.Labels {
			if label.Name == "__name__" {
				labels = append(labels, label.Value)
			}
		}
	}
	return labels
}

func createHandler(endpoint string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		now := time.Now().Format("04:04:05")
		metrics := parseMetrics(w, r)

		fmt.Printf("--- %s %s: %d\n", now, endpoint, len(metrics.Timeseries))
		for _, label := range getLabels(metrics) {
			fmt.Printf("%s %s: %s\n", now, endpoint, label)
		}
	}
}

func main() {
	http.Handle("/normal", http.HandlerFunc(createHandler("normal")))
	http.Handle("/filter", http.HandlerFunc(createHandler("filter")))
	http.Handle("/metrics", promhttp.Handler())
	http.HandleFunc("/test-data/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, r.URL.Path[1:])
	})
	err := http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
