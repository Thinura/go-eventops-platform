package event

import (
	"github.com/Thinura/go-eventops-platform/internal/usecase"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func Routes() http.Handler {
	router := chi.NewRouter()

	ingestEventUseCase := usecase.NewIngestEventUseCase()
	handler := NewHandler(ingestEventUseCase)

	router.Post("/", handler.Create)

	return router
}
