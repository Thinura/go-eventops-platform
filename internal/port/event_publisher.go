package port

import (
	"context"

	"github.com/Thinura/go-eventops-platform/internal/domain"
)

type EventPublisher interface {
	Publish(ctx context.Context, event domain.Event) error
}
