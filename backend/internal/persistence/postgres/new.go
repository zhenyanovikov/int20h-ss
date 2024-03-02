package postgres

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/extra/bundebug"
	"oss-backend/internal/config"
)

type Postgres struct {
	db *bun.DB
}

func New(cfg config.Postgres) *Postgres {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Database)

	sqlDB := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))
	db := bun.NewDB(sqlDB, pgdialect.New())
	if cfg.Log {
		db.AddQueryHook(bundebug.NewQueryHook(bundebug.WithVerbose(true)))
	}

	return &Postgres{
		db: db,
	}
}

func (p *Postgres) err(err error) error {
	if errors.Is(err, sql.ErrNoRows) {
		return nil
	}
	return err
}
