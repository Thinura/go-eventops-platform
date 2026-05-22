package http

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

func NewRouter() http.Handler {
	r := chi.NewRouter()
	r.Route("/api/v1", func(r chi.Router) {
		r.Get("/health", HealthHandler)
	})
	return r
}
