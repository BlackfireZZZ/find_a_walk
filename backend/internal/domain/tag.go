package domain

import "net/http"

type Tag struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}


func (e *Tag) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}