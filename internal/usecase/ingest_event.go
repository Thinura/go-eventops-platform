package usecase

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/Thinura/go-eventops-platform/internal/apperror"
	"github.com/Thinura/go-eventops-platform/internal/domain"
	"github.com/Thinura/go-eventops-platform/internal/port"
)

type IngestEventInput struct {
	Source     string
	EventType  string
	EntityID   string
	Payload    map[string]any
	OccurredAt time.Time
}

type IngestEventOutput struct {
	ID     string
	Status string
}

type IngestEventUseCase struct {
	publisher port.EventPublisher
}

func NewIngestEventUseCase(publisher port.EventPublisher) *IngestEventUseCase {
	return &IngestEventUseCase{
		publisher: publisher,
	}
}

func (uc *IngestEventUseCase) Execute(ctx context.Context, input IngestEventInput) (*IngestEventOutput, error) {
	event := domain.Event{
		ID:             uuid.NewString(),
		Source:         input.Source,
		EventType:      domain.EventType(input.EventType),
		EntityID:       input.EntityID,
		Payload:        input.Payload,
		OccurredAt:     input.OccurredAt,
		ReceivedAt:     time.Now().UTC(),
		IdempotencyKey: buildIdempotencyKey(input.Source, input.EventType, input.EntityID, input.OccurredAt),
	}

	if err := event.Validate(); err != nil {
		return nil, err
	}

	if err := uc.publisher.Publish(ctx, event); err != nil {
		return nil, apperror.Wrap(
			apperror.CodeEventPublishFailed,
			"failed to publish event",
			err,
		)
	}

	return &IngestEventOutput{
		ID:     event.ID,
		Status: "accepted",
	}, nil
}

func buildIdempotencyKey(source, eventType, entityID string, occurredAt time.Time) string {
	raw := fmt.Sprintf("%s:%s:%s:%s", source, eventType, entityID, occurredAt.UTC().Format(time.RFC3339Nano))

	hash := sha256.Sum256([]byte(raw))

	return hex.EncodeToString(hash[:])
}