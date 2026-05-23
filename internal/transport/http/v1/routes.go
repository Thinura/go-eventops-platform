package v1

import (
	"net/http"

	alert "github.com/Thinura/go-eventops-platform/internal/transport/http/v1/alert"
	event "github.com/Thinura/go-eventops-platform/internal/transport/http/v1/event"
	health "github.com/Thinura/go-eventops-platform/internal/transport/http/v1/health"
	stats "github.com/Thinura/go-eventops-platform/internal/transport/http/v1/stats"
	"github.com/go-chi/chi/v5"
)

func Router() http.Handler {
	router := chi.NewRouter()

	router.Mount("/health", health.HealthRoutes())
	router.Mount("/events", event.Routes())
	router.Mount("/alerts", alert.AlertRoutes())

	router.Mount("/stats", stats.StatsRoutes())

	return router
}
