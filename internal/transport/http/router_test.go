package http

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewRouter_MountsV1HealthRoute(t *testing.T) {
	router := NewRouter(RouterConfig{
		AppLogging: false,
	})

	req := httptest.NewRequest(http.MethodGet, "/api/v1/health", nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	require.Equal(t, http.StatusOK, rec.Code, "body: %s", rec.Body.String())
}

func TestNewRouter_MountsV1EventRoute(t *testing.T) {
	router := NewRouter(RouterConfig{
		AppLogging: false,
	})

	body := `{
		"source": "payment-service",
		"event_type": "payment.failed",
		"entity_id": "order-1001",
		"payload": {},
		"occurred_at": "2026-05-23T10:30:00Z"
	}`

	req := httptest.NewRequest(http.MethodPost, "/api/v1/events", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	require.Equal(t, http.StatusAccepted, rec.Code, "body: %s", rec.Body.String())
}

func TestNewRouter_ReturnsNotFoundForUnknownRoute(t *testing.T) {
	router := NewRouter(RouterConfig{
		AppLogging: false,
	})

	req := httptest.NewRequest(http.MethodGet, "/unknown", nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	require.Equal(t, http.StatusNotFound, rec.Code)
}

func TestNewRouter_WithAppLoggingEnabled(t *testing.T) {
	router := NewRouter(RouterConfig{
		AppLogging: true,
	})

	req := httptest.NewRequest(http.MethodGet, "/api/v1/health", nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	require.Equal(t, http.StatusOK, rec.Code, "body: %s", rec.Body.String())
}