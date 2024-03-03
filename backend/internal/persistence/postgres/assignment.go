package postgres

import (
	"context"

	"github.com/google/uuid"
	"oss-backend/internal/models"
)

func (p *Postgres) GetAssignmentByID(ctx context.Context, id uuid.UUID) (*models.Assignment, error) {
	var assignment models.Assignment

	if err := p.db.NewSelect().
		Model(&assignment).
		Where("id = ?", id).
		Scan(ctx); err != nil {
		return nil, p.err(err)
	}
	return &assignment, nil
}

func (p *Postgres) ListAssignmentsBySubject(ctx context.Context, subjectID uuid.UUID) ([]models.Assignment, error) {
	var assignments []models.Assignment

	err := p.db.NewSelect().
		Model(&assignments).
		Where("subject_id = ?", subjectID).
		Scan(ctx)
	if err != nil {
		return nil, p.err(err)
	}

	return assignments, nil
}

func (p *Postgres) CreateAssignment(ctx context.Context, assignment *models.Assignment) error {
	if _, err := p.db.NewInsert().
		Model(assignment).
		Returning("*").
		Exec(ctx); err != nil {
		return p.err(err)
	}
	return nil
}

func (p *Postgres) CreateSubmittedAssignment(ctx context.Context, submittedAssigment *models.SubmittedAssigment) error {
	if _, err := p.db.NewInsert().
		Model(submittedAssigment).
		Returning("*").
		Exec(ctx); err != nil {
		return p.err(err)
	}

	return nil
}

func (p *Postgres) UpdateAssignment(ctx context.Context, assignment *models.Assignment) error {
	if _, err := p.db.NewUpdate().
		Model(assignment).
		Returning("*").
		WherePK().
		Exec(ctx); err != nil {
		return p.err(err)
	}
	return nil
}

func (p *Postgres) DeleteAssignment(ctx context.Context, id uuid.UUID) error {
	if _, err := p.db.NewDelete().
		Model(&models.Assignment{}).
		Where("id = ?", id).
		Exec(ctx); err != nil {
		return p.err(err)
	}
	return nil
}
