package eventprocessors

import (
	"context"
	"time"

	"github.com/jmssalas/scale-lab/internal/db"
	"github.com/jmssalas/scale-lab/internal/event"
	"gorm.io/gorm"
)

type DummyProcessor struct {
	DB *gorm.DB
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

	return gorm.G[db.Event](p.DB).Create(ctx, &event)
}
