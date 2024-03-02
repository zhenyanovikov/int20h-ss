package auth

import (
	"oss-backend/internal/config"
	"oss-backend/internal/persistence"
)

type Service struct {
	cfg config.Config

	repo persistence.Repo
}

func New(cfg config.Config, repo persistence.Repo) *Service {
	return &Service{
		cfg: cfg,

		repo: repo,
	}
}
