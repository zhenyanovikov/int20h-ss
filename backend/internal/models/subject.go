package models

import "github.com/google/uuid"

type Subject struct {
	ID        uuid.UUID `json:"id" bun:",pk,nullzero"`
	TeacherID uuid.UUID `json:"teacherID" bun:",nullzero"`
	GroupID   uuid.UUID `json:"groupID" bun:",nullzero"`
	Name      string    `json:"name"`
}

type SubjectFilterDTO struct {
	TeacherID *uuid.UUID `json:"teacherID"`
	GroupID   *uuid.UUID `json:"groupID"`
	Name      *string    `json:"name"`
}
