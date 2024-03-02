package httpserver

import (
	"net/http"

	"github.com/google/uuid"
)

func (s *HTTPServer) getMe(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("user_id").(uuid.UUID)

	user, err := s.userSrv.GetByID(r.Context(), userID)
	if err != nil {
		s.respondError(w, http.StatusInternalServerError, err)
		return
	}

	s.respond(w, http.StatusOK, user)
}
