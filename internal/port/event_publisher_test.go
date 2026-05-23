
package port

import (
	"context"
	"testing"
	"time"

	"github.com/Thinura/go-eventops-platform/internal/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type fakeEventPublisher struct {
	publishedEvents []domain.Event
}

func (p *fakeEventPublisher) Publish(ctx context.Context, event domain.Event) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}

	p.publishedEvents = append(p.publishedEvents, event)
	return nil
}

func TestEventPublisherContract(t *testing.T) {
	var publisher EventPublisher = &fakeEventPublisher{}

	event := domain.Event{
		ID:         "event-1001",
		Source:     "payment-service",
		EventType:  domain.EventPaymentFailed,
		EntityID:   "order-1001",
		Payload:    map[string]any{"amount": 4500},
		OccurredAt: time.Now().UTC(),
		ReceivedAt: time.Now().UTC(),
	}

	err := publisher.Publish(context.Background(), event)

	require.NoError(t, err)

	fakePublisher, ok := publisher.(*fakeEventPublisher)
	require.True(t, ok)
	require.Len(t, fakePublisher.publishedEvents, 1)
	assert.Equal(t, event.ID, fakePublisher.publishedEvents[0].ID)
	assert.Equal(t, event.Source, fakePublisher.publishedEvents[0].Source)
	assert.Equal(t, event.EventType, fakePublisher.publishedEvents[0].EventType)
	assert.Equal(t, event.EntityID, fakePublisher.publishedEvents[0].EntityID)
}

func TestEventPublisherContract_ReturnsContextError(t *testing.T) {
	var publisher EventPublisher = &fakeEventPublisher{}

	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	err := publisher.Publish(ctx, domain.Event{})

	require.ErrorIs(t, err, context.Canceled)
}
