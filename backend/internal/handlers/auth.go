package handlers

import (
	"find_a_walk/internal/domain"
	"net/http"

	"github.com/go-chi/render"
)

type AuthHandler struct {
	service UserService
}

func NewAuthHandler(service UserService) *AuthHandler {
	return &AuthHandler{service: service}
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	userSchema := &domain.UserAuth{}
	render.Bind(r, userSchema)

	var token *domain.Token
	token, err := h.service.Login(r.Context(), userSchema)
	if err != nil {
		render.Render(w, r, domain.ErrInvalidRequest(err, http.StatusInternalServerError))
	}

	render.Render(w, r, token)
}
