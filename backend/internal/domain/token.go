package domain

import "net/http"


type Token struct {
	Token string `json:"token"`
}


func (e *Token) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}