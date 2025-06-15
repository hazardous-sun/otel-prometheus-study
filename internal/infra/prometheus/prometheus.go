package prometheus

import (
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// StartPrometheusEndpoint
// Starts an HTTP server at /metrics
func StartPrometheusEndpoint(port string) {
	go func() {
		http.Handle("/metrics", promhttp.Handler())
		log.Printf("Serving Prometheus metrics at http://localhost:%s/metrics\n", port)
		if err := http.ListenAndServe(":"+port, nil); err != nil {
			log.Fatalf("failed to start Prometheus HTTP server: %v", err)
		}
	}()
}
