package user

import "oss-backend/internal/persistence"

type Service struct {
	repo persistence.Repo
}

func New(repo persistence.Repo) *Service {
	return &Service{
		repo: repo,
	}
}
