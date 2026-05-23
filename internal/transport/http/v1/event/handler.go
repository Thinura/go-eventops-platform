package event

import (
	"encoding/json"
	"net/http"
	"time"

		"github.com/Thinura/go-eventops-platform/internal/infrastructure/logger"

	"github.com/Thinura/go-eventops-platform/internal/apperror"
	"github.com/Thinura/go-eventops-platform/internal/transport/http/response"
	"github.com/Thinura/go-eventops-platform/internal/usecase"
)

type Handler struct {
	ingestEventUseCase *usecase.IngestEventUseCase
}

func NewHandler(ingestEventUseCase *usecase.IngestEventUseCase) *Handler {
	return &Handler{
		ingestEventUseCase: ingestEventUseCase,
	}
}

type createEventRequest struct {
	Source     string         `json:"source"`
	EventType  string         `json:"event_type"`
	EntityID   string         `json:"entity_id"`
	Payload    map[string]any `json:"payload"`
	OccurredAt time.Time      `json:"occurred_at"`
}

type createEventResponse struct {
	ID     string `json:"id"`
	Status string `json:"status"`
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	var request createEventRequest

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(&request); err != nil {
		response.AppError(w, apperror.New(
			apperror.CodeInvalidRequestBody,
			"invalid request body",
		))
		return
	}

	logger.DebugJSON("Creating event", "payload", request)
	output, err := h.ingestEventUseCase.Execute(r.Context(), usecase.IngestEventInput{
		Source:     request.Source,
		EventType:  request.EventType,
		EntityID:   request.EntityID,
		Payload:    request.Payload,
		OccurredAt: request.OccurredAt,
	})
	if err != nil {
		response.AppError(w, err)
		return
	}

	response.JSON(w, http.StatusAccepted, createEventResponse{
		ID:     output.ID,
		Status: output.Status,
	})
}
