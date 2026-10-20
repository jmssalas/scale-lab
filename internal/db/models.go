package db

import (
	"time"

	"github.com/jmssalas/scale-lab/internal/event"
	"gorm.io/gorm"
)

type Event struct {
	gorm.Model
	event.Event
	ReceivedAt time.Time
}
