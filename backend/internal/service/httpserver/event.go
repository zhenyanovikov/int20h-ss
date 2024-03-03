package httpserver

import (
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"log/slog"
	"net/http"
	"oss-backend/internal/models"
)

func (s *HTTPServer) createEvent(w http.ResponseWriter, r *http.Request) {
	var event models.Event
	if err := s.decodeJSON(r, &event); err != nil {
		s.respondError(w, http.StatusBadRequest, err)
		return
	}

	if err := s.eventSrv.CreateEvent(r.Context(), &event); err != nil {
		s.respondError(w, http.StatusInternalServerError, err)
		return
	}
}

func (s *HTTPServer) getEvents(w http.ResponseWriter, r *http.Request) {
	var event models.FilterEventDTO
	eventType := models.EventType(r.URL.Query().Get("type"))
	if eventType != "" {
		event.Type = &(eventType)
	}

	studentIDRaw := r.URL.Query().Get("studentID")
	if studentIDRaw != "" {
		studentID, err := uuid.Parse(studentIDRaw)
		if err == nil {
			s.respondError(w, http.StatusBadRequest, err)
			return
		}

		event.StudentID = &studentID
	}

	subjectIDRaw := r.URL.Query().Get("subjectID")
	if subjectIDRaw != "" {
		subjectID, err := uuid.Parse(subjectIDRaw)
		if err == nil {
			s.respondError(w, http.StatusBadRequest, err)
			return
		}

		event.SubjectID = &subjectID
	}

	events, err := s.eventSrv.ListEvents(r.Context(), &event)
	if err != nil {
		s.respondError(w, http.StatusInternalServerError, err)
		return
	}

	if err = s.respond(w, http.StatusOK, events); err != nil {
		slog.Error("failed to respond: %v", err)
		return
	}
}

func (s *HTTPServer) updateEvent(w http.ResponseWriter, r *http.Request) {
	var event models.Event
	if err := s.decodeJSON(r, &event); err != nil {
		s.respondError(w, http.StatusBadRequest, err)
		return
	}

	if err := s.eventSrv.UpdateEvent(r.Context(), &event); err != nil {
		s.respondError(w, http.StatusInternalServerError, err)
		return
	}
}

func (s *HTTPServer) deleteEvent(w http.ResponseWriter, r *http.Request) {
	eventID, err := uuid.Parse(mux.Vars(r)["event_id"])
	if err != nil {
		s.respondError(w, http.StatusBadRequest, err)
		return
	}

	if err := s.eventSrv.DeleteEvent(r.Context(), eventID); err != nil {
		s.respondError(w, http.StatusInternalServerError, err)
		return
	}
}
