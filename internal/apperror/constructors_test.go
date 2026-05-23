

package apperror

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {
	err := New(CodeValidation, "validation failed")

	require.NotNil(t, err)
	assert.Equal(t, CodeValidation, err.Code)
	assert.Equal(t, "validation failed", err.Message)
	assert.Nil(t, err.Cause)
}

func TestWrap(t *testing.T) {
	cause := errors.New("database timeout")

	err := Wrap(CodeInternal, "failed to save event", cause)

	require.NotNil(t, err)
	assert.Equal(t, CodeInternal, err.Code)
	assert.Equal(t, "failed to save event", err.Message)
	assert.Equal(t, cause, err.Cause)
	assert.True(t, errors.Is(err, cause))
}

func TestValidation(t *testing.T) {
	err := Validation(CodeEventSourceRequired, "source is required")

	require.NotNil(t, err)
	assert.Equal(t, CodeEventSourceRequired, err.Code)
	assert.Equal(t, "source is required", err.Message)
	assert.Nil(t, err.Cause)
}

func TestNotFound(t *testing.T) {
	err := NotFound("event not found")

	require.NotNil(t, err)
	assert.Equal(t, CodeNotFound, err.Code)
	assert.Equal(t, "event not found", err.Message)
	assert.Nil(t, err.Cause)
}

func TestConflict(t *testing.T) {
	err := Conflict(CodeEventAlreadyExists, "event already exists")

	require.NotNil(t, err)
	assert.Equal(t, CodeEventAlreadyExists, err.Code)
	assert.Equal(t, "event already exists", err.Message)
	assert.Nil(t, err.Cause)
}

func TestInternal(t *testing.T) {
	cause := errors.New("redis unavailable")

	err := Internal("failed to publish event", cause)

	require.NotNil(t, err)
	assert.Equal(t, CodeInternal, err.Code)
	assert.Equal(t, "failed to publish event", err.Message)
	assert.Equal(t, cause, err.Cause)
	assert.True(t, errors.Is(err, cause))
}