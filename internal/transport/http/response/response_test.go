package response

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestJSON_WritesStatusContentTypeAndBody(t *testing.T) {
	rec := httptest.NewRecorder()

	payload := map[string]any{
		"status":  "ok",
		"service": "eventops-api",
	}

	JSON(rec, http.StatusCreated, payload)

	assert.Equal(t, http.StatusCreated, rec.Code)
	assert.Equal(t, "application/json", rec.Header().Get("Content-Type"))

	var response map[string]any
	require.NoError(t, json.NewDecoder(rec.Body).Decode(&response))
	assert.Equal(t, "ok", response["status"])
	assert.Equal(t, "eventops-api", response["service"])
}

func TestJSON_WritesNilPayloadAsNull(t *testing.T) {
	rec := httptest.NewRecorder()

	JSON(rec, http.StatusOK, nil)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, "application/json", rec.Header().Get("Content-Type"))
	assert.JSONEq(t, `null`, rec.Body.String())
}

func TestError_WritesErrorResponse(t *testing.T) {
	rec := httptest.NewRecorder()

	Error(rec, http.StatusBadRequest, "INVALID_REQUEST_BODY", "invalid request body")

	assert.Equal(t, http.StatusBadRequest, rec.Code)
	assert.Equal(t, "application/json", rec.Header().Get("Content-Type"))

	var response ErrorResponse
	require.NoError(t, json.NewDecoder(rec.Body).Decode(&response))
	assert.Equal(t, "INVALID_REQUEST_BODY", response.Error.Code)
	assert.Equal(t, "invalid request body", response.Error.Message)
}

func TestError_ResponseShape(t *testing.T) {
	rec := httptest.NewRecorder()

	Error(rec, http.StatusNotFound, "NOT_FOUND", "event not found")

	expected := `{
		"error": {
			"code": "NOT_FOUND",
			"message": "event not found"
		}
	}`

	assert.JSONEq(t, expected, rec.Body.String())
}
