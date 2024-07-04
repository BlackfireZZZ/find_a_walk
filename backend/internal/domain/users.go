package domain

import (
	"errors"
	"net/http"
	"regexp"

	"github.com/google/uuid"
)

type User struct {
	ID    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Email string    `json:"email"`
}

type UserIn struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func NewUser(name string, email string) User {
	return User{
		ID:    uuid.New(),
		Name:  name,
		Email: email,
	}
}

func (a *UserIn) Bind(r *http.Request) error {
	if a.Name == "" || a.Email == "" {
		return errors.New("missing required field")
	}
	matchStatus, _ := regexp.MatchString(`^\w+([-+.']\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*$`, a.Email)
	if !matchStatus {
		return errors.New("email incorrect")
	}

	return nil
}

func (e *User) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
