package httpserver

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"oss-backend/internal/models"
)

func (s *HTTPServer) createActivity(w http.ResponseWriter, r *http.Request) {
	var activity models.Activity

	if err := s.decodeJSON(r, &activity); err != nil {
		s.respondError(w, http.StatusBadRequest, err)
		return
	}

	studentID, err := uuid.Parse(mux.Vars(r)["student_id"])
	if err != nil {
		s.respondError(w, http.StatusBadRequest, err)
		return
	}

	if err = s.activitySrv.CreateActivity(r.Context(), &activity, studentID); err != nil {
		s.respondError(w, http.StatusInternalServerError, err)
		return
	}

	s.respond(w, http.StatusCreated, &activity)
}

func (s *HTTPServer) getActivitiesByStudentID(w http.ResponseWriter, r *http.Request) {
	studentID, err := uuid.Parse(mux.Vars(r)["student_id"])
	if err != nil {
		s.respondError(w, http.StatusBadRequest, err)
		return
	}

	activities, err := s.activitySrv.ListActivitiesByStudent(r.Context(), studentID)
	if err != nil {
		s.respondError(w, http.StatusInternalServerError, err)
		return
	}

	s.respond(w, http.StatusOK, activities)
}

func (s *HTTPServer) updateActivity(w http.ResponseWriter, r *http.Request) {
	var activity models.Activity

	if err := s.decodeJSON(r, &activity); err != nil {
		s.respondError(w, http.StatusBadRequest, err)
		return
	}

	activityID, err := uuid.Parse(mux.Vars(r)["activity_id"])
	if err != nil {
		s.respondError(w, http.StatusBadRequest, err)
		return
	}

	if err = s.activitySrv.UpdateActivity(r.Context(), &activity, activityID); err != nil {
		s.respondError(w, http.StatusInternalServerError, err)
		return
	}

	s.respond(w, http.StatusOK, &activity)
}

func (s *HTTPServer) deleteActivity(w http.ResponseWriter, r *http.Request) {
	activityID, err := uuid.Parse(mux.Vars(r)["activity_id"])
	if err != nil {
		s.respondError(w, http.StatusBadRequest, err)
		return
	}

	if err = s.activitySrv.DeleteActivity(r.Context(), activityID); err != nil {
		s.respondError(w, http.StatusInternalServerError, err)
		return
	}

	s.respond(w, http.StatusNoContent, nil)
}
