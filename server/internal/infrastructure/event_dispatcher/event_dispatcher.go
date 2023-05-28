package eventdispatcher

import (
	"context"
	"sync"

	"github.com/glopezep/arithmetic-calculator/internal/common"
)

type EventSubscriber interface {
	Subscribe(event common.Event, handler common.EventHandler)
}

type EventPublisher interface {
	Publish(ctx context.Context, events ...common.Event) error
}

type EventDispatcher struct {
	handlers map[string][]common.EventHandler
	mu       sync.Mutex
}

var _ interface {
	EventSubscriber
	EventPublisher
} = (*EventDispatcher)(nil)

func NewEventDispatcher() *EventDispatcher {
	return &EventDispatcher{
		handlers: make(map[string][]common.EventHandler),
	}
}

func (h *EventDispatcher) Subscribe(event common.Event, handler common.EventHandler) {
	h.mu.Lock()
	defer h.mu.Unlock()

	h.handlers[event.String()] = append(h.handlers[event.String()], handler)
}

func (h *EventDispatcher) Publish(ctx context.Context, events ...common.Event) error {
	for _, event := range events {
		for _, handler := range h.handlers[event.String()] {
			err := handler(ctx, event)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
