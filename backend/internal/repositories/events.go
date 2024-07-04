package repositories

import (
	"context"
	"errors"
	"find_a_walk/internal/domain"
	"log"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type EventRepository struct {
	db *pgxpool.Pool
}

func NewEventRepository(db *pgxpool.Pool) *EventRepository {
	return &EventRepository{db: db}
}

func (r *EventRepository) CreateEvent(ctx context.Context, event *domain.EventIn) (*domain.Event, error) {
	eventSchema := domain.NewEvent(event.AuthorID, event.StartLongitude,
		event.StartLongitude, event.EndLatitude, event.EndLongitude,
		event.Date, event.Capacity)

	query_events := squirrel.Insert("events").
		Columns("id", "start_longitude",
			"start_latitude", "end_longitude",
			"end_latitude", "date", "capacity",
			"author_id").
		Values(eventSchema.ID, eventSchema.StartLongitude,
			eventSchema.StartLatitude, eventSchema.EndLongitude,
			eventSchema.EndLatitude, eventSchema.Date, eventSchema.Capacity,
			eventSchema.AuthorID).
		PlaceholderFormat(squirrel.Dollar)

	query_tags := squirrel.Insert("event_tags").
		Columns("event_id", "tag_id").
		PlaceholderFormat(squirrel.Dollar)

	for _, tag := range event.Tags {
		query_tags = query_tags.Values(eventSchema.ID, tag)
	}

	tx, err := r.db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return nil, err
	}
	defer tx.Rollback(ctx)

	events_stmt, events_args, err := query_events.ToSql()
	if err != nil {
		return nil, err
	}

	tags_stmt, tags_args, err := query_tags.ToSql()
	if err != nil {
		return nil, err
	}

	_, err = tx.Exec(ctx, events_stmt, events_args...)
	if err != nil {
		return nil, err
	}

	_, err = tx.Exec(ctx, tags_stmt, tags_args...)
	if err != nil {
		return nil, errors.New("one or several tags does`t exists")
	}

	if err = tx.Commit(ctx); err != nil {
		return nil, err
	}

	log.Println("Created event: ", &eventSchema.ID)
	return &eventSchema, nil
}

func (r *EventRepository) GetEvents(ctx context.Context, tags []string) ([]*domain.Event, error) {
	query := squirrel.
		Select("events.*", "count(members.event_id) as members_count").
		From("events").
		JoinClause("FULL JOIN members ON members.event_id = events.id").
		GroupBy("events.id").
		PlaceholderFormat(squirrel.Dollar)
	stmt, args, error := query.ToSql()
	log.Println(stmt)

	if error != nil {
		return nil, error
	}

	rows, err := r.db.Query(ctx, stmt, args...)
	if err != nil {
		return nil, error
	}

	var events []*domain.Event
	for rows.Next() {
		event := &domain.Event{}
		err = rows.
			Scan(&event.ID, &event.AuthorID, &event.StartLatitude,
				&event.StartLongitude, &event.EndLatitude,
				&event.EndLongitude, &event.Date, &event.Capacity,
				&event.MembersCount)
		if err != nil {
			return nil, err
		}
		event.Tags, err = r.GetTagsByEventID(ctx, event.ID)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}

	return events, nil
}

func (r *EventRepository) GetEventByID(ctx context.Context, id uuid.UUID) (*domain.Event, error) {
	event := &domain.Event{}

	query := squirrel.Select("events.*, count(members.event_id)").
		From("events").
		LeftJoin("members ON members.event_id = events.id").
		Where(squirrel.Eq{"events.id": id}).
		GroupBy("events.id").
		PlaceholderFormat(squirrel.Dollar)
	stmt, args, error := query.ToSql()

	if error != nil {
		return nil, error
	}

	err := r.db.QueryRow(ctx, stmt, args...).
		Scan(&event.ID, &event.AuthorID, &event.StartLatitude,
			&event.StartLongitude, &event.EndLatitude,
			&event.EndLongitude, &event.Date, &event.Capacity,
			&event.MembersCount)
	if err != nil {
		return nil, err
	}
	return event, nil
}

func (r *EventRepository) GetEventsByAnglesCoordinates(ctx context.Context, lon1, lat1, lon2, lat2 float64, tags []string) ([]*domain.Event, error) {
	result := []*domain.Event{}

	query := squirrel.
		Select("events.*", "count(members.event_id) as members_count").
		From("events").
		JoinClause("FULL JOIN members ON members.event_id = events.id").
		GroupBy("events.id").
		PlaceholderFormat(squirrel.Dollar)

	stmt, args, error := query.
		Where(squirrel.And{
			squirrel.GtOrEq{"start_longitude": lon1},
			squirrel.LtOrEq{"start_longitude": lon2},
			squirrel.GtOrEq{"start_latitude": lat1},
			squirrel.LtOrEq{"start_latitude": lat2},
		}).
		ToSql()

	if error != nil {
		return nil, error
	}

	rows, err := r.db.Query(ctx, stmt, args...)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		event := &domain.Event{}
		err = rows.
			Scan(&event.ID, &event.AuthorID, &event.StartLatitude,
				&event.StartLongitude, &event.EndLatitude,
				&event.EndLongitude, &event.Date, &event.Capacity,
				&event.MembersCount)
		if err != nil {
			return nil, err
		}
		event.Tags, err = r.GetTagsByEventID(ctx, event.ID)
		if err != nil {
			return nil, err
		}
		result = append(result, event)
	}
	return result, nil
}

func (r *EventRepository) DeleteExpiredEvents(ctx context.Context) error {
	query, args, err := squirrel.
		Delete("events").
		Where(squirrel.LtOrEq{"date": time.Now()}).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return err
	}

	_, err = r.db.Exec(ctx, query, args...)
	if err != nil {
		return err
	}
	return nil
}
