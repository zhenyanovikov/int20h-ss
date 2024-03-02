package auth

import (
	"oss-backend/internal/config"
	"oss-backend/internal/persistence"
)

type Service struct {
	cfg config.Config

	authRepo    persistence.Auth
	userRepo    persistence.User
	facultyRepo persistence.Faculty
}

func New(cfg config.Config, authRepo persistence.Auth, userRepo persistence.User, facultyRepo persistence.Faculty) *Service {
	return &Service{
		cfg: cfg,

		authRepo:    authRepo,
		userRepo:    userRepo,
		facultyRepo: facultyRepo,
	}
}
