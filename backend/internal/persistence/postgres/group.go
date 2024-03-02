package postgres

import (
	"context"
	"github.com/google/uuid"
	"oss-backend/internal/models"
)

func (p *Postgres) ListGroups(ctx context.Context) ([]*models.Group, error) {
	var groups []*models.Group

	err := p.db.NewSelect().
		Model(&groups).
		Scan(ctx)
	if err != nil {
		return nil, p.err(err)
	}

	return groups, nil
}

func (p *Postgres) ListGroupsByFacultyID(ctx context.Context, facultyID uuid.UUID) ([]*models.Group, error) {
	var groups []*models.Group

	err := p.db.NewSelect().
		Model(&groups).
		Where("faculty_id = ?", facultyID).
		Scan(ctx)
	if err != nil {
		return nil, p.err(err)
	}

	return groups, nil
}

func (p *Postgres) UpdateGroup(ctx context.Context, group *models.Group) error {
	_, err := p.db.NewUpdate().
		Model(group).
		WherePK().
		Exec(ctx)
	if err != nil {
		return p.err(err)
	}

	return nil
}

func (p *Postgres) CreateGroup(ctx context.Context, group *models.Group) error {
	_, err := p.db.NewInsert().
		Model(group).
		Returning("*").
		Exec(ctx)
	if err != nil {
		return p.err(err)
	}

	return nil
}

func (p *Postgres) DeleteGroup(ctx context.Context, groupID uuid.UUID) error {
	_, err := p.db.NewDelete().
		Model(&models.Group{}).
		Where("id = ?", groupID).
		Exec(ctx)
	if err != nil {
		return p.err(err)
	}

	return nil
}
