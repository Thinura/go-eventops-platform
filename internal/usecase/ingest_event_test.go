package usecase

import (
	"context"
	"testing"
	"time"

	"github.com/Thinura/go-eventops-platform/internal/apperror"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestIngestEventUseCase_Execute(t *testing.T) {
	uc := NewIngestEventUseCase()

	input := IngestEventInput{
		Source:     "payment-service",
		EventType:  "payment.failed",
		EntityID:   "order-1001",
		Payload:    map[string]any{"amount": 4500},
		OccurredAt: time.Now().UTC(),
	}

	output, err := uc.Execute(context.Background(), input)
	require.NoError(t, err)
	require.NotNil(t, output)
	assert.NotEmpty(t, output.ID)
	assert.Equal(t, "accepted", output.Status)
}

func TestIngestEventUseCase_Execute_InvalidInput(t *testing.T) {
	uc := NewIngestEventUseCase()

	input := IngestEventInput{
		Source:     "",
		EventType:  "payment.failed",
		EntityID:   "order-1001",
		Payload:    map[string]any{},
		OccurredAt: time.Now().UTC(),
	}

	output, err := uc.Execute(context.Background(), input)
	require.Error(t, err)
	assert.Nil(t, output)

	var appErr *apperror.Error
	require.ErrorAs(t, err, &appErr)
	assert.Equal(t, apperror.CodeEventSourceRequired, appErr.Code)
}
