package user

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"oss-backend/internal/models"
)

func (s *Service) InviteTeacher(ctx context.Context, dto *models.InviteTeacherDTO) error {
	user := &models.User{
		ID:    uuid.New(),
		Email: dto.Email,
		Role:  models.RoleTeacher,
	}

	teacher := &models.Teacher{
		UserID: user.ID,
	}

	if err := s.repo.User().CreateUser(ctx, user); err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}

	if err := s.repo.User().CreateTeacher(ctx, teacher); err != nil {
		return fmt.Errorf("failed to create teacher: %w", err)
	}

	return nil
}
