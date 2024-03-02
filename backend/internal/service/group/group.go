package group

import (
	"context"
	"github.com/google/uuid"
	"oss-backend/internal/models"
	"time"
)

func (s *Service) CreateGroup(ctx context.Context, groupModel *models.Group) error {
	return s.repo.Group().CreateGroup(ctx, groupModel)
}

func (s *Service) ListGroups(ctx context.Context) ([]*models.Group, error) {
	groups, err := s.repo.Group().ListGroups(ctx)
	if err != nil {
		return nil, err
	}

	return calculateYearsOfStudy(groups), nil
}

func (s *Service) ListGroupsByFacultyID(ctx context.Context, facultyID uuid.UUID) ([]*models.Group, error) {
	groups, err := s.repo.Group().ListGroupsByFacultyID(ctx, facultyID)
	if err != nil {
		return nil, err
	}

	return calculateYearsOfStudy(groups), nil
}

func calculateYearsOfStudy(groups []*models.Group) []*models.Group {
	additionalPoint := 1
	if time.Now().Month() < time.August {
		additionalPoint = 0
	}

	for i, group := range groups {
		groups[i].YearOfStudy = time.Now().Year() - group.YearStart + additionalPoint
	}

	return groups
}

func (s *Service) UpdateGroup(ctx context.Context, groupModel *models.Group) error {
	return s.repo.Group().UpdateGroup(ctx, groupModel)
}

func (s *Service) DeleteGroup(ctx context.Context, groupID uuid.UUID) error {
	return s.repo.Group().DeleteGroup(ctx, groupID)
}
