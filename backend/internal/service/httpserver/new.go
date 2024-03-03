package httpserver

import (
	"github.com/gorilla/mux"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"oss-backend/internal/config"
	"oss-backend/internal/service"
)

type HTTPServer struct {
	cfg            config.Config
	router         *mux.Router
	googleOAuthCfg *oauth2.Config

	authSrv     service.Auth
	userSrv     service.User
	subjectSrv  service.Subject
	groupSrv    service.Group
	facultySrv  service.Faculty
	mediaSrv    service.Media
	notifierSrv service.Notifier
}

func New(cfg config.Config, authSrv service.Auth,
	userSrv service.User, mediaSrv service.Media,
	facultySrv service.Faculty, groupSrv service.Group,
	notifierSrv service.Notifier, subjectSrv service.Subject) *HTTPServer {
	server := &HTTPServer{
		cfg:         cfg,
		authSrv:     authSrv,
		userSrv:     userSrv,
		mediaSrv:    mediaSrv,
		facultySrv:  facultySrv,
		groupSrv:    groupSrv,
		notifierSrv: notifierSrv,
		subjectSrv:  subjectSrv,
		googleOAuthCfg: &oauth2.Config{
			RedirectURL:  cfg.Oauth.Google.RedirectURL,
			ClientID:     cfg.Oauth.Google.ClientID,
			ClientSecret: cfg.Oauth.Google.ClientSecret,
			Scopes:       cfg.Oauth.Google.Scopes,
			Endpoint:     google.Endpoint,
		},
	}

	server.router = server.newRouter(cfg)

	return server
}
