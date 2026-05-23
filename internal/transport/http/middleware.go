package http

import (
	"net/http"
	"time"

	"github.com/Thinura/go-eventops-platform/internal/infrastructure/logger"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func registerMiddlewares(router chi.Router, config RouterConfig) {
	router.Use(middleware.RequestID)
	router.Use(middleware.Recoverer)
	router.Use(middleware.RealIP)

	if config.AppLogging {
		router.Use(requestLogger)
	}
}

func requestLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		wrapped := &responseWriter{
			ResponseWriter: w,
			statusCode:     http.StatusOK,
		}

		next.ServeHTTP(wrapped, r)

		logger.Info(
			"http request completed",
			"request_id", middleware.GetReqID(r.Context()),
			"method", r.Method,
			"path", r.URL.Path,
			"status", wrapped.statusCode,
			"duration_ms", time.Since(start).Milliseconds(),
		)
	})
}

type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}