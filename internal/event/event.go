// Package event defines the event model shared by all services.
package event

import (
	"encoding/json"
	"time"
)

type Event struct {
	SourceID  string          `json:"sourceId"`
	Type      string          `json:"type"`
	Timestamp time.Time       `json:"ts"`
	Payload   json.RawMessage `json:"payload,omitempty"`
}
