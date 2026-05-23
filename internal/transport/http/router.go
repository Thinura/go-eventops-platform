package http

import (
	"github.com/go-chi/chi/v5"
	"net/http"

	v1 "github.com/Thinura/go-eventops-platform/internal/transport/http/v1"
)

type RouterConfig struct {
	AppLogging bool
}

func NewRouter(config RouterConfig) http.Handler {
	router := chi.NewRouter()
	registerMiddlewares(router, config)
	router.Mount("/api/v1", v1.Router())
	return router
}
