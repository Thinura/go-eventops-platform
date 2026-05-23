package event

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/Thinura/go-eventops-platform/internal/usecase"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHandler_Create_ValidRequest(t *testing.T) {
	handler := NewHandler(usecase.NewIngestEventUseCase())

	body := `{
		"source": "payment-service",
		"event_type": "payment.failed",
		"entity_id": "order-1001",
		"payload": {
			"amount": 4500,
			"reason": "card_declined"
		},
		"occurred_at": "2026-05-23T10:30:00Z"
	}`

	req := httptest.NewRequest(http.MethodPost, "/api/v1/events", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()

	handler.Create(rec, req)

	require.Equal(t, http.StatusAccepted, rec.Code, "body: %s", rec.Body.String())

	var response struct {
		ID     string `json:"id"`
		Status string `json:"status"`
	}

	require.NoError(t, json.NewDecoder(rec.Body).Decode(&response))
	assert.NotEmpty(t, response.ID)
	assert.Equal(t, "accepted", response.Status)
}

func TestHandler_Create_InvalidJSON(t *testing.T) {
	handler := NewHandler(usecase.NewIngestEventUseCase())

	body := `{"source": "payment-service",`

	req := httptest.NewRequest(http.MethodPost, "/api/v1/events", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()

	handler.Create(rec, req)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
}

func TestHandler_Create_UnsupportedEventType(t *testing.T) {
	handler := NewHandler(usecase.NewIngestEventUseCase())

	body := `{
		"source": "payment-service",
		"event_type": "wrong.type",
		"entity_id": "order-1001",
		"payload": {},
		"occurred_at": "2026-05-23T10:30:00Z"
	}`

	req := httptest.NewRequest(http.MethodPost, "/api/v1/events", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()

	handler.Create(rec, req)

	require.Equal(t, http.StatusBadRequest, rec.Code, "body: %s", rec.Body.String())

	var response struct {
		Error struct {
			Code    string `json:"code"`
			Message string `json:"message"`
		} `json:"error"`
	}

	require.NoError(t, json.NewDecoder(rec.Body).Decode(&response))
	assert.Equal(t, "UNSUPPORTED_EVENT_TYPE", response.Error.Code)
}

func TestHandler_Create_MissingSource(t *testing.T) {
	handler := NewHandler(usecase.NewIngestEventUseCase())

	body := `{
		"event_type": "payment.failed",
		"entity_id": "order-1001",
		"payload": {},
		"occurred_at": "2026-05-23T10:30:00Z"
	}`

	req := httptest.NewRequest(http.MethodPost, "/api/v1/events", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()

	handler.Create(rec, req)

	require.Equal(t, http.StatusBadRequest, rec.Code, "body: %s", rec.Body.String())

	var response struct {
		Error struct {
			Code    string `json:"code"`
			Message string `json:"message"`
		} `json:"error"`
	}

	require.NoError(t, json.NewDecoder(rec.Body).Decode(&response))
	assert.Equal(t, "EVENT_SOURCE_REQUIRED", response.Error.Code)
}
