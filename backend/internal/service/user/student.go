package user

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"oss-backend/internal/models"
)

func (s *Service) InviteStudent(ctx context.Context, dto *models.InviteStudentDTO) error {
	user := &models.User{
		ID:    uuid.New(),
		Email: dto.Email,
		Role:  models.RoleStudent,
	}

	teacher := &models.Student{
		GroupID: dto.GroupID,
		UserID:  user.ID,
	}

	if err := s.repo.User().CreateUser(ctx, user); err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}

	if err := s.repo.User().CreateStudent(ctx, teacher); err != nil {
		return fmt.Errorf("failed to create teacher: %w", err)
	}

	return nil
}
