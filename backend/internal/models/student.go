package models

import (
	"time"

	"github.com/google/uuid"
)

type Student struct {
	ID     uuid.UUID `json:"id" bun:",pk,nullzero"`
	UserID uuid.UUID `json:"-"`
	User   *User     `json:"user" bun:"rel:belongs-to"`

	GroupID   uuid.UUID `json:"-" bun:",nullzero"`
	CreatedAt time.Time `json:"createdAt" bun:",nullzero"`
}

type InviteStudentDTO struct {
	Email   string    `json:"email"`
	GroupID uuid.UUID `json:"groupId"`
}
