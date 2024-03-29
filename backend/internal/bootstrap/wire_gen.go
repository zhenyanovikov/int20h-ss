// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package bootstrap

import (
	"oss-backend/internal/config"
	"oss-backend/internal/persistence/postgres"
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

// Injectors from wire.go:

func Up() (*Dependencies, error) {
	configConfig, err := config.New()
	if err != nil {
		return nil, err
	}
	configPostgres := getPostgresConfig(configConfig)
	postgresPostgres := postgres.New(configPostgres)
	service := auth.New(configConfig, postgresPostgres)
	userService := user.New(postgresPostgres)
	mediaService, err := media.New(configConfig)
	if err != nil {
		return nil, err
	}
	facultyService := faculty.New(postgresPostgres)
	groupService := group.New(postgresPostgres)
	notifierService, err := notifier.New(configConfig, postgresPostgres)
	if err != nil {
		return nil, err
	}
	subjectService := subject.New(postgresPostgres)
	activitiesService := activities.New(postgresPostgres)
	eventService := event.New(postgresPostgres)
	assignmentsService := assignments.New(postgresPostgres)
	httpServer := httpserver.New(configConfig, service, userService, mediaService, facultyService, groupService, notifierService, subjectService, activitiesService, eventService, assignmentsService)
	dependencies := NewDependencies(configConfig, httpServer, service, userService, postgresPostgres, notifierService, activitiesService, assignmentsService)
	return dependencies, nil
}

// wire.go:

func getPostgresConfig(cfg config.Config) config.Postgres {
	return cfg.Postgres
}
