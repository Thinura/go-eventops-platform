package event

import (
	"github.com/go-chi/chi/v5"
	"net/http"
	"github.com/Thinura/go-eventops-platform/internal/usecase"
)

func Routes() http.Handler {
	router := chi.NewRouter()

	ingestEventUseCase := usecase.NewIngestEventUseCase()
	handler := NewHandler(ingestEventUseCase)

	router.Post("/", handler.Create)
	
	return router
}
