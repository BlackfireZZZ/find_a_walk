package domain

import (
	"errors"
	"net/http"
	"time"

	"github.com/google/uuid"
)

type Event struct {
	ID             uuid.UUID `json:"id"`
	AuthorID       uuid.UUID `json:"author_id"`
	StartLongitude float64   `json:"start_longitude"`
	StartLatitude  float64   `json:"start_latitude"`
	EndLongitude   float64   `json:"end_longitude"`
	EndLatitude    float64   `json:"end_latitude"`
	Date           time.Time `json:"date"`
	Capacity       int       `json:"capacity"`
	MembersCount   int       `json:"members_count"`
	Tags           []*Tag    `json:"tags"`
}

func NewEvent(authorID uuid.UUID, startLongitude float64,
	startLatitude float64, endLatitude float64,
	endLongitude float64, date time.Time,
	capacity int) Event {
	return Event{
		ID:             uuid.New(),
		AuthorID:       authorID,
		StartLongitude: startLongitude,
		StartLatitude:  startLatitude,
		EndLongitude:   endLongitude,
		EndLatitude:    endLatitude,
		Date:           date,
		Capacity:       capacity,
		MembersCount:   0,
	}
}

type EventIn struct {
	AuthorID       uuid.UUID `json:"author_id"`
	StartLongitude float64   `json:"start_longitude"`
	StartLatitude  float64   `json:"start_latitude"`
	EndLongitude   float64   `json:"end_longitude"`
	EndLatitude    float64   `json:"end_latitude"`
	Date           time.Time `json:"date"`
	Capacity       int       `json:"capacity"`
	Tags           []string  `json:"tags"`
}

func (a *EventIn) Bind(r *http.Request) error {
	if a.StartLatitude < 0 || a.StartLongitude < 0 {
		return errors.New("invalid coordinates")
	}
	if a.Date.Before(time.Now()) {
		return errors.New("date in the past")
	}
	if a.Capacity < 0 || a.Capacity > 999 {
		return errors.New("invalid capacity. it must be between 0 and 999")
	}
	return nil
}

func (e *Event) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
