package domain

import "github.com/google/uuid"

// Структуры данных
type User struct {
	ID    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Email string    `json:"email"`
}
