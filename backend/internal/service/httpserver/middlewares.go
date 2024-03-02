package httpserver

import (
	"context"
	"fmt"
	"net/http"
	"strings"
)

func (s *HTTPServer) authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		t := strings.Split(authHeader, " ")

		if len(t) != 2 || t[0] != "Bearer" {
			s.respondError(w, http.StatusUnauthorized, fmt.Errorf("unauthorized"))
			return
		}

		authToken := t[1]

		user, err := s.authSrv.Login(r.Context(), authToken)
		if err != nil {
			s.respondError(w, http.StatusUnauthorized, fmt.Errorf("unauthorized"))
			return
		}

		r = r.WithContext(context.WithValue(r.Context(), "user", user))
		r = r.WithContext(context.WithValue(r.Context(), "user_id", user.ID))

		next.ServeHTTP(w, r)
		return
	})
}
