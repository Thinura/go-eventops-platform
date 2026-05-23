package health

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

func HealthRoutes() http.Handler {
	router := chi.NewRouter()
	router.Get("/", healthHandler)
	return router
}
