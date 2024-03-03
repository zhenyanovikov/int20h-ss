package postgres

import (
	"context"
	"errors"
	"oss-backend/internal/persistence"

	"github.com/google/uuid"
	"oss-backend/internal/models"
)

func (p *Postgres) ListSubjects(ctx context.Context, opts *models.SubjectFilterDTO) ([]*models.Subject, error) {
	var subjects []*models.Subject

	var optsParsed = make([]persistence.QueryBuilder, 0, 3)
	if opts.TeacherID != nil {
		optsParsed = append(optsParsed, persistence.Exact("teacher_id", *opts.TeacherID))
	}
	if opts.GroupID != nil {
		optsParsed = append(optsParsed, persistence.Exact("group_id", *opts.GroupID))
	}
	if opts.Name != nil {
		optsParsed = append(optsParsed, persistence.Like("name", *opts.Name))
	}

	err := p.db.NewSelect().
		Model(&subjects).
		ApplyQueryBuilder(p.apply(optsParsed)).
		Scan(ctx)
	if err != nil {
		return nil, p.err(err)
	}

	return subjects, nil
}

func (p *Postgres) CreateSubject(ctx context.Context, subject *models.Subject) error {
	list, err := p.ListSubjects(ctx, &models.SubjectFilterDTO{
		TeacherID: &subject.TeacherID,
		GroupID:   &subject.GroupID,
		Name:      &subject.Name,
	})
	if err != nil {
		return p.err(err)
	}
	if list != nil {
		return errors.New("this subjects is already created")
	}

	_, err = p.db.NewInsert().
		Model(subject).
		Exec(ctx)
	if err != nil {
		return p.err(err)
	}

	return nil
}

func (p *Postgres) UpdateSubject(ctx context.Context, subject *models.Subject) error {
	_, err := p.db.NewUpdate().
		Model(&models.Subject{}).
		Where("teacher_id = ? AND group_id = ?", subject.TeacherID, subject.GroupID).
		Set("name = ?", subject.Name).
		Exec(ctx)
	if err != nil {
		return p.err(err)
	}

	return nil
}

func (p *Postgres) DeleteSubject(ctx context.Context, subjectID uuid.UUID) error {
	_, err := p.db.NewDelete().
		Model(&models.Subject{}).
		Where("id = ?", subjectID).
		Exec(ctx)
	if err != nil {
		return p.err(err)
	}

	return nil
}
