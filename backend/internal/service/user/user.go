package user

import (
	"context"
	"oss-backend/internal/persistence"

	"github.com/google/uuid"
	"oss-backend/internal/models"
)

func (s *Service) GetByID(ctx context.Context, userID uuid.UUID) (*models.User, error) {
	return s.repo.User().GetUserByID(ctx, userID)
}

func (s *Service) GetByEmail(ctx context.Context, email string) (*models.User, error) {
	return s.repo.User().GetUserByEmail(ctx, email)
}

func (s *Service) ListTeachers(ctx context.Context) ([]models.Teacher, error) {
	return s.repo.User().ListTeachers(ctx)
}

func (s *Service) ListStudents(ctx context.Context, filter models.FilterUserDTO) ([]models.Student, error) {
	opts := make([]persistence.QueryBuilder, 0, 2)
	if filter.Name != nil {
		opts = append(opts, persistence.SimilarName(*filter.Name))
	}

	if filter.GroupID != nil {
		opts = append(opts, persistence.Exact("group_id", *filter.GroupID))
	}

	return s.repo.User().ListStudents(ctx, opts...)
}

func (s *Service) ListStudentsByGroupID(ctx context.Context, groupID uuid.UUID) ([]models.Student, error) {
	return s.repo.User().ListStudentsByGroupID(ctx, groupID)
}

func (s *Service) Update(ctx context.Context, user *models.User) error {
	return s.repo.User().UpdateUser(ctx, user)
}
