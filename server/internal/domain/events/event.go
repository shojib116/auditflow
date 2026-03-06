package events

import (
	"time"
)

type DomainEvent interface {
	EventName() string
	OccurredAt() time.Time
}

// HandlerFunc is the signature all handlers must implement
type HandlerFunc func(event DomainEvent) error

// Bus is the interface your app depends on — lives in domain, no imports from infra
type Bus interface {
	Publish(event DomainEvent) error
	Subscribe(eventName string, handler HandlerFunc)
}
