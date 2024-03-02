package service

import (
	"context"
	"io"

	"github.com/google/uuid"
	"oss-backend/internal/models"
)

type Auth interface {
	GenerateToken(ctx context.Context, user *models.User) (*models.UserCredentials, error)
	GetCredentials(ctx context.Context, user *models.User) (*models.UserCredentials, error)
	Login(ctx context.Context, accessToken string) (*models.User, error)
}

type User interface {
	GetUser(id uuid.UUID) (*models.User, error)
}

type Faculty interface {
	GetFacultyByID(id uuid.UUID) (*models.Faculty, error)
	ListFaculties() ([]*models.Faculty, error)
	CreateFaculty(faculty *models.Faculty) error
	UpdateFaculty(faculty *models.Faculty) error
	DeleteFaculty(id uuid.UUID) error
}

type Media interface {
	Upload(ctx context.Context, fileReader io.Reader, extension string) (url string, err error)
}
