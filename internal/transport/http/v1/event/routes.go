package event

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/Thinura/go-eventops-platform/internal/infrastructure/memory"
	"github.com/Thinura/go-eventops-platform/internal/usecase"
)

func Routes() http.Handler {
	router := chi.NewRouter()

	publisher := memory.NewEventPublisher()
	ingestEventUseCase := usecase.NewIngestEventUseCase(publisher)
	handler := NewHandler(ingestEventUseCase)

	router.Post("/", handler.Create)

	return router
}
