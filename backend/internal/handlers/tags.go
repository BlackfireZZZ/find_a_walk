package handlers

import (
	"context"
	"find_a_walk/internal/domain"
	"net/http"

	"github.com/go-chi/render"
)

type TagsService interface {
	GetTags(ctx context.Context) ([]*domain.Tag, error)
}

type TagsHandler struct {
	service TagsService
}

func NewTagsHandler(service TagsService) *TagsHandler {
	return &TagsHandler{service: service}
}

func (h *TagsHandler) GetTags(w http.ResponseWriter, r *http.Request) {
	tags, err := h.service.GetTags(r.Context())
	if err != nil {
		render.Render(w, r, domain.ErrInvalidRequest(err, http.StatusNotImplemented))
	}

	render.RenderList(w, r, newTagList(tags))
}

func newTagList(tags []*domain.Tag) []render.Renderer {
	list := []render.Renderer{}
	for _, tag := range tags {
		list = append(list, tag)
	}
	return list
}
