package handlers

import (
	"context"
	"find_a_walk/internal/domain"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
	"github.com/go-chi/render"
	"github.com/google/uuid"
)

type EventService interface {
	CreateEvent(ctx context.Context, event *domain.EventIn) (*domain.Event, error)
	DeleteEvent(ctx context.Context, id uuid.UUID, userID uuid.UUID) error
	GetEventByID(ctx context.Context, id uuid.UUID) (*domain.Event, error)

	GetEvents(ctx context.Context, tags []string) ([]*domain.Event, error)
	GetEventsByUserID(ctx context.Context, userID uuid.UUID) ([]*domain.Event, error)
	GetEventsByAnglesCoordinates(ctx context.Context, lon1, lat1, lon2, lat2 float64, tags []string) ([]*domain.Event, error)

	CreateEventMember(ctx context.Context, eventID uuid.UUID, userID uuid.UUID) error
	DeleteEventMember(ctx context.Context, eventID uuid.UUID, userID uuid.UUID) error
	// GetEventMembers(ctx context.Context, eventID uuid.UUID) ([]*domain.User, error)
	GetMyEventMembers(ctx context.Context, userID uuid.UUID) ([]*domain.Event, error)
}

type EventHandler struct {
	service EventService
}

func NewEventHandler(service EventService) *EventHandler {
	return &EventHandler{service: service}
}

func (h *EventHandler) DeleteEvent(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	var eventID uuid.UUID
	var err error
	if eventID, err = uuid.Parse(id); err != nil {
		render.Render(w, r, domain.ErrInvalidRequest(err, http.StatusBadRequest))
		return
	}

	userID, err := getUserIDFromContext(r)
	if err != nil {
		render.Render(w, r, domain.ErrInvalidRequest(err, http.StatusUnauthorized))
		return
	}

	err = h.service.DeleteEvent(r.Context(), eventID, userID)
	if err != nil {
		render.Render(w, r, domain.ErrInvalidRequest(err, http.StatusBadRequest))
		return
	}

	render.Render(w, r, domain.OKRequest("ok", http.StatusNoContent))
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
			return
		}
		render.RenderList(w, r, newEventList(events))
	} else {
		events, err := h.service.GetEvents(r.Context(), tags)
		if err != nil {
			render.Render(w, r, domain.ErrInvalidRequest(err, http.StatusInternalServerError))
			return
		}
		render.RenderList(w, r, newEventList(events))
	}
}

func (h *EventHandler) GetMyEvents(w http.ResponseWriter, r *http.Request) {
	userID, err := getUserIDFromContext(r)
	if err != nil {
		render.Render(w, r, domain.ErrInvalidRequest(err, http.StatusUnauthorized))
		return
	}

	events, err := h.service.GetEventsByUserID(r.Context(), userID)
	if err != nil {
		render.Render(w, r, domain.ErrInvalidRequest(err, http.StatusBadRequest))
		return
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

	var eventID uuid.UUID
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

func getUserIDFromContext(r *http.Request) (uuid.UUID, error) {
	_, claims, err := jwtauth.FromContext(r.Context())

	if err != nil {
		return uuid.UUID{}, err
	}

	userID, err := uuid.Parse(fmt.Sprintf("%v", claims["user_id"]))
	if err != nil {
		return uuid.UUID{}, err
	}
	return userID, nil

}

func (h *EventHandler) CreateEvent(w http.ResponseWriter, r *http.Request) {
	userID, err := getUserIDFromContext(r)
	if err != nil {
		render.Render(w, r, domain.ErrInvalidRequest(err, http.StatusUnauthorized))
		return
	}

	event := &domain.EventIn{AuthorID: userID}

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

func (h *EventHandler) CreateEventMember(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	var eventID uuid.UUID
	var err error
	if eventID, err = uuid.Parse(id); err != nil {
		render.Render(w, r, domain.ErrInvalidRequest(err, http.StatusBadRequest))
		return
	}

	userID, err := getUserIDFromContext(r)
	if err != nil {
		render.Render(w, r, domain.ErrInvalidRequest(err, http.StatusUnauthorized))
		return
	}

	err = h.service.CreateEventMember(r.Context(), eventID, userID)
	if err != nil {
		render.Render(w, r, domain.ErrInvalidRequest(err, http.StatusBadRequest))
		return
	}

	render.Render(w, r, domain.OKRequest("ok", http.StatusCreated))
}

func (h *EventHandler) DeleteEventMember(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	var eventID uuid.UUID
	var err error
	if eventID, err = uuid.Parse(id); err != nil {
		render.Render(w, r, domain.ErrInvalidRequest(err, http.StatusBadRequest))
		return
	}

	userID, err := getUserIDFromContext(r)
	if err != nil {
		render.Render(w, r, domain.ErrInvalidRequest(err, http.StatusUnauthorized))
		return
	}

	err = h.service.DeleteEventMember(r.Context(), eventID, userID)
	if err != nil {
		render.Render(w, r, domain.ErrInvalidRequest(err, http.StatusBadRequest))
		return
	}

	render.Render(w, r, domain.OKRequest("ok", http.StatusNoContent))
}

// func (h *EventHandler) GetEventMembers(w http.ResponseWriter, r *http.Request) {
// }

func (h *EventHandler) GetMyEventMembers(w http.ResponseWriter, r *http.Request) {
	userID, err := getUserIDFromContext(r)
	if err != nil {
		render.Render(w, r, domain.ErrInvalidRequest(err, http.StatusUnauthorized))
		return
	}

	var events []*domain.Event
	events, err = h.service.GetMyEventMembers(r.Context(), userID)
	if err != nil {
		render.Render(w, r, domain.ErrInvalidRequest(err, http.StatusBadRequest))
		return
	}

	render.RenderList(w, r, newEventList(events))
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
