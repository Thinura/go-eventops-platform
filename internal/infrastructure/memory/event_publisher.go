package memory

import (
	"context"
	"sync"

	"github.com/Thinura/go-eventops-platform/internal/domain"
)

type EventPublisher struct {
	mu     sync.Mutex
	events []domain.Event
}

func NewEventPublisher() *EventPublisher {
	return &EventPublisher{
		events: make([]domain.Event, 0),
	}
}

func (p *EventPublisher) Publish(ctx context.Context, event domain.Event) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}

	p.mu.Lock()
	defer p.mu.Unlock()

	p.events = append(p.events, event)

	return nil
}

func (p *EventPublisher) Events() []domain.Event {
	p.mu.Lock()
	defer p.mu.Unlock()

	copied := make([]domain.Event, len(p.events))
	copy(copied, p.events)

	return copied
}
