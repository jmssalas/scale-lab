package api

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

type EventMetrics struct {
	Requests        prometheus.Counter
	ReqDuration     prometheus.Histogram
	ProcessDuration prometheus.Histogram
}

func CreateEventMetrics(reg prometheus.Registerer) *EventMetrics {
	m := &EventMetrics{
		Requests: promauto.With(reg).NewCounter(prometheus.CounterOpts{
			Name: "api_events_requests_total",
			Help: "HTTP requests handled.",
		}),
		ReqDuration: promauto.With(reg).NewHistogram(prometheus.HistogramOpts{
			Name:    "api_events_request_duration_seconds",
			Help:    "Total event request duration in seconds.",
			Buckets: prometheus.DefBuckets,
		}),
		ProcessDuration: promauto.With(reg).NewHistogram(prometheus.HistogramOpts{
			Name:    "api_events_processing_duration_seconds",
			Help:    "Duration of event processing in seconds.",
			Buckets: prometheus.DefBuckets,
		}),
	}

	return m
}
