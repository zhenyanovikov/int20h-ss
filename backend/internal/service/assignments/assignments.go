package assignments

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"oss-backend/internal/models"
)

func (s *Service) GetAssignmentByID(ctx context.Context, assignmentID uuid.UUID) (*models.Assignment, error) {
	assignment, err := s.repo.Assignment().GetAssignmentByID(ctx, assignmentID)
	if err != nil {
		return nil, fmt.Errorf("failed to get assignment: %v", err)
	}

	return assignment, nil
}

func (s *Service) ListAssignmentsBySubject(ctx context.Context, subjectID uuid.UUID) ([]models.Assignment, error) {
	assignments, err := s.repo.Assignment().ListAssignmentsBySubject(ctx, subjectID)
	if err != nil {
		return nil, fmt.Errorf("failed to list assignments: %v", err)
	}

	return assignments, nil
}

func (s *Service) CreateAssignment(ctx context.Context, assignment *models.Assignment, subjectID uuid.UUID) error {
	assignment.SubjectID = subjectID

	if err := s.repo.Assignment().CreateAssignment(ctx, assignment); err != nil {
		return fmt.Errorf("failed to create assignment: %v", err)
	}

	return nil
}

func (s *Service) SubmitAssignment(ctx context.Context, studentID uuid.UUID, submittedAssigment *models.SubmittedAssigment) error {
	submittedAssigment.StudentID = studentID

	if err := s.repo.Assignment().CreateSubmittedAssignment(ctx, submittedAssigment); err != nil {
		return fmt.Errorf("failed to submit assignment: %v", err)
	}

	// todo create event

	return nil

}

func (s *Service) UpdateAssignment(ctx context.Context, assignment *models.Assignment) error {
	if err := s.repo.Assignment().UpdateAssignment(ctx, assignment); err != nil {
		return fmt.Errorf("failed to update assignment: %v", err)
	}

	return nil
}

func (s *Service) DeleteAssignment(ctx context.Context, assignmentID uuid.UUID) error {
	if err := s.repo.Assignment().DeleteAssignment(ctx, assignmentID); err != nil {
		return fmt.Errorf("failed to delete assignment: %v", err)
	}

	return nil
}
