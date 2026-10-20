package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	ep "github.com/jmssalas/scale-lab/internal/api/event_processors"
	"github.com/jmssalas/scale-lab/internal/event"
	"gorm.io/gorm"
)

type EventHandler struct {
	DB *gorm.DB
}

func (h *EventHandler) ProcessEvent(c *gin.Context) {
	var event event.Event

	if err := c.ShouldBindJSON(&event); err != nil {
		Fail(c, http.StatusBadRequest, "BAD_REQUEST", err.Error())
		return
	}

	// @TODO: Select Processor by Event Type
	var p ep.EventProcessor = &ep.DummyProcessor{DB: h.DB}

	if err := p.Process(&event); err != nil {
		Fail(c, http.StatusInternalServerError, "INTERNAL_SERVER_ERROR", err.Error())
		return
	}

	OK(c, nil)
}
