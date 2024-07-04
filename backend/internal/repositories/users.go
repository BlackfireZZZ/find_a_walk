package repositories

import (
	"context"
	"log"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"

	"find_a_walk/internal/domain"

	"github.com/Masterminds/squirrel"
)

type UserRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) GetUserByID(ctx context.Context, id uuid.UUID) (*domain.User, error) {
	user := &domain.User{}

	query := squirrel.Select("id", "name", "email").From("users").
		Where(squirrel.Eq{"id": id}).
		PlaceholderFormat(squirrel.Dollar)
	stmt, args, error := query.ToSql()

	if error != nil {
		return nil, error
	}

	err := r.db.QueryRow(ctx, stmt, args...).
		Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) CreateUser(ctx context.Context, user *domain.UserIn) (*domain.User, error) {
	userSchema := domain.NewUser(user.Name, user.Email)

	query := squirrel.Insert("users").
		Columns("id", "name", "email").
		Values(userSchema.ID, userSchema.Name, userSchema.Email).
		PlaceholderFormat(squirrel.Dollar)

	stmt, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	_, err = r.db.Exec(ctx, stmt, args...)

	if err != nil {
		return nil, err
	}
	log.Println("Created user: ", &userSchema.ID)

	return &userSchema, nil
}
