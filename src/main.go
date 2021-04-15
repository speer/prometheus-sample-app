package main

import (
  "net/http"
  "fmt"

  "github.com/prometheus/client_golang/prometheus"
  "github.com/prometheus/client_golang/prometheus/promauto"
  "github.com/prometheus/client_golang/prometheus/promhttp"
)

func handler(w http.ResponseWriter, r *http.Request) {
  requestCounter.WithLabelValues(r.URL.Path).Inc()
  fmt.Fprintf(w, "Hello %s!", r.URL.Path[1:])
}

var (
  requestCounter = promauto.NewCounterVec(prometheus.CounterOpts{
    Name: "sample_app_request_count",
    Help: "The total number of requests",
  }, []string{"path"})
)

func main() {
  http.Handle("/metrics", promhttp.Handler())
  http.HandleFunc("/", handler)
  http.ListenAndServe(":8080", nil)
}
