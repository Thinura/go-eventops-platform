package memory

import (
	"context"
	"testing"
	"time"

	"github.com/Thinura/go-eventops-platform/internal/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewEventPublisher(t *testing.T) {
	publisher := NewEventPublisher()

	require.NotNil(t, publisher)
	assert.Empty(t, publisher.Events())
}

func TestEventPublisher_Publish(t *testing.T) {
	publisher := NewEventPublisher()

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

	events := publisher.Events()
	require.Len(t, events, 1)
	assert.Equal(t, event.ID, events[0].ID)
	assert.Equal(t, event.Source, events[0].Source)
	assert.Equal(t, event.EventType, events[0].EventType)
	assert.Equal(t, event.EntityID, events[0].EntityID)
	assert.Equal(t, event.Payload, events[0].Payload)
}

func TestEventPublisher_PublishMultipleEvents(t *testing.T) {
	publisher := NewEventPublisher()

	firstEvent := domain.Event{
		ID:         "event-1001",
		Source:     "payment-service",
		EventType:  domain.EventPaymentFailed,
		EntityID:   "order-1001",
		OccurredAt: time.Now().UTC(),
		ReceivedAt: time.Now().UTC(),
	}

	secondEvent := domain.Event{
		ID:         "event-1002",
		Source:     "job-service",
		EventType:  domain.EventJobFailed,
		EntityID:   "job-1001",
		OccurredAt: time.Now().UTC(),
		ReceivedAt: time.Now().UTC(),
	}

	require.NoError(t, publisher.Publish(context.Background(), firstEvent))
	require.NoError(t, publisher.Publish(context.Background(), secondEvent))

	events := publisher.Events()
	require.Len(t, events, 2)
	assert.Equal(t, firstEvent.ID, events[0].ID)
	assert.Equal(t, secondEvent.ID, events[1].ID)
}

func TestEventPublisher_PublishReturnsContextError(t *testing.T) {
	publisher := NewEventPublisher()

	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	err := publisher.Publish(ctx, domain.Event{})

	require.ErrorIs(t, err, context.Canceled)
	assert.Empty(t, publisher.Events())
}

func TestEventPublisher_EventsReturnsCopy(t *testing.T) {
	publisher := NewEventPublisher()

	event := domain.Event{
		ID:         "event-1001",
		Source:     "payment-service",
		EventType:  domain.EventPaymentFailed,
		EntityID:   "order-1001",
		OccurredAt: time.Now().UTC(),
		ReceivedAt: time.Now().UTC(),
	}

	require.NoError(t, publisher.Publish(context.Background(), event))

	events := publisher.Events()
	require.Len(t, events, 1)

	events[0].ID = "mutated-event-id"

	storedEvents := publisher.Events()
	require.Len(t, storedEvents, 1)
	assert.Equal(t, event.ID, storedEvents[0].ID)
}
