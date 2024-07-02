package handlers

import (
	"context"
	"encoding/json"
	"find_a_walk/internal/domain"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
)

type EventService interface {
	GetEventByID(ctx context.Context, id int) (*domain.Event, error)
	CreateEvent(ctx context.Context, event *domain.Event) error
	GetEventTags(ctx context.Context, id int) ([]*domain.Tag, error)
	GetEventMembers(ctx context.Context, eventID int) ([]*domain.User, error)
}

// Обработчики HTTP запросов
type EventHandler struct {
	service EventService
}

func NewEventHandler(service EventService) *EventHandler {

	return &EventHandler{service: service}
}

func (h *EventHandler) GetEventByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	eventID := 0
	var err error
	if eventID, err = strconv.Atoi(id); err != nil {
		http.Error(w, "Invalid event ID", http.StatusBadRequest)
		return
	}

	event, err := h.service.GetEventByID(r.Context(), eventID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(event)
}

func (h *EventHandler) CreateEvent(w http.ResponseWriter, r *http.Request) {
	var event domain.Event
	err := json.NewDecoder(r.Body).Decode(&event)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.service.CreateEvent(r.Context(), &event)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *EventHandler) GetEventTags(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	eventID := 0
	var err error
	if eventID, err = strconv.Atoi(id); err != nil {
		http.Error(w, "Invalid event ID", http.StatusBadRequest)
		return
	}

	tags, err := h.service.GetEventTags(r.Context(), eventID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tags)
}

func (h *EventHandler) GetEventMembers(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	eventID := 0
	var err error
	if eventID, err = strconv.Atoi(id); err != nil {
		http.Error(w, "Invalid event ID", http.StatusBadRequest)
		return
	}

	members, err := h.service.GetEventMembers(r.Context(), eventID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(members)
}
