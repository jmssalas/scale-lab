package eventprocessors

import "github.com/jmssalas/scale-lab/internal/event"

type EventProcessor interface {
	Process(e *event.Event) error
}
