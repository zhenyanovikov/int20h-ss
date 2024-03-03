package models

import "github.com/google/uuid"

type Activities struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	StudentID uuid.UUID `json:"-"`
}
