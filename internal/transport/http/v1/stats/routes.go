package stats

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

func StatsRoutes() http.Handler {
	router := chi.NewRouter()
	router.Get("/", statsHandler)
	return router
}
