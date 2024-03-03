package models

import "github.com/google/uuid"

type Group struct {
	ID        uuid.UUID `json:"id" bun:",pk,nullzero"`
	Name      string    `json:"name"`
	YearStart int       `json:"yearStart"`
	FacultyID uuid.UUID `json:"facultyID" bun:",nullzero"`
	Faculty   *Faculty  `json:"faculty" bun:"rel:belongs-to"`

	YearOfStudy int `json:"yearOfStudy" bun:",scanonly"`
}
