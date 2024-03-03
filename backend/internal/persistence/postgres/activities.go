package postgres

import (
	"context"

	"github.com/google/uuid"
	"oss-backend/internal/models"
)

func (p *Postgres) GetActivityByID(ctx context.Context, id uuid.UUID) (*models.Activity, error) {
	var activity models.Activity

	if err := p.db.NewSelect().
		Model(&activity).
		Where("id = ?", id).
		Scan(ctx); err != nil {
		return nil, p.err(err)
	}

	return &activity, nil
}

func (p *Postgres) ListActivitiesByStudent(ctx context.Context, id uuid.UUID) ([]models.Activity, error) {
	var activities []models.Activity

	if err := p.db.NewSelect().
		Model(&activities).
		Where("student_id = ?", id).
		Scan(ctx); err != nil {
		return nil, p.err(err)
	}

	return activities, nil
}

func (p *Postgres) CreateActivity(ctx context.Context, activity *models.Activity) error {
	if _, err := p.db.NewInsert().
		Model(activity).
		Returning("*").
		Exec(ctx); err != nil {
		return p.err(err)
	}

	return nil
}

func (p *Postgres) UpdateActivity(ctx context.Context, activity *models.Activity) error {
	if _, err := p.db.NewUpdate().
		Model(activity).
		Where("id = ?", activity.ID).
		Returning("*").
		Exec(ctx); err != nil {
		return p.err(err)
	}

	return nil
}

func (p *Postgres) DeleteActivity(ctx context.Context, activityID uuid.UUID) error {
	if _, err := p.db.NewDelete().
		Model((*models.Activity)(nil)).
		Where("id = ?", activityID).
		Exec(ctx); err != nil {
		return p.err(err)
	}

	return nil
}
