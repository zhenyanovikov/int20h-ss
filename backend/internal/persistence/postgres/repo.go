package postgres

import "oss-backend/internal/persistence"

func (p *Postgres) Auth() persistence.Auth {
	return p
}

func (p *Postgres) User() persistence.User {
	return p
}
