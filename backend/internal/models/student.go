package models

import (
	"github.com/google/uuid"
	"time"
)

type Student struct {
	ID     uuid.UUID `json:"id" bun:",pk,nullzero"`
	UserID uuid.UUID `json:"-"`
	User   *User     `json:"user" bun:"rel:belongs-to"`

	GroupID   uuid.UUID `json:"group_id" bun:",nullzero"`
	CreatedAt time.Time `json:"created_at" bun:",nullzero"`
}

type InviteStudentDTO struct {
	Email   string    `json:"email"`
	GroupID uuid.UUID `json:"group_id"`
}
