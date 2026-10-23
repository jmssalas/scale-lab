package loadgen

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

type GeneratorMetrics struct {
	requests   *prometheus.CounterVec
	latency    prometheus.Histogram
	targetRate prometheus.Gauge
	dropped    prometheus.Counter
}

func CreateGeneratorMetrics(reg prometheus.Registerer) *GeneratorMetrics {
	m := &GeneratorMetrics{
		requests: promauto.With(reg).NewCounterVec(prometheus.CounterOpts{
			Name: "loadgen_requests_total",
			Help: "Requests sent, by HTTP status (or 'error').",
		}, []string{"status"}),
		latency: promauto.With(reg).NewHistogram(prometheus.HistogramOpts{
			Name:    "loadgen_latency_seconds",
			Help:    "Client-side latency of each requests.",
			Buckets: prometheus.DefBuckets,
		}),
		targetRate: promauto.With(reg).NewGauge(prometheus.GaugeOpts{
			Name: "loadgen_target_requests_per_second",
			Help: "Configured target request rate.",
		}),
		dropped: promauto.With(reg).NewCounter(prometheus.CounterOpts{
			Name: "loadgen_dropped_requests_total",
			Help: "Total scheduled requests dropped because the work queue was full.",
		}),
	}

	return m
}
