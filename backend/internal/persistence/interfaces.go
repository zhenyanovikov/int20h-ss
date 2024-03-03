package persistence

import (
	"context"

	"github.com/google/uuid"
	"oss-backend/internal/models"
)

type Repo interface {
	Auth() Auth
	User() User
	Faculty() Faculty
	Group() Group
	Subject() Subject
	Notification() Notification
	Activity() Activity
	Assignment() Assignment
}

type Auth interface {
	CreateCredentials(ctx context.Context, credentials *models.UserCredentials) error
}

type User interface {
	GetUserByID(ctx context.Context, id uuid.UUID) (*models.User, error)
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
	ListTeachers(ctx context.Context) ([]models.Teacher, error)
	ListStudents(ctx context.Context, opts ...QueryBuilder) ([]models.Student, error)
	ListStudentsByGroupID(ctx context.Context, groupID uuid.UUID) ([]models.Student, error)
	ListStudentsByFacultyID(ctx context.Context, facultyID uuid.UUID) ([]models.Student, error)

	CreateUser(ctx context.Context, user *models.User) error
	CreateTeacher(ctx context.Context, teacher *models.Teacher) error
	CreateStudent(ctx context.Context, student *models.Student) error

	UpdateUser(ctx context.Context, user *models.User) error
}

type Group interface {
	ListGroups(ctx context.Context) ([]*models.Group, error)
	ListGroupsByFacultyID(ctx context.Context, facultyID uuid.UUID) ([]*models.Group, error)
	UpdateGroup(ctx context.Context, group *models.Group) error
	CreateGroup(ctx context.Context, group *models.Group) error
	DeleteGroup(ctx context.Context, groupID uuid.UUID) error
}

type Faculty interface {
	ListFaculties(ctx context.Context) ([]*models.Faculty, error)
	GetFacultyByID(ctx context.Context, facultyID uuid.UUID) (*models.Faculty, error)
	CreateFaculty(ctx context.Context, faculty *models.Faculty) error
	UpdateFaculty(ctx context.Context, faculty *models.Faculty) error
	DeleteFaculty(ctx context.Context, facultyID uuid.UUID) error
}

type Subject interface {
	ListSubjects(ctx context.Context, opts *models.SubjectFilterDTO) ([]*models.Subject, error)
	CreateSubject(ctx context.Context, subject *models.Subject) error
	UpdateSubject(ctx context.Context, subject *models.Subject) error
	DeleteSubject(ctx context.Context, subjectID uuid.UUID) error
}

type Notification interface {
	GetTemplates(ctx context.Context) ([]models.NotificationTemplate, error)
	CreateTemplate(ctx context.Context, template *models.NotificationTemplate) error
	UpdateTemplate(ctx context.Context, template *models.NotificationTemplate) error
	DeleteTemplate(ctx context.Context, templateID uuid.UUID) error
}

type Activity interface {
	GetActivityByID(ctx context.Context, id uuid.UUID) (*models.Activity, error)
	ListActivitiesByStudent(ctx context.Context, id uuid.UUID) ([]models.Activity, error)
	CreateActivity(ctx context.Context, activity *models.Activity) error
	UpdateActivity(ctx context.Context, activity *models.Activity) error
	DeleteActivity(ctx context.Context, activityID uuid.UUID) error
}

type Assignment interface {
	GetAssignmentByID(ctx context.Context, assignmentID uuid.UUID) (*models.Assignment, error)
	// ListAssignmentsByStudent(ctx context.Context, id uuid.UUID) ([]models.Assignment, error)
	ListAssignmentsBySubject(ctx context.Context, subjectID uuid.UUID) ([]models.Assignment, error)
	CreateAssignment(ctx context.Context, assignment *models.Assignment) error
	UpdateAssignment(ctx context.Context, assignment *models.Assignment) error
	DeleteAssignment(ctx context.Context, assignmentID uuid.UUID) error
}
