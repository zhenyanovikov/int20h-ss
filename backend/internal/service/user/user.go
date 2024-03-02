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

func (s *Service) ListTeachers(ctx context.Context) ([]models.Teacher, error) {
	return s.repo.User().ListTeachers(ctx)
}

func (s *Service) ListStudents(ctx context.Context) ([]models.Student, error) {
	return s.repo.User().ListStudents(ctx)
}

func (s *Service) ListStudentsByGroupID(ctx context.Context, groupID uuid.UUID) ([]models.Student, error) {
	return s.repo.User().ListStudentsByGroupID(ctx, groupID)
}

func (s *Service) Update(ctx context.Context, user *models.User) error {
	return s.repo.User().Update(ctx, user)
}
