package models

import (
	"time"

	"github.com/google/uuid"
)

type Teacher struct {
	ID     uuid.UUID `json:"id" bun:",pk,nullzero"`
	UserID uuid.UUID `json:"-"`
	User   *User     `json:"user" bun:"rel:belongs-to"`

	CreatedAt time.Time `json:"createdAt" bun:",nullzero"`
}

type InviteTeacherDTO struct {
	Email string `json:"email"`
}
