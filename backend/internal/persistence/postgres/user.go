package postgres

import (
	"context"

	"github.com/google/uuid"
	"oss-backend/internal/models"
)

func (p *Postgres) GetByID(ctx context.Context, id uuid.UUID) (*models.User, error) {
	var user models.User

	err := p.db.NewSelect().
		Model(&user).
		Where("id = ?", id).
		Scan(ctx)
	if err != nil {
		return nil, p.err(err)
	}

	return &user, nil
}

func (p *Postgres) UpsertOnEmail(ctx context.Context, user *models.User) error {
	_, err := p.db.NewInsert().
		On("CONFLICT (email) DO UPDATE").
		Model(user).
		Returning("*").
		Exec(ctx)
	if err != nil {
		return p.err(err)
	}

	return nil
}
