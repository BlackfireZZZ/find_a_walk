package repositories

import (
	"context"
	"log"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"

	"find_a_walk/internal/domain"
)

// Реализация репозитория
type UserRepository struct {
	db *pgxpool.Pool
}

// GetUserInterests implements services.UserRepository.
func (r *UserRepository) GetUserInterests(ctx context.Context, id int) ([]*domain.Interest, error) {
	panic("unimplemented")
}

func NewUserRepository(db *pgxpool.Pool) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) GetUserByID(ctx context.Context, id uuid.UUID) (*domain.User, error) {
	user := &domain.User{}
	err := r.db.QueryRow(ctx, "SELECT id, name, email FROM users WHERE id = $1", id).
		Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) CreateUser(ctx context.Context, user *domain.User) error {
	user.ID = uuid.New()
	_, err := r.db.Exec(ctx, "INSERT INTO users (id, name, email) VALUES ($1, $2, $3)", user.ID, user.Name, user.Email)
	log.Println("Created new user: ", user)
	return err
}
