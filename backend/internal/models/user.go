package models

import "github.com/google/uuid"

type User struct {
	ID        uuid.UUID `json:"id" bun:",pk,nullzero"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Email     string    `json:"email"`
	AvatarURL string    `json:"avatarUrl"`
}
