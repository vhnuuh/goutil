// https://studygolang.com/articles/32055?fr=sidebar
package main

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"math"
	"math/rand"
	"net/http"
	"time"
)

var (
	TestCounter = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "test_couter",
		Help: "test_counter",
	})
	TestGauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "test_gauge",
		Help: "test_gauge",
	})
	TestHistogram = prometheus.NewHistogram(prometheus.HistogramOpts{
		Name:    "test_histogram",
		Help:    "test_histogram",
		Buckets: prometheus.LinearBuckets(20, 5, 5),
	})
	TestSummary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name:       "test_summary",
		Help:       "test_summary",
		Objectives: map[float64]float64{0.5: 0.5, 0.9: 0.01, 0.99: 0.1},
	})
)

func main() {
	prometheus.MustRegister(TestGauge)
	prometheus.MustRegister(TestHistogram)
	prometheus.MustRegister(TestSummary)
	prometheus.MustRegister(TestCounter)

	go func() {
		i := 0.0
		for {
			TestGauge.Add(1)
			TestCounter.Add(1)
			TestHistogram.Observe(30 + math.Floor(float64(rand.Intn(120))*math.Sin(i*0.1))/10)
			TestSummary.Observe(30 + math.Floor(float64(rand.Intn(120))*math.Sin(i*0.1))/10)
			time.Sleep(2 * time.Second)
			i += 1
		}
	}()
	http.Handle("/metrics", promhttp.Handler())
	err := http.ListenAndServe("localhost:2112", nil)
	if err != nil {
		fmt.Println(err)
	}
}
