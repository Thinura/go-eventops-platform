
package response

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Thinura/go-eventops-platform/internal/apperror"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAppError_WritesMappedValidationError(t *testing.T) {
	rec := httptest.NewRecorder()

	AppError(rec, apperror.Validation(
		apperror.CodeEventSourceRequired,
		"source is required",
	))

	assert.Equal(t, http.StatusBadRequest, rec.Code)
	assert.Equal(t, "application/json", rec.Header().Get("Content-Type"))

	var response ErrorResponse
	require.NoError(t, json.NewDecoder(rec.Body).Decode(&response))
	assert.Equal(t, string(apperror.CodeEventSourceRequired), response.Error.Code)
	assert.Equal(t, "source is required", response.Error.Message)
}

func TestAppError_WritesMappedNotFoundError(t *testing.T) {
	rec := httptest.NewRecorder()

	AppError(rec, apperror.NotFound("event not found"))

	assert.Equal(t, http.StatusNotFound, rec.Code)
	assert.Equal(t, "application/json", rec.Header().Get("Content-Type"))

	var response ErrorResponse
	require.NoError(t, json.NewDecoder(rec.Body).Decode(&response))
	assert.Equal(t, string(apperror.CodeNotFound), response.Error.Code)
	assert.Equal(t, "event not found", response.Error.Message)
}

func TestAppError_WritesMappedConflictError(t *testing.T) {
	rec := httptest.NewRecorder()

	AppError(rec, apperror.Conflict(
		apperror.CodeEventAlreadyExists,
		"event already exists",
	))

	assert.Equal(t, http.StatusConflict, rec.Code)
	assert.Equal(t, "application/json", rec.Header().Get("Content-Type"))

	var response ErrorResponse
	require.NoError(t, json.NewDecoder(rec.Body).Decode(&response))
	assert.Equal(t, string(apperror.CodeEventAlreadyExists), response.Error.Code)
	assert.Equal(t, "event already exists", response.Error.Message)
}

func TestAppError_WritesInternalServerErrorForNonAppError(t *testing.T) {
	rec := httptest.NewRecorder()

	AppError(rec, errors.New("unexpected failure"))

	assert.Equal(t, http.StatusInternalServerError, rec.Code)
	assert.Equal(t, "application/json", rec.Header().Get("Content-Type"))

	var response ErrorResponse
	require.NoError(t, json.NewDecoder(rec.Body).Decode(&response))
	assert.Equal(t, string(apperror.CodeInternal), response.Error.Code)
	assert.Equal(t, "internal server error", response.Error.Message)
}

func TestAppError_WritesInternalServerErrorForUnknownAppError(t *testing.T) {
	rec := httptest.NewRecorder()

	AppError(rec, apperror.New(
		apperror.CodeEventPublishFailed,
		"failed to publish event",
	))

	assert.Equal(t, http.StatusInternalServerError, rec.Code)
	assert.Equal(t, "application/json", rec.Header().Get("Content-Type"))

	var response ErrorResponse
	require.NoError(t, json.NewDecoder(rec.Body).Decode(&response))
	assert.Equal(t, string(apperror.CodeInternal), response.Error.Code)
	assert.Equal(t, "internal server error", response.Error.Message)
}
