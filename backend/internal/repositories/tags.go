package repositories

import (
	"context"
	"find_a_walk/internal/domain"
	"fmt"

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

func (r *TagRepository) GetTagByTagame(ctx context.Context, name string) (*domain.Tag, error) {
	query := squirrel.Select("*").From("tags").Where(squirrel.Eq{"name": name})
	stmt, args, err := query.ToSql()

	if err != nil {
		return nil, err
	}

	row := r.db.QueryRow(ctx, stmt, args...)

	tag := &domain.Tag{}
	err = row.Scan(&tag.ID, &tag.Name)
	if err != nil {
		return nil, err
	}

	return tag, nil
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

func TagsToString(tags []string) string {
	stringTags := "("
	for i, tag := range tags {
		if len(tags) == i+1 {
			stringTags += fmt.Sprintf("'%s'", tag)
		} else {
			stringTags += fmt.Sprintf("'%s',", tag)
		}
	}
	stringTags += ")"
	return stringTags
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
