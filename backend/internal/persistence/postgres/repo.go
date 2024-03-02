package postgres

import "oss-backend/internal/persistence"

func (p *Postgres) Auth() persistence.Auth {
	return p
}

func (p *Postgres) User() persistence.User {
	return p
}

func (p *Postgres) Faculty() persistence.Faculty {
	return p
}

func (p *Postgres) Group() persistence.Group {
	return p
}

func (p *Postgres) Notification() persistence.Notification {
	return p
}
