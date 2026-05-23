
package apperror

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCodeValues(t *testing.T) {
	tests := []struct {
		name     string
		code     Code
		expected string
	}{
		{name: "validation", code: CodeValidation, expected: "VALIDATION_ERROR"},
		{name: "not found", code: CodeNotFound, expected: "NOT_FOUND"},
		{name: "conflict", code: CodeConflict, expected: "CONFLICT"},
		{name: "internal", code: CodeInternal, expected: "INTERNAL_ERROR"},
		{name: "invalid request body", code: CodeInvalidRequestBody, expected: "INVALID_REQUEST_BODY"},
		{name: "event source required", code: CodeEventSourceRequired, expected: "EVENT_SOURCE_REQUIRED"},
		{name: "event type required", code: CodeEventTypeRequired, expected: "EVENT_TYPE_REQUIRED"},
		{name: "event entity id required", code: CodeEventEntityIDRequired, expected: "EVENT_ENTITY_ID_REQUIRED"},
		{name: "event occurred at missing", code: CodeEventOccurredAtMissing, expected: "EVENT_OCCURRED_AT_REQUIRED"},
		{name: "unsupported event type", code: CodeUnsupportedEventType, expected: "UNSUPPORTED_EVENT_TYPE"},
		{name: "event publish failed", code: CodeEventPublishFailed, expected: "EVENT_PUBLISH_FAILED"},
		{name: "event save failed", code: CodeEventSaveFailed, expected: "EVENT_SAVE_FAILED"},
		{name: "event already exists", code: CodeEventAlreadyExists, expected: "EVENT_ALREADY_EXISTS"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, string(tt.code))
		})
	}
}
