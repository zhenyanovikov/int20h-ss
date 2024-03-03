package activities

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"oss-backend/internal/models"
)

func (s *Service) ListActivitiesByStudent(ctx context.Context, studentID uuid.UUID) ([]models.Activity, error) {
	return s.repo.Activity().ListActivitiesByStudent(ctx, studentID)
}

func (s *Service) CreateActivity(ctx context.Context, activity *models.Activity, studentID uuid.UUID) error {
	activity.StudentID = studentID

	return s.repo.Activity().CreateActivity(ctx, activity)
}

func (s *Service) UpdateActivity(ctx context.Context, newActivity *models.Activity, activityID uuid.UUID) error {
	activity, err := s.repo.Activity().GetActivityByID(ctx, activityID)
	if err != nil {
		return fmt.Errorf("failed to get activity: %w", err)
	}

	newActivity.ID = activityID
	newActivity.StudentID = activity.StudentID
	newActivity.CreatedAt = activity.CreatedAt

	return s.repo.Activity().UpdateActivity(ctx, newActivity)
}

func (s *Service) DeleteActivity(ctx context.Context, activityID uuid.UUID) error {
	return s.repo.Activity().DeleteActivity(ctx, activityID)
}
