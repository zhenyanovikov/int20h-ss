package httpserver

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
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

	user, err := s.userSrv.GetByEmail(r.Context(), gUser.Email)
	if err != nil {
		s.respondError(w, http.StatusInternalServerError, fmt.Errorf("get user by email: %w", err))
		return
	}

	if user == nil {
		s.respondError(w, http.StatusUnauthorized, fmt.Errorf("unauthorized"))
		return
	}

	if err = s.userSrv.UpdateAvatar(r.Context(), user.ID, gUser.Picture); err != nil {
		s.respondError(w, http.StatusInternalServerError, fmt.Errorf("update avatar: %w", err))
		return
	}

	res, err := s.authSrv.GenerateToken(context.Background(), user)
	if err != nil {
		s.respondError(w, http.StatusInternalServerError, fmt.Errorf("login: %w", err))
		return
	}

	if err = s.respond(w, http.StatusOK, res); err != nil {
		slog.Error(fmt.Sprintf("respond: %s", err))
		return
	}

	return
}
