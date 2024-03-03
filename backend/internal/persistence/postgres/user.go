package postgres

import (
	"context"

	"github.com/google/uuid"
	"oss-backend/internal/models"
	"oss-backend/internal/persistence"
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

func (p *Postgres) GetByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User

	err := p.db.NewSelect().
		Model(&user).
		Where("email = ?", email).
		Scan(ctx)
	if err != nil {
		return nil, p.err(err)
	}

	return &user, nil
}

func (p *Postgres) ListTeachers(ctx context.Context) ([]models.Teacher, error) {
	var teachers []models.Teacher

	err := p.db.NewSelect().
		Model(&teachers).
		Relation(RelationUser).
		Where("role = ?", models.RoleTeacher).
		Scan(ctx)
	if err != nil {
		return nil, p.err(err)
	}

	return teachers, nil
}

func (p *Postgres) ListStudents(ctx context.Context, opts ...persistence.QueryBuilder) ([]models.Student, error) {
	var students []models.Student

	err := p.db.NewSelect().
		Model(&students).
		Relation(RelationUser).
		Where("role = ?", models.RoleStudent).
		ApplyQueryBuilder(p.apply(opts)).
		Scan(ctx)
	if err != nil {
		return nil, p.err(err)
	}

	return students, nil
}

func (p *Postgres) ListStudentsByGroupID(ctx context.Context, groupID uuid.UUID) ([]models.Student, error) {
	var students []models.Student

	err := p.db.NewSelect().
		Model(&students).
		Relation(RelationUser).
		Where("group_id = ?", groupID).
		Scan(ctx)
	if err != nil {
		return nil, p.err(err)
	}

	return students, nil
}

func (p *Postgres) ListStudentsByFacultyID(ctx context.Context, facultyID uuid.UUID) ([]models.Student, error) {
	var students []models.Student

	err := p.db.NewSelect().
		Model(&students).
		Relation(RelationUser).
		Join("JOIN groups ON groups.id = student.group_id").
		Where("groups.faculty_id = ?", facultyID).
		Scan(ctx)
	if err != nil {
		return nil, p.err(err)
	}

	return students, nil

}

func (p *Postgres) CreateUser(ctx context.Context, user *models.User) error {
	_, err := p.db.NewInsert().
		Model(user).
		Returning("*").
		Exec(ctx)
	if err != nil {
		return p.err(err)
	}

	return nil
}

func (p *Postgres) CreateTeacher(ctx context.Context, teacher *models.Teacher) error {
	_, err := p.db.NewInsert().
		Model(teacher).
		Returning("*").
		Exec(ctx)
	if err != nil {
		return p.err(err)
	}

	return nil
}

func (p *Postgres) CreateStudent(ctx context.Context, student *models.Student) error {
	_, err := p.db.NewInsert().
		Model(student).
		Returning("*").
		Exec(ctx)
	if err != nil {
		return p.err(err)
	}

	return nil
}

func (p *Postgres) Update(ctx context.Context, user *models.User) error {
	_, err := p.db.NewUpdate().
		Model(user).
		WherePK().
		Exec(ctx)
	if err != nil {
		return p.err(err)
	}

	return nil
}
