package event

import (
	"context"
	"github.com/google/uuid"
	"oss-backend/internal/persistence"

	"oss-backend/internal/models"
)

func (s *Service) CreateEvent(ctx context.Context, event *models.Event) error {
	return s.repo.Event().CreateEvent(ctx, event)
}

func (s *Service) ListEvents(ctx context.Context, dto *models.FilterEventDTO) ([]*models.Event, error) {
	opts := make([]persistence.QueryBuilder, 0, 3)
	if dto != nil {
		if dto.StudentID != nil {
			opts = append(opts, persistence.Exact("student_id", *dto.StudentID))
		}
		if dto.SubjectID != nil {
			opts = append(opts, persistence.Exact("subject_id", *dto.SubjectID))
		}
		if dto.Type != nil {
			opts = append(opts, persistence.Exact("type", *dto.Type))
		}
	}

	return s.repo.Event().ListEvents(ctx)
}

func (s *Service) UpdateEvent(ctx context.Context, event *models.Event) error {
	return s.repo.Event().UpdateEvent(ctx, event)
}

func (s *Service) DeleteEvent(ctx context.Context, eventID uuid.UUID) error {
	return s.repo.Event().DeleteEvent(ctx, eventID)
}
