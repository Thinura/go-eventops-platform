package http

import (
	"github.com/go-chi/chi/v5"
	"net/http"

	v1 "github.com/Thinura/go-eventops-platform/internal/transport/http/v1"
)

func NewRouter() http.Handler {
	router := chi.NewRouter()
	registerMiddlewares(router)
	router.Mount("/api/v1", v1.Router())
	return router
}
