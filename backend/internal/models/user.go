package models

import "github.com/google/uuid"

type Role string

const (
	RoleAdmin   Role = "admin"
	RoleTeacher Role = "teacher"
	RoleStudent Role = "student"
)

type User struct {
	ID        uuid.UUID `json:"id" bun:",pk,nullzero"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Email     string    `json:"email"`
	AvatarURL string    `json:"avatarUrl"`
	Role      Role      `json:"role" bun:",nullzero"`
}
