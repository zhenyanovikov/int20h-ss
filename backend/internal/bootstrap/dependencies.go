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

	authRepo persistence.Auth
	userRepo persistence.User
}

func NewDependencies(config config.Config, httpServer *httpserver.HTTPServer,
	authSrv service.Auth, userSrv service.User,
	authRepo persistence.Auth, userRepo persistence.User) *Dependencies {
	return &Dependencies{
		Config:      config,
		HTTPServer:  httpServer,
		AuthService: authSrv,
		UserService: userSrv,
		authRepo:    authRepo,
		userRepo:    userRepo,
	}
}
