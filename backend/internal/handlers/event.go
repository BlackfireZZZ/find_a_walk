package handlers

import (
	"context"
	"find_a_walk/internal/domain"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/google/uuid"
)

type EventService interface {
	GetEventByID(ctx context.Context, id uuid.UUID) (*domain.Event, error)
	CreateEvent(ctx context.Context, event *domain.EventIn) (*domain.Event, error)
	GetEvents(ctx context.Context) ([]*domain.Event, error)
	// GetEventTags(ctx context.Context, id int) ([]*domain.Tag, error)
	// GetEventMembers(ctx context.Context, eventID int) ([]*domain.User, error)
}

type EventHandler struct {
	service EventService
}

func NewEventHandler(service EventService) *EventHandler {
	return &EventHandler{service: service}
}

func (h *EventHandler) GetEvents(w http.ResponseWriter, r *http.Request) {
	events, err := h.service.GetEvents(r.Context())
	if err != nil {
		render.Render(w, r, domain.ErrInvalidRequest(err, http.StatusInternalServerError))
	}

	render.RenderList(w, r, newEventList(events))
}

func newEventList(events []*domain.Event) []render.Renderer {
	list := []render.Renderer{}
	for _, event := range events {
		list = append(list, event)
	}
	return list
}

func (h *EventHandler) GetEventByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	eventID := uuid.New()
	var err error
	if eventID, err = uuid.Parse(id); err != nil {
		render.Render(w, r, domain.ErrInvalidRequest(err, http.StatusBadRequest))
		return
	}

	event, err := h.service.GetEventByID(r.Context(), eventID)
	if err != nil {
		render.Render(w, r, domain.ErrInvalidRequest(err, http.StatusNotFound))
		return
	}

	render.Render(w, r, event)
}

func (h *EventHandler) CreateEvent(w http.ResponseWriter, r *http.Request) {
	event := &domain.EventIn{}

	var err error
	if err = render.Bind(r, event); err != nil {
		render.Render(w, r, domain.ErrInvalidRequest(err, http.StatusBadRequest))
		return
	}

	var eventSchema *domain.Event
	eventSchema, err = h.service.CreateEvent(r.Context(), event)
	if err != nil {
		render.Render(w, r, domain.ErrInvalidRequest(err, http.StatusInternalServerError))
		return
	}

	render.Status(r, http.StatusCreated)
	render.Render(w, r, eventSchema)
}
