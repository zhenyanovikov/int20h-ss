package httpserver

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"oss-backend/internal/models"
)

func (s *HTTPServer) inviteTeacher(w http.ResponseWriter, r *http.Request) {
	var dto models.InviteTeacherDTO

	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		s.respondError(w, http.StatusBadRequest, err)
		return
	}

	if err := s.userSrv.InviteTeacher(r.Context(), &dto); err != nil {
		s.respondError(w, http.StatusInternalServerError, err)
		return
	}

	if err := s.respond(w, http.StatusCreated, nil); err != nil {
		slog.Error("failed to respond: %v", err)
		return
	}
}
