package models

import "github.com/google/uuid"

type EventType string

const (
	EventTypeLab      EventType = "Lab"
	EventTypeExam     EventType = "Exam"
	EventTypePractice EventType = "Practice"
	EventTypeLecture  EventType = "Lecture"
	EventTypeAbsent   EventType = "Absent"
)

type Event struct {
	ID   uuid.UUID
	Type EventType
	Mark int

	StudentID uuid.UUID
	SubjectID uuid.UUID
}

type FilterEventDTO struct {
	Type      *EventType `json:"type"`
	StudentID *uuid.UUID `json:"studentID"`
	SubjectID *uuid.UUID `json:"subjectID"`
}
