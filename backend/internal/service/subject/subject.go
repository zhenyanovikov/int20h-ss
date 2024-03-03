package subject

import (
	"context"
	"github.com/google/uuid"
	"oss-backend/internal/models"
)

func (s *Service) ListSubjects(ctx context.Context, request *models.SubjectFilterDTO) ([]*models.Subject, error) {
	return s.repo.Subject().ListSubjects(ctx, request)
}

func (s *Service) CreateSubject(ctx context.Context, subject *models.Subject) error {
	return s.repo.Subject().CreateSubject(ctx, subject)
}

func (s *Service) UpdateSubject(ctx context.Context, subject *models.Subject) error {
	return s.repo.Subject().UpdateSubject(ctx, subject)
}

func (s *Service) DeleteSubject(ctx context.Context, id uuid.UUID) error {
	return s.repo.Subject().DeleteSubject(ctx, id)
}
