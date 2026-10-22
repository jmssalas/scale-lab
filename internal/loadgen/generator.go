package loadgen

import (
	"bytes"
	"context"
	"encoding/json"
	"log/slog"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/jmssalas/scale-lab/internal/event"
)

const (
	DEFAULT_TARGET_URL = "http://localhost:8080/api/v1/events"
	DEFAULT_RATE       = 5000
	DEFAULT_WORKERS    = 10
	DEFAULT_QUEUE_SIZE = 100
)

type config struct {
	targetUrl         string
	requestsPerSecond int
	workers           int
	queueSize         int
}

type Generator struct {
	config  *config
	client  *http.Client
	metrics *GeneratorMetrics
}

func CreateGenerator(m *GeneratorMetrics) *Generator {
	return &Generator{
		config:  loadConfig(),
		client:  &http.Client{},
		metrics: m,
	}
}

func (g *Generator) Run(ctx context.Context) {
	jobs := make(chan struct{}, g.config.queueSize)
	var workers sync.WaitGroup

	for workerId := 0; workerId < g.config.workers; workerId++ {
		workers.Add(1)
		go func(id int) {
			defer workers.Done()
			g.runWorker(ctx, jobs, id)
		}(workerId)
	}

	ticker := time.NewTicker(time.Second / time.Duration(g.config.requestsPerSecond))
	defer ticker.Stop()

	slog.Info("Load generator initialized", "config", g.config)

	for {
		select {
		case <-ctx.Done():
			close(jobs)
			workers.Wait()
			return
		case <-ticker.C:
			select {
			case jobs <- struct{}{}:
			default:
				// @TODO: Increase dropped counter
			}
		}
	}

}

func (g *Generator) runWorker(ctx context.Context, jobs chan struct{}, id int) {
	for range jobs {
		g.sendRequest(ctx, id)
	}
}

func (g *Generator) sendRequest(ctx context.Context, workerId int) {
	payload, err := json.Marshal(event.Event{
		SourceID:  "loadgen-device-" + strconv.Itoa(workerId),
		Type:      "telemetry",
		Timestamp: time.Now(),
		Payload:   json.RawMessage(`{"simulated":true}`),
	})
	if err != nil {
		// @TODO: Handle error
		slog.Error(err.Error())
		return
	}

	request, err := http.NewRequestWithContext(ctx, http.MethodPost, g.config.targetUrl, bytes.NewReader(payload))
	if err != nil {
		// @TODO: Handle error
		slog.Error(err.Error())
		return
	}
	request.Header.Set("Content-Type", "application/json")

	response, err := g.client.Do(request)
	if err != nil {
		// @TODO: Handle error
		slog.Error(err.Error())
		return
	}
	defer response.Body.Close()

	if response.StatusCode < http.StatusOK || response.StatusCode >= http.StatusMultipleChoices {
		// @TODO: Handle request status
	}
}

func loadConfig() *config {
	return &config{
		targetUrl:         envStr("LOADGEN_TARGET_URL", DEFAULT_TARGET_URL),
		requestsPerSecond: envInt("LOADGEN_RATE", DEFAULT_RATE),
		workers:           envInt("LOADGEN_WORKERS", DEFAULT_WORKERS),
		queueSize:         envInt("LOADGEN_QUEUE_SIZE", DEFAULT_QUEUE_SIZE),
	}
}

func envInt(k string, def int) int {
	if v := os.Getenv(k); v != "" {
		if n, err := strconv.Atoi(v); err == nil {
			return n
		}
	}
	return def
}

func envStr(k, def string) string {
	if v := os.Getenv(k); v != "" {
		return v
	}
	return def
}
