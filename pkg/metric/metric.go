package metric

import (
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type Metric struct {
	labelDefault        []string
	latencyAPIHistogram *prometheus.HistogramVec
	latencyAPIGauge     *prometheus.GaugeVec
}

func New() (M *Metric) {
	M = &Metric{}

	M.labelDefault = []string{
		"name",
		"domain",
		"method",
		"test_result",
		"http_code",
		"location",
		"geohash",
		"ISP",
	}

	M.latencyAPIHistogram = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "API_latency_histo",
			Help:    "latency of API being tested",
			Buckets: prometheus.ExponentialBuckets(0.2, 5, 5),
		}, M.labelDefault)

	M.latencyAPIGauge = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "API_latency_gauge",
			Help: "latency of API being tested",
		}, M.labelDefault)

	return
}

// Init metric runtime
// it adds http handler to default mux, run before http listen and serve
func (M *Metric) Init() {
	prometheus.MustRegister(M.latencyAPIHistogram)
	prometheus.MustRegister(M.latencyAPIGauge)

	// add metrics handler to DefaultServeMux
	http.Handle("/metrics", promhttp.Handler())
}

// MeasureAPILatency use this to measure API latency
// label are "domain", "method", "test_result", "http_code", "location", "geohash", "ISP"
func (M *Metric) MeasureAPILatency(latency time.Duration, labels ...string) {
	M.latencyAPIHistogram.WithLabelValues(labels...).Observe(latency.Seconds())
	M.latencyAPIGauge.WithLabelValues(labels...).Set(latency.Seconds())
}
