//go:build wireinject
// +build wireinject

package bootstrap

import (
	"github.com/google/wire"
	"oss-backend/internal/config"
	"oss-backend/internal/persistence"
	"oss-backend/internal/persistence/postgres"
	"oss-backend/internal/service"
	"oss-backend/internal/service/activities"
	"oss-backend/internal/service/assignments"
	"oss-backend/internal/service/auth"
	"oss-backend/internal/service/aws/media"
	"oss-backend/internal/service/event"
	"oss-backend/internal/service/faculty"
	"oss-backend/internal/service/group"
	"oss-backend/internal/service/httpserver"
	"oss-backend/internal/service/notifier"
	"oss-backend/internal/service/subject"
	"oss-backend/internal/service/user"
)

func Up() (*Dependencies, error) {
	wire.Build(
		wire.Bind(new(service.User), new(*user.Service)),
		wire.Bind(new(service.Auth), new(*auth.Service)),
		wire.Bind(new(service.Media), new(*media.Service)),
		wire.Bind(new(service.Faculty), new(*faculty.Service)),
		wire.Bind(new(service.Notifier), new(*notifier.Service)),
		wire.Bind(new(service.Group), new(*group.Service)),
		wire.Bind(new(service.Activity), new(*activities.Service)),
		wire.Bind(new(service.Subject), new(*subject.Service)),
		wire.Bind(new(service.Event), new(*event.Service)),
		wire.Bind(new(service.Assignment), new(*assignments.Service)),

		wire.Bind(new(persistence.Repo), new(*postgres.Postgres)),

		config.New,
		httpserver.New,
		getPostgresConfig,

		postgres.New,
		auth.New,
		user.New,
		faculty.New,
		notifier.New,
		media.New,
		group.New,
		activities.New,
		subject.New,
		event.New,
		assignments.New,
		NewDependencies,
	)
	return &Dependencies{}, nil
}

func getPostgresConfig(cfg config.Config) config.Postgres {
	return cfg.Postgres
}
