package service

import (
	"context"
	"io"

	"github.com/google/uuid"
	"oss-backend/internal/models"
)

type Auth interface {
	GenerateToken(ctx context.Context, user *models.User) (*models.UserCredentials, error)
	Login(ctx context.Context, accessToken string) (*models.User, error)
}

type User interface {
	GetByID(ctx context.Context, userID uuid.UUID) (*models.User, error)
	GetByEmail(ctx context.Context, email string) (*models.User, error)
	GetTeachers(ctx context.Context) ([]models.Teacher, error)
	InviteTeacher(ctx context.Context, dto *models.InviteTeacherDTO) error
	Update(ctx context.Context, user *models.User) error
}

type Media interface {
	Upload(ctx context.Context, fileReader io.Reader, extension string) (url string, err error)
}
