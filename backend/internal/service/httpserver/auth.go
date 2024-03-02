package httpserver

import (
	"context"
	"fmt"
	"net/http"

	"oss-backend/internal/models"
)

func (s *HTTPServer) oauthCallback(w http.ResponseWriter, r *http.Request) {
	exchange, err := s.googleOAuthCfg.Exchange(context.Background(), r.URL.Query().Get("code"))
	if err != nil {
		s.respondError(w, http.StatusInternalServerError, fmt.Errorf("exchange: %w", err))
		return
	}

	gUser, err := fetchGoogleUser(exchange.AccessToken)
	if err != nil {
		s.respondError(w, http.StatusInternalServerError, fmt.Errorf("complete oauth: %w", err))
		return
	}

	user := models.User{
		FirstName: gUser.FirstName,
		LastName:  gUser.LastName,
		Email:     gUser.Email,
		AvatarURL: gUser.Picture,
	}

	res, err := s.authSrv.GetCredentials(context.Background(), &user)
	if err != nil {
		s.respondError(w, http.StatusInternalServerError, fmt.Errorf("login: %w", err))
		return
	}

	s.respond(w, http.StatusOK, res)

	return
}
