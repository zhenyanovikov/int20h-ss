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

type SubmittedAssigment struct {
	ID            uuid.UUID `json:"id" bun:",pk,nullzero"`
	AssignmentID  uuid.UUID `json:"assignmentId"`
	StudentID     uuid.UUID `json:"studentId"`
	Text          string    `json:"text"`
	AttachmentURL string    `json:"attachment"`
	//EventID	   uuid.UUID `json:"eventId"`
	//Event        *Event    `json:"event" bun:"rel:belongs-to"`
	Grade     *int      `json:"grade" bun:",scanonly"`
	CreatedAt time.Time `json:"createdAt" bun:",nullzero"`
}
