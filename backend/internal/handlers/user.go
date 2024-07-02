package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"

	"hse_school/internal/domain"
)

type UserService interface {
	GetUserByID(ctx context.Context, id int) (*domain.User, error)
	CreateUser(ctx context.Context, user *domain.User) error
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

	userID := 0
	var err error
	if userID, err = strconv.Atoi(id); err != nil {
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
