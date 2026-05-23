package event

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/Thinura/go-eventops-platform/internal/domain"
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
		response.Error(w, http.StatusBadRequest, "invalid request body")
		return
	}

	output, err := h.ingestEventUseCase.Execute(r.Context(), usecase.IngestEventInput{
		Source:     request.Source,
		EventType:  request.EventType,
		EntityID:   request.EntityID,
		Payload:    request.Payload,
		OccurredAt: request.OccurredAt,
	})
	if err != nil {
		writeCreateEventError(w, err)
		return
	}

	response.JSON(w, http.StatusAccepted, createEventResponse{
		ID:     output.ID,
		Status: output.Status,
	})
}

func writeCreateEventError(w http.ResponseWriter, err error) {
	switch {
	case errors.Is(err, domain.ErrEventSourceRequired):
		response.Error(w, http.StatusBadRequest, "source is required")
	case errors.Is(err, domain.ErrEventTypeRequired):
		response.Error(w, http.StatusBadRequest, "event_type is required")
	case errors.Is(err, domain.ErrUnsupportedEventType):
		response.Error(w, http.StatusBadRequest, "unsupported event_type")
	case errors.Is(err, domain.ErrEventEntityIDRequired):
		response.Error(w, http.StatusBadRequest, "entity_id is required")
	case errors.Is(err, domain.ErrEventOccurredAtMissing):
		response.Error(w, http.StatusBadRequest, "occurred_at is required")
	default:
		response.Error(w, http.StatusInternalServerError, "failed to create event")
	}
}