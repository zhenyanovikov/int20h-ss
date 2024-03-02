package postgres

import (
	"context"

	"oss-backend/internal/models"
)

func (p *Postgres) GetCredentialsByAccessToken(ctx context.Context, accessToken string) (*models.UserCredentials, error) {
	var credentials models.UserCredentials

	err := p.db.NewSelect().
		Model(&credentials).
		Where("access_token = ?", accessToken).
		Scan(ctx)
	if err != nil {
		return nil, p.err(err)
	}

	return &credentials, nil
}

func (p *Postgres) GetCredentialsByRefreshToken(ctx context.Context, refreshToken string) (*models.UserCredentials, error) {
	var credentials models.UserCredentials

	err := p.db.NewSelect().
		Model(&credentials).
		Where("refresh_token = ?", refreshToken).
		Scan(ctx)
	if err != nil {
		return nil, p.err(err)
	}

	return &credentials, nil
}

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
