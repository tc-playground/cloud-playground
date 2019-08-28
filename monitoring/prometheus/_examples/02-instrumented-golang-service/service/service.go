package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	// Prometheus
	//
	// Histogram to collect required metrics
	pongHistogram := prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "pong_seconds",
			Help:    "Time taken to pong a ping",
			Buckets: []float64{1, 2, 5, 6, 10}, 		// Define Bucket ranges.
		}, 
		[]string{"code"}, 								// Partition the results by code.
	)
	prometheus.Register(pongHistogram)					// Register the histogram metric with the default 'registry'.

	// Service
	//
	// Define a simple `gorilla` router with two basic services.
	//
	router := mux.NewRouter()
	router.Handle("/ping", Ping())
	router.Handle("/echo/{toEcho}", Echo())
	router.Handle("/metrics", promhttp.Handler())			// Add a metrics endpoint for scraping.

	fmt.Println("Starting service: http://localhost:8009/ping")
	fmt.Println("Starting service: http://localhost:8009/echo/:toEcho")
	log.Fatal(http.ListenAndServe(":8009", router))
}

// Ping returns 'Pong'!
func Ping() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		code := http.StatusBadRequest
		if r.Method == "GET" {
			code = http.StatusOK
			w.Write([]byte("Pong"))
		} else {
			w.WriteHeader(code)
		}
	}
}

// Echo return an echo!
func Echo() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		code := http.StatusBadRequest
		if r.Method == "GET" {
			code = http.StatusOK
			vars := mux.Vars(r)
			toEcho := vars["toEcho"]
			echoed := fmt.Sprintf("%s", toEcho)
			w.Write([]byte(echoed))
		} else {
			w.WriteHeader(code)
		}
	}
}
