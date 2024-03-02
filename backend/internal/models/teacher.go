package models

import (
	"time"

	"github.com/google/uuid"
)

type Teacher struct {
	ID     uuid.UUID `json:"id" bun:",pk,nullzero"`
	UserID uuid.UUID `json:"user_id"`
	User   *User     `json:"user" bun:"rel:belongs-to"`

	CreatedAt time.Time `json:"created_at" bun:",nullzero"`
}

type InviteTeacherDTO struct {
	Email     string `json:"email"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}
