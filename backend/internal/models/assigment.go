package models

import (
	"time"

	"github.com/google/uuid"
)

type Assignment struct {
	ID uuid.UUID `json:"id" bun:",pk,nullzero"`

	SubjectID uuid.UUID `json:"-"`
	//Subject     *Subject  `json:"subject" bun:"rel:belongs-to"`

	Title         string    `json:"title"`
	Description   string    `json:"description"`
	AttachmentURL string    `json:"attachmentUrl,omitempty"`
	Deadline      time.Time `json:"deadline,omitempty"`

	CreatedAt time.Time `json:"createdAt" bun:",nullzero"`
}
