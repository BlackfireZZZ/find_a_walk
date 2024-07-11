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

func (r *UserRepository) GetTagsByUserID(ctx context.Context, userID uuid.UUID) ([]*domain.Tag, error) {
	query := squirrel.Select("tags.*").
		From("user_tags").
		Join("tags ON tags.id = user_tags.tag_id").
		Where(squirrel.Eq{"user_tags.user_id": userID}).
		GroupBy("tags.id").
		PlaceholderFormat(squirrel.Dollar)
	stmt, args, err := query.ToSql()

	if err != nil {
		return nil, err
	}

	rows, err := r.db.Query(ctx, stmt, args...)
	if err != nil {
		return nil, err
	}

	tags, err := ToTagsFromRows(rows)
	if err != nil {
		return nil, err
	}
	return tags, nil
}

func (r *UserRepository) DeleteInterests(ctx context.Context, id uuid.UUID, interests []string) (error) {
	query_tags := squirrel.Delete("user_tags").
		Where(squirrel.And{
			squirrel.Eq{"user_id": id},
			squirrel.Eq{"tag_id": interests},
		}).
		PlaceholderFormat(squirrel.Dollar)

	stmt, args, err := query_tags.ToSql()
	if err != nil {
		return err
	}
	log.Println(stmt)
	_, err = r.db.Exec(ctx, stmt, args...)
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) CreateInterests(ctx context.Context, userID uuid.UUID, interests []string) (*domain.User, error) {
	query_tags := squirrel.Insert("user_tags").
		Columns("user_id", "tag_id").
		PlaceholderFormat(squirrel.Dollar)

	for _, tag := range interests {
		query_tags = query_tags.Values(userID, tag)
	}

	stmt, args, err := query_tags.ToSql()
	if err != nil {
		return nil, err
	}

	_, err = r.db.Exec(ctx, stmt, args...)
	if err != nil {
		return nil, err
	}

	return r.GetUserByID(ctx, userID)
}

func (r *UserRepository) GetUserByID(ctx context.Context, id uuid.UUID) (*domain.User, error) {
	user := &domain.User{}

	query := squirrel.Select("id", "name", "email", "password").From("users").
		Where(squirrel.Eq{"id": id}).
		PlaceholderFormat(squirrel.Dollar)
	stmt, args, error := query.ToSql()

	if error != nil {
		return nil, error
	}

	err := r.db.QueryRow(ctx, stmt, args...).
		Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}

	user.Interests, err = r.GetTagsByUserID(ctx, id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) IsUserExists(ctx context.Context, email string) (bool, error) {
	query := squirrel.Select("COUNT(1)").
		From("users").
		Where(squirrel.Eq{"email": email}).
		PlaceholderFormat(squirrel.Dollar)

	stmt, args, err := query.ToSql()
	if err != nil {
		return false, err
	}

	var result int
	err = r.db.QueryRow(ctx, stmt, args...).Scan(&result)
	if err != nil {
		return false, err
	}
	if result > 0 {
		return true, nil
	}
	return false, nil
}

func (r *UserRepository) GetUserByEmail(ctx context.Context, email string) (*domain.User, error) {
	user := &domain.User{}

	query := squirrel.Select("id", "name", "email", "password").From("users").
		Where(squirrel.Eq{"email": email}).
		PlaceholderFormat(squirrel.Dollar)
	stmt, args, error := query.ToSql()

	if error != nil {
		return nil, error
	}

	err := r.db.QueryRow(ctx, stmt, args...).
		Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}

	user.Interests, err = r.GetTagsByUserID(ctx, user.ID)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) CreateUser(ctx context.Context, user *domain.UserIn) (*domain.User, error) {
	userSchema := domain.NewUser(user.Name, user.Password, user.Email)
	query := squirrel.Insert("users").
		Columns("id", "name", "email", "password").
		Values(userSchema.ID, userSchema.Name, userSchema.Email, userSchema.Password).
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