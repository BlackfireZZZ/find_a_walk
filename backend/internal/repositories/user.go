package repositories

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"

	"find_a_walk/internal/domain"
)

// Реализация репозитория
type eventRepository struct {
	db *pgxpool.Pool
}

func NeweventRepository(db *pgxpool.Pool) *eventRepository {
	return &eventRepository{db: db}
}

func (r *eventRepository) GeteventByID(ctx context.Context, id int) (*domain.event, error) {
	event := &domain.event{}
	err := r.db.QueryRow(ctx, "SELECT id, name, email FROM events WHERE id = $1", id).
		Scan(&event.ID, &event.Name, &event.Email)
	if err != nil {
		return nil, err
	}
	return event, nil
}

func (r *eventRepository) Createevent(ctx context.Context, event *domain.event) error {
	_, err := r.db.Exec(ctx, "INSERT INTO events (name, email) VALUES ($1, $2)", event.Name, event.Email)
	return err
}
