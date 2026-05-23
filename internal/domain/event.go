package domain

import (
	"strings"
	"time"

	"github.com/Thinura/go-eventops-platform/internal/apperror"
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

func (e Event) Validate() error {
	if strings.TrimSpace(e.Source) == "" {
		return apperror.Validation(apperror.CodeEventSourceRequired, "source is required")
	}

	if strings.TrimSpace(string(e.EventType)) == "" {
		return apperror.Validation(apperror.CodeEventTypeRequired, "event_type is required")
	}

	if !IsSupportedEventType(e.EventType) {
		return apperror.Validation(apperror.CodeUnsupportedEventType, "unsupported event_type")
	}

	if strings.TrimSpace(e.EntityID) == "" {
		return apperror.Validation(apperror.CodeEventEntityIDRequired, "entity_id is required")
	}

	if e.OccurredAt.IsZero() {
		return apperror.Validation(apperror.CodeEventOccurredAtMissing, "occurred_at is required")
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
