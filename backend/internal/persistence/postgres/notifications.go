package postgres

import (
	"context"

	"github.com/google/uuid"
	"oss-backend/internal/models"
)

func (p *Postgres) GetTemplates(ctx context.Context) ([]models.NotificationTemplate, error) {
	var templates []models.NotificationTemplate

	if err := p.db.NewSelect().
		Model(&templates).
		Scan(ctx); err != nil {
		return nil, p.err(err)
	}

	return templates, nil
}

func (p *Postgres) CreateTemplate(ctx context.Context, template *models.NotificationTemplate) error {
	_, err := p.db.NewInsert().
		Model(template).
		Returning("*").
		Exec(ctx)
	if err != nil {
		return p.err(err)
	}

	return nil
}

func (p *Postgres) UpdateTemplate(ctx context.Context, template *models.NotificationTemplate) error {
	_, err := p.db.NewUpdate().
		Model(template).
		WherePK().
		Exec(ctx)
	if err != nil {
		return p.err(err)
	}

	return nil
}

func (p *Postgres) DeleteTemplate(ctx context.Context, templateID uuid.UUID) error {
	_, err := p.db.NewDelete().
		Model(&models.NotificationTemplate{}).
		Where("id = ?", templateID).
		Exec(ctx)
	if err != nil {
		return p.err(err)
	}

	return nil
}
