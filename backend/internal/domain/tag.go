package domain

type Tag struct {
	ID       int    `json:"id"`
	Event_id int    `json:"event_id"`
	Tag      string `json:"tag"`
}
