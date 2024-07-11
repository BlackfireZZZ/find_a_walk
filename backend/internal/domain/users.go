package domain

import (
	"errors"
	"net/http"
	"regexp"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"-"`
	Interests []*Tag    `json:"interests,omitempty"`
}

type UserIn struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserAuth struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserInterestIn struct {
	Interests []string `json:"interests"`
}

func (a *UserInterestIn) Bind(r *http.Request) error {
	return nil
}

func NewUser(name, password, email string) User {
	return User{
		ID:        uuid.New(),
		Name:      name,
		Email:     email,
		Password:  password,
		Interests: []*Tag{},
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

func (a *UserAuth) Bind(r *http.Request) error {
	return nil
}

func (e *User) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
