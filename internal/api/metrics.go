package api

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

type EventMetrics struct {
	Requests    prometheus.Counter
	ReqDuration prometheus.Histogram
	DbDuration  prometheus.Histogram
}

func CreateEventMetrics(reg prometheus.Registerer) *EventMetrics {
	m := &EventMetrics{
		Requests: promauto.With(reg).NewCounter(prometheus.CounterOpts{
			Name: "api_events_requests_total",
			Help: "HTTP requests handled.",
		}),
		ReqDuration: promauto.With(reg).NewHistogram(prometheus.HistogramOpts{
			Name:    "api_events_request_duration_seconds",
			Help:    "Total request duration.",
			Buckets: prometheus.ExponentialBuckets(0.001, 2, 14),
		}),
		DbDuration: promauto.With(reg).NewHistogram(prometheus.HistogramOpts{
			Name:    "api_events_request_db_duration_seconds",
			Help:    "Duration of the INSERT call (includes waiting for a pool connection).",
			Buckets: prometheus.ExponentialBuckets(0.001, 2, 14),
		}),
	}

	return m
}
