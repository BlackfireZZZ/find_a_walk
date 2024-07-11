package handlers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
	"github.com/go-chi/render"
	"github.com/google/uuid"

	"find_a_walk/internal/domain"
)

type UserService interface {
	Login(ctx context.Context, user *domain.UserAuth) (*domain.Token, error)
	GetUserByID(ctx context.Context, id uuid.UUID) (*domain.User, error)
	CreateUser(ctx context.Context, user *domain.UserIn) (*domain.User, error)
	CreateInterests(ctx context.Context, id uuid.UUID, interests []string) (*domain.User, error)
	DeleteInterests(ctx context.Context, id uuid.UUID, interests []string) (error)
	GetJWTConfig() *jwtauth.JWTAuth
}

type UserHandler struct {
	service UserService
}

func NewUserHandler(service UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) DeleteInterests(w http.ResponseWriter, r *http.Request) {
	var UserInterest domain.UserInterestIn
	err := render.Bind(r, &UserInterest)
	if err != nil {
		render.Render(w, r, domain.ErrInvalidRequest(err, http.StatusBadRequest))
		return 
	}

	_, claims, err := jwtauth.FromContext(r.Context())
	if err != nil {
		render.Render(w, r, domain.ErrInvalidRequest(err, http.StatusNotFound))
		return
	}
	userID, err := uuid.Parse(fmt.Sprintf("%v", claims["user_id"]))
	if err != nil {
		render.Render(w, r, domain.ErrInvalidRequest(err, http.StatusNotFound))
		return
	}
	err = h.service.DeleteInterests(r.Context(), userID, UserInterest.Interests)
	if err != nil {
		render.Render(w, r, domain.ErrInvalidRequest(err, http.StatusNotFound))
		return
	}

	render.Render(w, r, &domain.OKResponse{Status: "ok", HTTPStatusCode: http.StatusNoContent})
}

func (h *UserHandler) CreateInterest(w http.ResponseWriter, r *http.Request) {
	var userInterests domain.UserInterestIn
	err := render.Bind(r, &userInterests)
	if err != nil {
		render.Render(w, r, domain.ErrInvalidRequest(err, http.StatusBadRequest))
		return
	}

	_, claims, err := jwtauth.FromContext(r.Context())
	if err != nil {
		render.Render(w, r, domain.ErrInvalidRequest(err, http.StatusNotFound))
		return
	}
	userID, err := uuid.Parse(fmt.Sprintf("%v", claims["user_id"]))
	if err != nil {
		render.Render(w, r, domain.ErrInvalidRequest(err, http.StatusNotFound))
		return
	}

	user, err := h.service.CreateInterests(r.Context(), userID, userInterests.Interests)
	if err != nil {
		render.Render(w, r, domain.ErrInvalidRequest(err, http.StatusNotFound))
		return
	}

	render.Render(w, r, user)
}

func (h *UserHandler) GetUserByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	userID, err := uuid.Parse(id)
	if err != nil {
		render.Render(w, r, domain.ErrInvalidRequest(err, http.StatusBadRequest))
		return
	}

	user, err := h.service.GetUserByID(r.Context(), userID)
	if err != nil {
		render.Render(w, r, domain.ErrInvalidRequest(err, http.StatusNotFound))
		return
	}

	render.Render(w, r, user)
}

func (h *UserHandler) GetUserProfile(w http.ResponseWriter, r *http.Request) {
	_, claims, err := jwtauth.FromContext(r.Context())
	if err != nil {
		render.Render(w, r, domain.ErrInvalidRequest(err, http.StatusNotFound))
		return
	}

	userID, err := uuid.Parse(fmt.Sprintf("%v", claims["user_id"]))
	if err != nil {
		render.Render(w, r, domain.ErrInvalidRequest(err, http.StatusNotFound))
		return
	}

	user, err := h.service.GetUserByID(r.Context(), userID)
	if err != nil {
		render.Render(w, r, domain.ErrInvalidRequest(err, http.StatusNotFound))
		return
	}

	render.Render(w, r, user)
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	userSchema := &domain.UserIn{}
	err := render.Bind(r, userSchema)
	if err != nil {
		render.Render(w, r, domain.ErrInvalidRequest(err, http.StatusBadRequest))
		return
	}

	user, err := h.service.CreateUser(r.Context(), userSchema)
	if err != nil {
		render.Render(w, r, domain.ErrInvalidRequest(err, http.StatusInternalServerError))
		return
	}

	render.Status(r, http.StatusCreated)
	render.Render(w, r, user)
}
