package models

import "github.com/google/uuid"

type SendNotificationDTO struct {
	Subject     string      `json:"subject"`
	Body        string      `json:"body"`
	StudentIDs  []uuid.UUID `json:"student_ids"`
	GroupID     uuid.UUID   `json:"group_id"`
	YearOfStudy int         `json:"year_of_study"`
	FacultyID   uuid.UUID   `json:"faculty_id"`
}
