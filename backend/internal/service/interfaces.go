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
	ListTeachers(ctx context.Context) ([]models.Teacher, error)
	ListStudents(ctx context.Context) ([]models.Student, error)
	ListStudentsByGroupID(ctx context.Context, groupID uuid.UUID) ([]models.Student, error)
	InviteTeacher(ctx context.Context, dto *models.InviteTeacherDTO) error
	InviteStudent(ctx context.Context, dto *models.InviteStudentDTO) error
	Update(ctx context.Context, user *models.User) error
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
