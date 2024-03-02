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
	GetUserByID(ctx context.Context, id uuid.UUID) (*models.User, error)
	UpsertUserOnEmail(ctx context.Context, user *models.User) error
}

type Faculty interface {
	ListFaculties(ctx context.Context) ([]*models.Faculty, error)
	GetFacultyByID(ctx context.Context, facultyID uuid.UUID) (*models.Faculty, error)
	CreateFaculty(ctx context.Context, faculty *models.Faculty) error
	UpdateFaculty(ctx context.Context, faculty *models.Faculty) error
	DeleteFaculty(ctx context.Context, facultyID uuid.UUID) error
}
