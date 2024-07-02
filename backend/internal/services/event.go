package services

import (
	"context"
	"find_a_walk/internal/domain"
)

type eventRepository interface {
	GetEventByID(ctx context.Context, id int) (*domain.Event, error)
	CreateEvent(ctx context.Context, event *domain.Event) error
	GetEventTags(ctx context.Context, id int) ([]*domain.Tag, error)
}

// Реализация сервиса
type EventService struct {
	repo eventRepository
}

func NewDefaultEventService(repo eventRepository) *EventService {
	return &EventService{repo: repo}
}

func (s *EventService) GetEventByID(ctx context.Context, id int) (*domain.Event, error) {
	return s.repo.GetEventByID(ctx, id)
}

func (s *EventService) CreateEvent(ctx context.Context, event *domain.Event) error {
	return s.repo.CreateEvent(ctx, event)
}

func (s *EventService) GetEventTags(ctx context.Context, id int) ([]*domain.Tag, error) {
	return s.repo.GetEventTags(ctx, id)
}
