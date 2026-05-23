package domain

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

type EventType string

const (
	EventPaymentFailed       EventType = "payment.failed"
	EventOrderPlaced         EventType = "order.placed"
	EventDeviceOffline       EventType = "device.offline"
	EventJobFailed           EventType = "job.failed"
	EventJobCompleted        EventType = "job.completed"
	EventSystemError         EventType = "system.error"
	EventUserCreated         EventType = "user.created"
	EventDeploymentCompleted EventType = "deployment.completed"
)

type Event struct {
	ID             string
	Source         string
	EventType      EventType
	EntityID       string
	Payload        map[string]any
	OccurredAt     time.Time
	ReceivedAt     time.Time
	IdempotencyKey string
}

var (
	ErrEventSourceRequired    = errors.New("event source is required")
	ErrEventTypeRequired      = errors.New("event type is required")
	ErrEventEntityIDRequired  = errors.New("event entity_id is required")
	ErrEventOccurredAtMissing = errors.New("event occurred_at is required")
	ErrUnsupportedEventType   = errors.New("unsupported event type")
)

func (e Event) Validate() error {
	if strings.TrimSpace(e.Source) == "" {
		return ErrEventSourceRequired
	}

	if strings.TrimSpace(string(e.EventType)) == "" {
		return ErrEventTypeRequired
	}

	if !IsSupportedEventType(e.EventType) {
		return fmt.Errorf("%w: %s", ErrUnsupportedEventType, e.EventType)
	}

	if strings.TrimSpace(e.EntityID) == "" {
		return ErrEventEntityIDRequired
	}

	if e.OccurredAt.IsZero() {
		return ErrEventOccurredAtMissing
	}

	return nil
}

func IsSupportedEventType(eventType EventType) bool {
	switch eventType {
	case EventPaymentFailed,
		EventOrderPlaced,
		EventDeviceOffline,
		EventJobFailed,
		EventJobCompleted,
		EventSystemError,
		EventUserCreated,
		EventDeploymentCompleted:
		return true
	default:
		return false
	}
}
