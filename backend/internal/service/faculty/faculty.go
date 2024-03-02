package faculty

import (
	"context"
	"github.com/google/uuid"
	"oss-backend/internal/models"
)

func (s *Service) GetFacultyByID(id uuid.UUID) (*models.Faculty, error) {
	ctx := context.Background()

	return s.facultyRepo.GetFacultyByID(ctx, id)
}

func (s *Service) ListFaculties() ([]*models.Faculty, error) {
	ctx := context.Background()

	return s.facultyRepo.ListFaculties(ctx)
}

func (s *Service) CreateFaculty(faculty *models.Faculty) error {
	ctx := context.Background()

	return s.facultyRepo.CreateFaculty(ctx, faculty)
}

func (s *Service) UpdateFaculty(faculty *models.Faculty) error {
	ctx := context.Background()

	return s.facultyRepo.UpdateFaculty(ctx, faculty)
}

func (s *Service) DeleteFaculty(id uuid.UUID) error {
	ctx := context.Background()

	return s.facultyRepo.DeleteFaculty(ctx, id)
}
