package models

import (
	"time"

	"github.com/google/uuid"
)

type NotificationTemplate struct {
	ID        uuid.UUID `json:"id" bun:",pk,type:uuid,default:uuid_generate_v4()"`
	Subject   string    `json:"subject"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"created_at" bun:",nullzero"`
}

type SendNotificationDTO struct {
	Subject     string      `json:"subject"`
	Body        string      `json:"body"`
	StudentIDs  []uuid.UUID `json:"student_ids"`
	GroupID     uuid.UUID   `json:"group_id"`
	YearOfStudy int         `json:"year_of_study"`
	FacultyID   uuid.UUID   `json:"faculty_id"`
}
