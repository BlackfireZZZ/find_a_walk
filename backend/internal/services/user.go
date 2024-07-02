package services

import (
	"context"

	"hse_school/internal/domain"
)

type UserRepository interface {
	GetUserByID(ctx context.Context, id int) (*domain.User, error)
	CreateUser(ctx context.Context, user *domain.User) error
}

// Реализация сервиса
type UserService struct {
	repo UserRepository
}

func NewDefaultUserService(repo UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetUserByID(ctx context.Context, id int) (*domain.User, error) {
	return s.repo.GetUserByID(ctx, id)
}

func (s *UserService) CreateUser(ctx context.Context, user *domain.User) error {
	return s.repo.CreateUser(ctx, user)
}
