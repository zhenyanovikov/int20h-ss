package httpserver

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"log/slog"
	"net/http"
	"oss-backend/internal/models"
)

func (s *HTTPServer) getGroups(w http.ResponseWriter, r *http.Request) {
	groups, err := s.groupSrv.ListGroups(r.Context())
	if err != nil {
		s.respondError(w, http.StatusInternalServerError, err)
		return
	}

	if err = s.respond(w, http.StatusOK, groups); err != nil {
		slog.Error("failed to respond: %v", err)
		return
	}
}

func (s *HTTPServer) getGroupsByFacultyID(w http.ResponseWriter, r *http.Request) {
	facultyID, err := uuid.Parse(mux.Vars(r)["faculty_id"])
	if err != nil {
		s.respondError(w, http.StatusBadRequest, err)
		return
	}

	groups, err := s.groupSrv.ListGroupsByFacultyID(r.Context(), facultyID)
	if err != nil {
		s.respondError(w, http.StatusInternalServerError, err)
		return
	}

	if err = s.respond(w, http.StatusOK, groups); err != nil {
		slog.Error("failed to respond: %v", err)
		return
	}
}

func (s *HTTPServer) updateGroup(w http.ResponseWriter, r *http.Request) {
	var group models.Group
	err := json.NewDecoder(r.Body).Decode(&group)
	if err != nil {
		s.respondError(w, http.StatusBadRequest, err)
		return
	}

	if err := s.groupSrv.UpdateGroup(r.Context(), &group); err != nil {
		s.respondError(w, http.StatusInternalServerError, err)
		return
	}
}

func (s *HTTPServer) createGroup(w http.ResponseWriter, r *http.Request) {
	var group models.Group
	err := json.NewDecoder(r.Body).Decode(&group)
	if err != nil {
		s.respondError(w, http.StatusBadRequest, err)
		return
	}

	if err := s.groupSrv.CreateGroup(r.Context(), &group); err != nil {
		s.respondError(w, http.StatusInternalServerError, err)
		return
	}
}

func (s *HTTPServer) deleteGroup(w http.ResponseWriter, r *http.Request) {
	groupID, err := uuid.Parse(mux.Vars(r)["group_id"])
	if err != nil {
		s.respondError(w, http.StatusBadRequest, err)
	}

	if err := s.groupSrv.DeleteGroup(r.Context(), groupID); err != nil {
		s.respondError(w, http.StatusInternalServerError, err)
		return
	}
}
