package postgres

import (
	"context"

	"oss-backend/internal/models"
)

func (p *Postgres) CreateCredentials(ctx context.Context, credentials *models.UserCredentials) error {
	_, err := p.db.NewInsert().
		Model(credentials).
		Returning("*").
		Exec(ctx)
	if err != nil {
		return p.err(err)
	}

	return nil
}
