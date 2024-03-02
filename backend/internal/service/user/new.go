package user

import "oss-backend/internal/persistence"

type Service struct {
	userRepo persistence.User
}

func New(userRepo persistence.User) *Service {
	return &Service{
		userRepo: userRepo,
	}
}
