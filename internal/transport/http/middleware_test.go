package http

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRegisterMiddlewares_WithAppLoggingEnabled(t *testing.T) {
	router := chi.NewRouter()

	registerMiddlewares(router, RouterConfig{
		AppLogging: true,
	})

	router.Get("/test", func(w http.ResponseWriter, r *http.Request) {
		requestID := middleware.GetReqID(r.Context())
		require.NotEmpty(t, requestID)

		w.WriteHeader(http.StatusCreated)
	})

	req := httptest.NewRequest(http.MethodGet, "/test", nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusCreated, rec.Code)
}

func TestRegisterMiddlewares_WithAppLoggingDisabled(t *testing.T) {
	router := chi.NewRouter()

	registerMiddlewares(router, RouterConfig{
		AppLogging: false,
	})

	router.Get("/test", func(w http.ResponseWriter, r *http.Request) {
		requestID := middleware.GetReqID(r.Context())
		require.NotEmpty(t, requestID)

		w.WriteHeader(http.StatusOK)
	})

	req := httptest.NewRequest(http.MethodGet, "/test", nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestRequestLogger_DefaultStatusCode(t *testing.T) {
	handler := requestLogger(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("ok"))
		require.NoError(t, err)
	}))

	req := httptest.NewRequest(http.MethodGet, "/test", nil)
	rec := httptest.NewRecorder()

	handler.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, "ok", rec.Body.String())
}

func TestRequestLogger_ExplicitStatusCode(t *testing.T) {
	handler := requestLogger(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusAccepted)
	}))

	req := httptest.NewRequest(http.MethodPost, "/events", nil)
	rec := httptest.NewRecorder()

	handler.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusAccepted, rec.Code)
}

func TestResponseWriter_WriteHeaderStoresStatusCode(t *testing.T) {
	rec := httptest.NewRecorder()
	wrapped := &responseWriter{
		ResponseWriter: rec,
		statusCode:     http.StatusOK,
	}

	wrapped.WriteHeader(http.StatusCreated)

	assert.Equal(t, http.StatusCreated, wrapped.statusCode)
	assert.Equal(t, http.StatusCreated, rec.Code)
}
