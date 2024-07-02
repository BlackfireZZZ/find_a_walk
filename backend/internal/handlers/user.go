package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"

	"find_a_walk/internal/domain"
)

type UserService interface {
	GetUserByID(ctx context.Context, id uuid.UUID) (*domain.User, error)
	CreateUser(ctx context.Context, user *domain.User) error
	GetUserInterests(ctx context.Context, id uuid.UUID) ([]*domain.Interest, error)
}

// Обработчики HTTP запросов
type UserHandler struct {
	service UserService
}

func NewUserHandler(service UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) GetUserByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	userID := uuid.New()
	var err error
	if userID, err = uuid.Parse(id); err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	user, err := h.service.GetUserByID(r.Context(), userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user domain.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.service.CreateUser(r.Context(), &user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *UserHandler) GetUserInterests(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	userID := uuid.New()
	var err error
	if userID, err = uuid.Parse(id); err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}
	interests, err := h.service.GetUserInterests(r.Context(), userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(interests)
}
