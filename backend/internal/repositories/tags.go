package repositories

import (
	"context"
	"find_a_walk/internal/domain"

	"github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type TagRepository struct {
	db *pgxpool.Pool
}

func NewTagRepository(db *pgxpool.Pool) *TagRepository {
	return &TagRepository{db: db}
}

func (r *TagRepository) GetTags(ctx context.Context) ([]*domain.Tag, error) {
	query := squirrel.Select("*").From("tags")
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

func (r *EventRepository) GetTagsByEventID(ctx context.Context, eventID uuid.UUID) ([]*domain.Tag, error) {
	query := squirrel.Select("tags.*").
		From("event_tags").
		Join("tags ON tags.id = event_tags.tag_id").
		Where(squirrel.Eq{"event_tags.event_id": eventID}).
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

func ToTagsFromRows(rows pgx.Rows) ([]*domain.Tag, error) {
	var tags []*domain.Tag
	var err error
	for rows.Next() {
		tag := &domain.Tag{}
		err = rows.Scan(&tag.ID, &tag.Name)
		tags = append(tags, tag)
	}
	if err != nil {
		return nil, err
	}
	return tags, nil
}
