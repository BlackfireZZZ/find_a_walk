package services

import (
	"context"
<<<<<<< HEAD
	"github.com/BlackfireZZZ/find_a_walk/internal/domain"
=======

	"find_a_walk/internal/domain"
>>>>>>> 9a73157d27c4455022c57b14e97da788acc1039c
)

type eventRepository interface {
	GeteventByID(ctx context.Context, id int) (*domain.event, error)
	Createevent(ctx context.Context, event *domain.event) error
}

// Реализация сервиса
type eventService struct {
	repo eventRepository
}

func NewDefaulteventService(repo eventRepository) *eventService {
	return &eventService{repo: repo}
}

func (s *eventService) GeteventByID(ctx context.Context, id int) (*domain.event, error) {
	return s.repo.GeteventByID(ctx, id)
}

func (s *eventService) Createevent(ctx context.Context, event *domain.event) error {
	return s.repo.Createevent(ctx, event)
}
