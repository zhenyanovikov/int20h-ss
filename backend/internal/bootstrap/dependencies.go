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

	AuthService     service.Auth
	UserService     service.User
	NotifierService service.Notifier
	ActivityService service.Activity

	repo persistence.Repo
}

func NewDependencies(config config.Config, httpServer *httpserver.HTTPServer,
	authSrv service.Auth, userSrv service.User,
	repo persistence.Repo, notifierSrv service.Notifier,
	activitySrv service.Activity) *Dependencies {
	return &Dependencies{
		Config:          config,
		HTTPServer:      httpServer,
		AuthService:     authSrv,
		UserService:     userSrv,
		NotifierService: notifierSrv,
		ActivityService: activitySrv,
		repo:            repo,
	}
}
