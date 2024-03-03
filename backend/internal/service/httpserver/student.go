package httpserver

import (
	"log/slog"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"

	"oss-backend/internal/models"
)

func (s *HTTPServer) getStudents(w http.ResponseWriter, r *http.Request) {
	var request models.FilterUserDTO
	name := r.URL.Query().Get("name")
	if name != "" {
		request.Name = &name
	}

	groupID := r.URL.Query().Get("groupID")
	if groupID != "" {
		parsedID, err := uuid.Parse(groupID)
		if err != nil {
			s.respondError(w, http.StatusBadRequest, err)
			return
		}
		request.GroupID = &parsedID
	}

	students, err := s.userSrv.ListStudents(r.Context(), request)
	if err != nil {
		s.respondError(w, http.StatusInternalServerError, err)
		return
	}

	if students == nil {
		students = []models.Student{}
	}

	if err = s.respond(w, http.StatusOK, students); err != nil {
		slog.Error("failed to respond: %v", err)
		return
	}
}

func (s *HTTPServer) getStudentsByGroupID(w http.ResponseWriter, r *http.Request) {
	groupID, err := uuid.Parse(mux.Vars(r)["group_id"])
	if err != nil {
		s.respondError(w, http.StatusBadRequest, err)
		return
	}

	teachers, err := s.userSrv.ListStudentsByGroupID(r.Context(), groupID)
	if err != nil {
		s.respondError(w, http.StatusInternalServerError, err)
		return
	}

	if err = s.respond(w, http.StatusOK, teachers); err != nil {
		slog.Error("failed to respond: %v", err)
		return
	}
}

func (s *HTTPServer) inviteStudent(w http.ResponseWriter, r *http.Request) {
	var dto models.InviteStudentDTO

	if err := s.decodeJSON(r, &dto); err != nil {
		s.respondError(w, http.StatusBadRequest, err)
		return
	}

	if err := s.userSrv.InviteStudent(r.Context(), &dto); err != nil {
		s.respondError(w, http.StatusInternalServerError, err)
		return
	}

	if err := s.respond(w, http.StatusCreated, nil); err != nil {
		slog.Error("failed to respond: %v", err)
		return
	}
}
