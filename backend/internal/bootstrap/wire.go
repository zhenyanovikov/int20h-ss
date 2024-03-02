//go:build wireinject
// +build wireinject

package bootstrap

import (
	"github.com/google/wire"
	"oss-backend/internal/config"
	"oss-backend/internal/persistence"
	"oss-backend/internal/persistence/postgres"
	"oss-backend/internal/service"
	"oss-backend/internal/service/auth"
	"oss-backend/internal/service/aws/mail"
	"oss-backend/internal/service/aws/media"
	"oss-backend/internal/service/faculty"
	"oss-backend/internal/service/group"
	"oss-backend/internal/service/httpserver"
	"oss-backend/internal/service/user"
)

func Up() (*Dependencies, error) {
	wire.Build(
		wire.Bind(new(service.User), new(*user.Service)),
		wire.Bind(new(service.Auth), new(*auth.Service)),
		wire.Bind(new(service.Media), new(*media.Service)),
		wire.Bind(new(service.Faculty), new(*faculty.Service)),
		wire.Bind(new(service.Notifier), new(*mail.Service)),
		wire.Bind(new(service.Group), new(*group.Service)),

		wire.Bind(new(persistence.Repo), new(*postgres.Postgres)),
		//wire.Bind(new(persistence.Cache), new(*redis.Redis)),

		config.New,
		httpserver.New,
		getPostgresConfig,

		postgres.New,
		auth.New,
		user.New,
		faculty.New,
		mail.New,
		media.New,
		group.New,
		//redis.New,
		NewDependencies,
	)
	return &Dependencies{}, nil
}

func getPostgresConfig(cfg config.Config) config.Postgres {
	return cfg.Postgres
}
