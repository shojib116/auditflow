package memory

import (
	"log/slog"
	"sync"

	"github.com/shojib116/auditflow-api/internal/domain/events"
)

type Bus struct {
	mu       sync.RWMutex
	handlers map[string][]events.HandlerFunc
}

func NewBus() *Bus {
	return &Bus{handlers: make(map[string][]events.HandlerFunc)}
}

func (b *Bus) Subscribe(eventName string, handler events.HandlerFunc) {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.handlers[eventName] = append(b.handlers[eventName], handler)
}

func (b *Bus) Publish(event events.DomainEvent) error {
	b.mu.RLock()
	handlers := b.handlers[event.EventName()]
	b.mu.RUnlock()

	for _, h := range handlers {
		if err := h(event); err != nil {
			slog.Error("event handler failed",
				"event", event.EventName(), "err", err)
		}
	}
	return nil
}
