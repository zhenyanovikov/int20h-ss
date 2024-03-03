package models

import (
	"time"

	"github.com/google/uuid"
)

type NotificationTemplate struct {
	ID        uuid.UUID `json:"id" bun:",pk,type:uuid,default:uuid_generate_v4()"`
	Subject   string    `json:"subject"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"createdAt" bun:",nullzero"`
}

type SendNotificationDTO struct {
	Subject     string      `json:"subject"`
	Body        string      `json:"body"`
	StudentIDs  []uuid.UUID `json:"studentIds"`
	GroupID     uuid.UUID   `json:"groupId"`
	YearOfStudy int         `json:"yearOfStudy"`
	FacultyID   uuid.UUID   `json:"facultyId"`
}
