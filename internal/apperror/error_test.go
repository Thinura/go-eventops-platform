

package apperror

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestError_ErrorWithoutCause(t *testing.T) {
	err := &Error{
		Code:    CodeValidation,
		Message: "validation failed",
	}

	assert.Equal(t, "VALIDATION_ERROR: validation failed", err.Error())
}

func TestError_ErrorWithCause(t *testing.T) {
	cause := errors.New("database timeout")
	err := &Error{
		Code:    CodeInternal,
		Message: "failed to save event",
		Cause:   cause,
	}

	assert.Equal(t, "INTERNAL_ERROR: failed to save event: database timeout", err.Error())
}

func TestError_UnwrapWithCause(t *testing.T) {
	cause := errors.New("redis unavailable")
	err := &Error{
		Code:    CodeInternal,
		Message: "failed to publish event",
		Cause:   cause,
	}

	assert.Equal(t, cause, err.Unwrap())
	assert.True(t, errors.Is(err, cause))
}

func TestError_UnwrapWithoutCause(t *testing.T) {
	err := &Error{
		Code:    CodeNotFound,
		Message: "event not found",
	}

	assert.Nil(t, err.Unwrap())
}

func TestError_CanBeMatchedWithErrorsAs(t *testing.T) {
	original := &Error{
		Code:    CodeConflict,
		Message: "event already exists",
	}

	wrapped := errors.Join(original)

	var appErr *Error
	require.ErrorAs(t, wrapped, &appErr)
	assert.Equal(t, CodeConflict, appErr.Code)
	assert.Equal(t, "event already exists", appErr.Message)
}