package postgres

import (
	"context"

	"github.com/google/uuid"
	"oss-backend/internal/models"
)

func (p *Postgres) GetFacultyByID(ctx context.Context, facultyID uuid.UUID) (*models.Faculty, error) {
	var faculty models.Faculty

	err := p.db.NewSelect().
		Model(&faculty).
		Where("id = ?", facultyID).
		Scan(ctx)
	if err != nil {
		return nil, p.err(err)
	}

	return &faculty, nil
}

func (p *Postgres) ListFaculties(ctx context.Context) ([]*models.Faculty, error) {
	var faculties []*models.Faculty

	err := p.db.NewSelect().
		Model(&faculties).
		Scan(ctx)
	if err != nil {
		return nil, p.err(err)
	}

	return faculties, nil
}

func (p *Postgres) CreateFaculty(ctx context.Context, faculty *models.Faculty) error {
	_, err := p.db.NewInsert().
		Model(faculty).
		Returning("*").
		Exec(ctx)
	if err != nil {
		return p.err(err)
	}

	return nil
}

func (p *Postgres) UpdateFaculty(ctx context.Context, faculty *models.Faculty) error {
	_, err := p.db.NewUpdate().
		Model(faculty).
		WherePK().
		Returning("*").
		Exec(ctx)
	if err != nil {
		return p.err(err)
	}

	return nil
}

func (p *Postgres) DeleteFaculty(ctx context.Context, facultyID uuid.UUID) error {
	_, err := p.db.NewDelete().
		Model(&models.Faculty{}).
		Where("id = ?", facultyID).
		Exec(ctx)
	if err != nil {
		return p.err(err)
	}

	return nil
}
