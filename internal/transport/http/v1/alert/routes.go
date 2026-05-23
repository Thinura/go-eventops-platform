package alert

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

func AlertRoutes() http.Handler {
	router := chi.NewRouter()
	router.Get("/", alertHandler)
	return router
}
