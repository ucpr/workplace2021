package main

import (
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	httpReqs = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "How many HTTP requests processed, partitioned by status code and HTTP method.",
		},
		[]string{"method"},
	)
)

func init() {
	prometheus.MustRegister(httpReqs)
}

func main() {
	http.Handle("/", http.HandlerFunc(count))
	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func count(w http.ResponseWriter, r *http.Request) {
	a := httpReqs.WithLabelValues("1")
	b := httpReqs.WithLabelValues("2")
	a.Inc()
	b.Inc()
}
