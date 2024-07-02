package repositories

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"

	"find_a_walk/internal/domain"
)

// Реализация репозитория
type UserRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) GetUserByID(ctx context.Context, id int) (*domain.User, error) {
	event := &domain.User{}
	err := r.db.QueryRow(ctx, "SELECT id, name, email FROM users WHERE id = $1", id).
		Scan(&event.ID, &event.Name, &event.Email)
	if err != nil {
		return nil, err
	}
	return event, nil
}

func (r *UserRepository) CreateUser(ctx context.Context, event *domain.User) error {
	_, err := r.db.Exec(ctx, "INSERT INTO users (name, email) VALUES ($1, $2)", event.Name, event.Email)
	return err
}
