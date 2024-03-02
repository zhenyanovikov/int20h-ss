package bootstrap

import (
	"oss-backend/internal/config"
	"oss-backend/internal/persistence"
	"oss-backend/internal/service"
	"oss-backend/internal/service/httpserver"
)

type Dependencies struct {
	Config config.Config

	HTTPServer *httpserver.HTTPServer

	AuthService service.Auth
	UserService service.User

	repo persistence.Repo
}

func NewDependencies(config config.Config, httpServer *httpserver.HTTPServer,
	authSrv service.Auth, userSrv service.User,
	repo persistence.Repo) *Dependencies {
	return &Dependencies{
		Config:      config,
		HTTPServer:  httpServer,
		AuthService: authSrv,
		UserService: userSrv,
		repo:        repo,
	}
}
