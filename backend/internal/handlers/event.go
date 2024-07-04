package handlers

import (
	"context"
	"find_a_walk/internal/domain"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/google/uuid"
)

type EventService interface {
	GetEventByID(ctx context.Context, id uuid.UUID) (*domain.Event, error)
	CreateEvent(ctx context.Context, event *domain.EventIn) (*domain.Event, error)
	GetEvents(ctx context.Context, tags []string) ([]*domain.Event, error)
	GetEventsByAnglesCoordinates(ctx context.Context, lon1, lat1, lon2, lat2 float64, tags []string) ([]*domain.Event, error)
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
	query := r.URL.Query()
	tags := query["tags"]
	if _, ok := query["lat1"]; ok {
		lat1 := query["lat1"]
		lat2 := query["lat2"]
		lon1 := query["lon1"]
		lon2 := query["lon2"]
		coordinates, err := StrToFloat64([]string{lat1[0], lat2[0], lon1[0], lon2[0]})
		if err != nil {
			render.Render(w, r, domain.ErrInvalidRequest(err, http.StatusBadRequest))
			return
		}
		events, err := h.service.GetEventsByAnglesCoordinates(r.Context(), coordinates[0], coordinates[1], coordinates[2], coordinates[3], tags)
		if err != nil {
			render.Render(w, r, domain.ErrInvalidRequest(err, http.StatusInternalServerError))
		}
		render.RenderList(w, r, newEventList(events))
	} else {
		events, err := h.service.GetEvents(r.Context(), tags)
		if err != nil {
			render.Render(w, r, domain.ErrInvalidRequest(err, http.StatusInternalServerError))
		}
		render.RenderList(w, r, newEventList(events))
	}
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

func StrToFloat64(list []string) ([]float64, error) {
	var res []float64
	for _, str := range list {
		f, err := strconv.ParseFloat(str, 64)
		if err != nil {
			return nil, err
		}
		res = append(res, f)
	}
	return res, nil
}
