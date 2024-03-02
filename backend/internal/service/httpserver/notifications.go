package httpserver

import (
	"net/http"

	"oss-backend/internal/models"
)

func (s *HTTPServer) sendNotification(w http.ResponseWriter, r *http.Request) {
	var sendNotificationRequest models.SendNotificationDTO

	if err := s.decodeJSON(r, &sendNotificationRequest); err != nil {
		s.respondError(w, http.StatusBadRequest, err)
		return
	}

	if err := s.notifierSrv.SendNotification(r.Context(), &sendNotificationRequest); err != nil {
		s.respondError(w, http.StatusInternalServerError, err)
		return
	}

	s.respond(w, http.StatusCreated, nil)

	return
}
