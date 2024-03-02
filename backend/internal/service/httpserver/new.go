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

	authSrv    service.Auth
	userSrv    service.User
	facultySrv service.Faculty
	mediaSrv   service.Media
}

func New(cfg config.Config, authSrv service.Auth,
	userSrv service.User, mediaSrv service.Media, facultyServ service.Faculty) *HTTPServer {
	server := &HTTPServer{
		cfg:        cfg,
		authSrv:    authSrv,
		userSrv:    userSrv,
		mediaSrv:   mediaSrv,
		facultySrv: facultyServ,
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
