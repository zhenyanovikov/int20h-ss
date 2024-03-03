package models

import (
	"time"

	"github.com/google/uuid"
)

type Activity struct {
	ID        uuid.UUID `json:"id" bun:",pk,nullzero"`
	Name      string    `json:"name"`
	StudentID uuid.UUID `json:"-"`
	BonusMark int       `json:"bonusMark"`
	CreatedAt time.Time `json:"createdAt" bun:",nullzero"`
}
