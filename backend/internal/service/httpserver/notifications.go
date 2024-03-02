package httpserver

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
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

func (s *HTTPServer) getNotificationTemplates(w http.ResponseWriter, r *http.Request) {
	templates, err := s.notifierSrv.GetTemplates(r.Context())
	if err != nil {
		s.respondError(w, http.StatusInternalServerError, err)
		return
	}

	s.respond(w, http.StatusOK, templates)
}

func (s *HTTPServer) createNotificationTemplate(w http.ResponseWriter, r *http.Request) {
	var template models.NotificationTemplate

	if err := s.decodeJSON(r, &template); err != nil {
		s.respondError(w, http.StatusBadRequest, err)
		return
	}

	if err := s.notifierSrv.CreateTemplate(r.Context(), &template); err != nil {
		s.respondError(w, http.StatusInternalServerError, err)
		return
	}

	s.respond(w, http.StatusCreated, &template)
}

func (s *HTTPServer) updateNotificationTemplate(w http.ResponseWriter, r *http.Request) {
	var template models.NotificationTemplate

	if err := s.decodeJSON(r, &template); err != nil {
		s.respondError(w, http.StatusBadRequest, err)
		return
	}

	if err := s.notifierSrv.UpdateTemplate(r.Context(), &template); err != nil {
		s.respondError(w, http.StatusInternalServerError, err)
		return
	}

	s.respond(w, http.StatusOK, &template)
}

func (s *HTTPServer) deleteNotificationTemplates(w http.ResponseWriter, r *http.Request) {
	templateID, err := uuid.Parse(mux.Vars(r)["template_id"])
	if err != nil {
		s.respondError(w, http.StatusBadRequest, err)
		return
	}

	if err = s.notifierSrv.DeleteTemplate(r.Context(), templateID); err != nil {
		s.respondError(w, http.StatusInternalServerError, err)
		return
	}

	s.respond(w, http.StatusOK, nil)
}
