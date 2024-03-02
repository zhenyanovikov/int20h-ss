package postgres

import (
	"context"

	"github.com/google/uuid"
	"oss-backend/internal/models"
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

func (p *Postgres) UpdateAvatar(ctx context.Context, userID uuid.UUID, pictureURL string) error {
	_, err := p.db.NewUpdate().
		Model(&models.User{}).
		Set("picture_url", pictureURL).
		Where("id = ?", userID).
		Exec(ctx)
	if err != nil {
		return p.err(err)
	}

	return nil
}
