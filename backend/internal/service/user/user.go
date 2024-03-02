package user

import (
	"context"

	"github.com/google/uuid"
	"oss-backend/internal/models"
)

func (s *Service) GetUser(id uuid.UUID) (*models.User, error) {
	ctx := context.Background()

	return s.userRepo.GetByID(ctx, id)
}
