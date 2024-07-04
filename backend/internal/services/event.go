package services

import (
	"context"
	"find_a_walk/internal/domain"
	"github.com/google/uuid"
)

type EventRepository interface {
	GetEventByID(ctx context.Context, id uuid.UUID) (*domain.Event, error)
	CreateEvent(ctx context.Context, event *domain.EventIn) (*domain.Event, error)
	GetEvents(ctx context.Context, tags []string) ([]*domain.Event, error)
	GetEventsByAnglesCoordinates(ctx context.Context, lon1, lat1, lon2, lat2 float64, tags []string) ([]*domain.Event, error)
	DeleteExpiredEvents(ctx context.Context) error
	// GetEventTags(ctx context.Context, id uuid.UUID) ([]*domain.Tag, error)
	// GetEventMembers(ctx context.Context, eventID int) ([]*domain.User, error)
}

// Реализация сервиса
type EventService struct {
	repo EventRepository
}

func NewDefaultEventService(repo EventRepository) *EventService {
	return &EventService{repo: repo}
}

func (s *EventService) GetEventByID(ctx context.Context, id uuid.UUID) (*domain.Event, error) {
	return s.repo.GetEventByID(ctx, id)
}

func (s *EventService) CreateEvent(ctx context.Context, event *domain.EventIn) (*domain.Event, error) {
	return s.repo.CreateEvent(ctx, event)
}

func (s *EventService) GetEvents(ctx context.Context, tags []string) ([]*domain.Event, error) {
	return s.repo.GetEvents(ctx, tags)
}

func (s *EventService) GetEventsByAnglesCoordinates(ctx context.Context, lon1, lat1, lon2, lat2 float64, tags []string) ([]*domain.Event, error) {
	return s.repo.GetEventsByAnglesCoordinates(ctx, lon1, lat1, lon2, lat2, tags)
}

func (s *EventService) DeleteExpiredEvents(ctx context.Context) error {
	return s.repo.DeleteExpiredEvents(ctx)
}

// func (s *EventService) GetEventTags(ctx context.Context, id uuid.UUID) ([]*domain.Tag, error) {
// 	return s.repo.GetEventTags(ctx, id)
// }

// func (s *EventService) GetEventMembers(ctx context.Context, eventID int) ([]*domain.User, error) {
// 	return s.repo.GetEventMembers(ctx, eventID)
// }
