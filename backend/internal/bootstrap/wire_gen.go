// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package bootstrap

import (
	"oss-backend/internal/config"
	"oss-backend/internal/persistence/postgres"
	"oss-backend/internal/service/auth"
	"oss-backend/internal/service/httpserver"
	"oss-backend/internal/service/media"
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
	httpServer := httpserver.New(configConfig, service, userService, mediaService)
	dependencies := NewDependencies(configConfig, httpServer, service, userService, postgresPostgres)
	return dependencies, nil
}

// wire.go:

func getPostgresConfig(cfg config.Config) config.Postgres {
	return cfg.Postgres
}
