package user

import (
	"context"

	"github.com/google/uuid"
	"oss-backend/internal/models"
)

func (s *Service) GetByID(ctx context.Context, userID uuid.UUID) (*models.User, error) {
	return s.repo.User().GetByID(ctx, userID)
}

func (s *Service) GetByEmail(ctx context.Context, email string) (*models.User, error) {
	return s.repo.User().GetByEmail(ctx, email)
}

func (s *Service) UpdateAvatar(ctx context.Context, userID uuid.UUID, pictureURL string) error {
	return s.repo.User().UpdateAvatar(ctx, userID, pictureURL)
}
