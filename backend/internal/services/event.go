package services

import (
	"context"
	"errors"
	"find_a_walk/internal/domain"

	"github.com/google/uuid"
)

type EventRepository interface {
	GetEventByID(ctx context.Context, id uuid.UUID) (*domain.Event, error)
	CreateEvent(ctx context.Context, event *domain.EventIn) (*domain.Event, error)
	DeleteEvent(ctx context.Context, id uuid.UUID) error
	GetEvents(ctx context.Context, tags []string) ([]*domain.Event, error)
	GetEventsByAnglesCoordinates(ctx context.Context, lon1, lat1, lon2, lat2 float64, tags []string) ([]*domain.Event, error)
	DeleteExpiredEvents(ctx context.Context) error
}

// Реализация сервиса
type EventService struct {
	repo EventRepository
}

func NewDefaultEventService(repo EventRepository) *EventService {
	return &EventService{repo: repo}
}

func (s *EventService) DeleteEvent(ctx context.Context, id uuid.UUID, userID uuid.UUID) error {
	event, err := s.repo.GetEventByID(ctx, id)
	if err != nil {
		return err
	}

	if event.AuthorID != userID {
		return errors.New("you`re not the author of this event")
	}

	return s.repo.DeleteEvent(ctx, id)
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
