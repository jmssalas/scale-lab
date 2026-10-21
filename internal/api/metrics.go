package api

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

type Metrics struct {
	EventsRequests prometheus.Counter
}

func CreateMetrics(reg prometheus.Registerer) *Metrics {
	m := &Metrics{
		EventsRequests: promauto.With(reg).NewCounter(prometheus.CounterOpts{
			Name: "api_events_requests_total",
			Help: "HTTP requests handled",
		}),
	}

	return m
}
