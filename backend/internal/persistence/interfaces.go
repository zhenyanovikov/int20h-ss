package persistence

import (
	"context"

	"github.com/google/uuid"
	"oss-backend/internal/models"
)

type Repo interface {
	Auth() Auth
	User() User
}

type Auth interface {
	CreateCredentials(ctx context.Context, credentials *models.UserCredentials) error
}

type User interface {
	GetByID(ctx context.Context, id uuid.UUID) (*models.User, error)
	GetByEmail(ctx context.Context, email string) (*models.User, error)
	GetByRole(ctx context.Context, role models.Role) ([]models.Teacher, error)

	CreateUser(ctx context.Context, user *models.User) error
	CreateTeacher(ctx context.Context, teacher *models.Teacher) error

	Update(ctx context.Context, user *models.User) error
}
