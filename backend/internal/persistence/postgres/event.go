package postgres

import (
	"context"
	"github.com/google/uuid"
	"oss-backend/internal/models"
	"oss-backend/internal/persistence"
)

func (p *Postgres) ListEvents(ctx context.Context, opts ...persistence.QueryBuilder) ([]*models.Event, error) {
	var events []*models.Event

	err := p.db.NewSelect().
		Model(&events).
		ApplyQueryBuilder(p.apply(opts)).
		Scan(ctx)
	if err != nil {
		return nil, p.err(err)
	}

	return events, nil
}

func (p *Postgres) CreateEvent(ctx context.Context, event *models.Event) error {
	_, err := p.db.NewInsert().
		Model(event).
		Exec(ctx)
	if err != nil {
		return p.err(err)
	}

	return nil
}

func (p *Postgres) UpdateEvent(ctx context.Context, event *models.Event) error {
	_, err := p.db.NewUpdate().
		Model(event).
		WherePK().
		Exec(ctx)
	if err != nil {
		return p.err(err)
	}

	return nil
}

func (p *Postgres) DeleteEvent(ctx context.Context, eventID uuid.UUID) error {
	_, err := p.db.NewDelete().
		Model(&models.Event{}).
		Where("id = ?", eventID).
		Exec(ctx)
	if err != nil {
		return p.err(err)
	}

	return nil
}
