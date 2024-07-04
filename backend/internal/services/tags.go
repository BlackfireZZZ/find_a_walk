package services

import (
	"context"
	"find_a_walk/internal/domain"
)

type TagRepository interface {
	GetTags(ctx context.Context) ([]*domain.Tag, error)
}

type TagService struct {
	repo TagRepository
}

func NewDefaultTagService(repo TagRepository) *TagService {
	return &TagService{repo: repo}
}

func (s *TagService) GetTags(ctx context.Context) ([]*domain.Tag, error) {
	return s.repo.GetTags(ctx)
}
