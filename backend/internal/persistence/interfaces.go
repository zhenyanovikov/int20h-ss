package persistence

import (
	"context"

	"github.com/google/uuid"
	"oss-backend/internal/models"
)

type Auth interface {
	GetCredentialsByAccessToken(ctx context.Context, accessToken string) (*models.UserCredentials, error)
	GetCredentialsByRefreshToken(ctx context.Context, refreshToken string) (*models.UserCredentials, error)

	CreateCredentials(ctx context.Context, credentials *models.UserCredentials) error
}

type User interface {
	GetByID(ctx context.Context, id uuid.UUID) (*models.User, error)
	UpsertOnEmail(ctx context.Context, user *models.User) error
}
