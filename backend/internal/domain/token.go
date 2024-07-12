package domain

import (
	"net/http"
)

type Token struct {
	Token string `json:"token"`
}

func (e *Token) Render(w http.ResponseWriter, r *http.Request) error {
	cookie := http.Cookie{
		Name:     "jwt",
		Value:    e.Token,
		Path:     "/",
		MaxAge:   3600,
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
	}
	http.SetCookie(w, &cookie)
	return nil
}
