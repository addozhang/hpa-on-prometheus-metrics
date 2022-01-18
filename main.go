package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
	"strconv"
)

func main() {
	metrics := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name:        "http_requests_total",
			Help:        "Number of total http requests",
		},
		[]string{"status"},
	)
	prometheus.MustRegister(metrics)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		statusCode := 200
		switch path {
		case "/metrics":
			promhttp.Handler().ServeHTTP(w, r)
		default:
			w.WriteHeader(statusCode)
			w.Write([]byte("Hello World!"))
		}
		metrics.WithLabelValues(strconv.Itoa(statusCode)).Inc()
	})
	http.ListenAndServe(":3000", nil)
}

