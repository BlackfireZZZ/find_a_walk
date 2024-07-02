package domain

import (
	"time"

	"github.com/google/uuid"
)

type Event struct {
	ID             uuid.UUID `json:"id"`
	AuthorID       uuid.UUID `json:"authour_id"`
	StartLongitude float64   `json:"start_longitude"`
	StartLatitude  float64   `json:"start_latitude"`
	EndLongitude   float64   `json:"end_longitude"`
	EndLatitude    float64   `json:"end_latitude"`
	Date           time.Time `json:"date"`
	Capacity       int       `json:"capacity"`
	MembersCount   int       `json:"members_count"`
}
