package domain

import "time"

type Event struct {
	ID             string
	Source         string
	EventType      string
	EntityID       string
	Payload        map[string]any
	OccurredAt     time.Time
	ReceivedAt     time.Time
	IdempotencyKey string
}
