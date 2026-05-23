package domain

import (
	"testing"
	"time"

	"github.com/Thinura/go-eventops-platform/internal/apperror"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestEvent_Validate(t *testing.T) {
	tests := []struct {
		name        string
		event       Event
		expectedErr apperror.Code
	}{
		{
			name: "valid event",
			event: Event{
				Source:     "payment-service",
				EventType:  EventPaymentFailed,
				EntityID:   "order-1001",
				Payload:    map[string]any{"amount": 4500},
				OccurredAt: time.Now().UTC(),
			},
			expectedErr: "",
		},
		{
			name: "missing source",
			event: Event{
				EventType:  EventPaymentFailed,
				EntityID:   "order-1001",
				OccurredAt: time.Now().UTC(),
			},
			expectedErr: apperror.CodeEventSourceRequired,
		},
		{
			name: "missing event type",
			event: Event{
				Source:     "payment-service",
				EntityID:   "order-1001",
				OccurredAt: time.Now().UTC(),
			},
			expectedErr: apperror.CodeEventTypeRequired,
		},
		{
			name: "unsupported event type",
			event: Event{
				Source:     "payment-service",
				EventType:  EventType("wrong.type"),
				EntityID:   "order-1001",
				OccurredAt: time.Now().UTC(),
			},
			expectedErr: apperror.CodeUnsupportedEventType,
		},
		{
			name: "missing entity id",
			event: Event{
				Source:     "payment-service",
				EventType:  EventPaymentFailed,
				OccurredAt: time.Now().UTC(),
			},
			expectedErr: apperror.CodeEventEntityIDRequired,
		},
		{
			name: "missing occurred at",
			event: Event{
				Source:    "payment-service",
				EventType: EventPaymentFailed,
				EntityID:  "order-1001",
			},
			expectedErr: apperror.CodeEventOccurredAtMissing,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.event.Validate()

			if tt.expectedErr == "" {
				require.NoError(t, err)
				return
			}

			require.Error(t, err)

			var appErr *apperror.Error
			require.ErrorAs(t, err, &appErr)
			assert.Equal(t, tt.expectedErr, appErr.Code)
		})
	}
}
