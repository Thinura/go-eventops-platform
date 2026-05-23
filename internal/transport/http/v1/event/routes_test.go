package event

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/require"
)

func TestRoutes_CreateEvent_RootPath(t *testing.T) {
	router := Routes()

	body := `{
		"source": "payment-service",
		"event_type": "payment.failed",
		"entity_id": "order-1001",
		"payload": {},
		"occurred_at": "2026-05-23T10:30:00Z"
	}`

	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	require.Equal(t, http.StatusAccepted, rec.Code, "body: %s", rec.Body.String())
}

func TestRoutes_CreateEvent_MountedPathWithoutTrailingSlash(t *testing.T) {
	root := chi.NewRouter()
	root.Mount("/events", Routes())

	body := `{
		"source": "payment-service",
		"event_type": "payment.failed",
		"entity_id": "order-1001",
		"payload": {},
		"occurred_at": "2026-05-23T10:30:00Z"
	}`

	req := httptest.NewRequest(http.MethodPost, "/events", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()

	root.ServeHTTP(rec, req)

	require.Equal(t, http.StatusAccepted, rec.Code, "body: %s", rec.Body.String())
}

func TestRoutes_MethodNotAllowed(t *testing.T) {
	router := Routes()

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	require.Equal(t, http.StatusMethodNotAllowed, rec.Code, "body: %s", rec.Body.String())
}