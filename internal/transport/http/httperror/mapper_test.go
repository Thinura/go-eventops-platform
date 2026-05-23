
package httperror

import (
	"errors"
	"net/http"
	"testing"

	"github.com/Thinura/go-eventops-platform/internal/apperror"
	"github.com/stretchr/testify/assert"
)

func TestMap_AppErrors(t *testing.T) {
	tests := []struct {
		name               string
		err                error
		expectedStatusCode int
		expectedCode       string
		expectedMessage    string
	}{
		{
			name:               "invalid request body maps to bad request",
			err:                apperror.New(apperror.CodeInvalidRequestBody, "invalid request body"),
			expectedStatusCode: http.StatusBadRequest,
			expectedCode:       string(apperror.CodeInvalidRequestBody),
			expectedMessage:    "invalid request body",
		},
		{
			name:               "validation error maps to bad request",
			err:                apperror.New(apperror.CodeValidation, "validation failed"),
			expectedStatusCode: http.StatusBadRequest,
			expectedCode:       string(apperror.CodeValidation),
			expectedMessage:    "validation failed",
		},
		{
			name:               "event source required maps to bad request",
			err:                apperror.New(apperror.CodeEventSourceRequired, "source is required"),
			expectedStatusCode: http.StatusBadRequest,
			expectedCode:       string(apperror.CodeEventSourceRequired),
			expectedMessage:    "source is required",
		},
		{
			name:               "event type required maps to bad request",
			err:                apperror.New(apperror.CodeEventTypeRequired, "event_type is required"),
			expectedStatusCode: http.StatusBadRequest,
			expectedCode:       string(apperror.CodeEventTypeRequired),
			expectedMessage:    "event_type is required",
		},
		{
			name:               "event entity id required maps to bad request",
			err:                apperror.New(apperror.CodeEventEntityIDRequired, "entity_id is required"),
			expectedStatusCode: http.StatusBadRequest,
			expectedCode:       string(apperror.CodeEventEntityIDRequired),
			expectedMessage:    "entity_id is required",
		},
		{
			name:               "event occurred at missing maps to bad request",
			err:                apperror.New(apperror.CodeEventOccurredAtMissing, "occurred_at is required"),
			expectedStatusCode: http.StatusBadRequest,
			expectedCode:       string(apperror.CodeEventOccurredAtMissing),
			expectedMessage:    "occurred_at is required",
		},
		{
			name:               "unsupported event type maps to bad request",
			err:                apperror.New(apperror.CodeUnsupportedEventType, "unsupported event_type"),
			expectedStatusCode: http.StatusBadRequest,
			expectedCode:       string(apperror.CodeUnsupportedEventType),
			expectedMessage:    "unsupported event_type",
		},
		{
			name:               "not found maps to not found",
			err:                apperror.New(apperror.CodeNotFound, "event not found"),
			expectedStatusCode: http.StatusNotFound,
			expectedCode:       string(apperror.CodeNotFound),
			expectedMessage:    "event not found",
		},
		{
			name:               "conflict maps to conflict",
			err:                apperror.New(apperror.CodeConflict, "conflict detected"),
			expectedStatusCode: http.StatusConflict,
			expectedCode:       string(apperror.CodeConflict),
			expectedMessage:    "conflict detected",
		},
		{
			name:               "event already exists maps to conflict",
			err:                apperror.New(apperror.CodeEventAlreadyExists, "event already exists"),
			expectedStatusCode: http.StatusConflict,
			expectedCode:       string(apperror.CodeEventAlreadyExists),
			expectedMessage:    "event already exists",
		},
		{
			name:               "unknown app error maps to internal server error",
			err:                apperror.New(apperror.CodeEventPublishFailed, "publish failed"),
			expectedStatusCode: http.StatusInternalServerError,
			expectedCode:       string(apperror.CodeInternal),
			expectedMessage:    "internal server error",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mappedErr := Map(tt.err)

			assert.Equal(t, tt.expectedStatusCode, mappedErr.StatusCode)
			assert.Equal(t, tt.expectedCode, mappedErr.Code)
			assert.Equal(t, tt.expectedMessage, mappedErr.Message)
		})
	}
}

func TestMap_NonAppError(t *testing.T) {
	mappedErr := Map(errors.New("unexpected failure"))

	assert.Equal(t, http.StatusInternalServerError, mappedErr.StatusCode)
	assert.Equal(t, string(apperror.CodeInternal), mappedErr.Code)
	assert.Equal(t, "internal server error", mappedErr.Message)
}
