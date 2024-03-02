package models

import "github.com/google/uuid"

type Faculty struct {
	ID   uuid.UUID `json:"id" bun:",pk,nullzero"`
	Name string    `json:"name"`
}
