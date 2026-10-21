package api

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jmssalas/scale-lab/internal/db"
	"github.com/jmssalas/scale-lab/internal/event"
	"gorm.io/gorm"
)

// TODO: Everything regarding EventProcessor should be refactored
// when multiple processors will be added. Til now, it will be here for simplicity

type EventProcessor interface {
	Process(e *event.Event) error
}

type DummyProcessor struct {
	db *gorm.DB
}

func (p *DummyProcessor) Process(e *event.Event) error {
	ctx := context.Background()
	event := db.Event{
		SourceID:   e.SourceID,
		Type:       e.Type,
		Timestamp:  e.Timestamp,
		Payload:    e.Payload,
		ReceivedAt: time.Now(),
	}

	return gorm.G[db.Event](p.db).Create(ctx, &event)
}

type EventHandler struct {
	db *gorm.DB
}

func (h *EventHandler) ProcessEvent(c *gin.Context) {
	var event event.Event

	if err := c.ShouldBindJSON(&event); err != nil {
		Fail(c, http.StatusBadRequest, "BAD_REQUEST", err.Error())
		return
	}

	var p EventProcessor = &DummyProcessor{db: h.db}

	if err := p.Process(&event); err != nil {
		Fail(c, http.StatusInternalServerError, "INTERNAL_SERVER_ERROR", err.Error())
		return
	}

	OK(c, nil)
}
