package httpserver

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"net/http"
	"oss-backend/internal/models"
)

func (s *HTTPServer) getFacultyByID(w http.ResponseWriter, r *http.Request) {
	facultyID, err := uuid.Parse(mux.Vars(r)["faculty_id"])
	if err != nil {
		s.respondError(w, http.StatusBadRequest, err)
		return
	}

	faculty, err := s.facultySrv.GetFacultyByID(facultyID)
	if err != nil {
		s.respondError(w, http.StatusNotFound, err)
		return
	}

	s.respond(w, http.StatusOK, faculty)
}

func (s *HTTPServer) listFaculties(w http.ResponseWriter, r *http.Request) {
	faculties, err := s.facultySrv.ListFaculties()
	if err != nil {
		s.respondError(w, http.StatusInternalServerError, err)
		return
	}

	s.respond(w, http.StatusOK, faculties)
}

func (s *HTTPServer) createFaculty(w http.ResponseWriter, r *http.Request) {
	var faculty models.Faculty
	err := json.NewDecoder(r.Body).Decode(&faculty)
	if err != nil {
		s.respondError(w, http.StatusBadRequest, err)
		return
	}

	err = s.facultySrv.CreateFaculty(&faculty)
	if err != nil {
		s.respondError(w, http.StatusInternalServerError, err)
		return
	}

	s.respond(w, http.StatusCreated, faculty)
}

func (s *HTTPServer) updateFaculty(w http.ResponseWriter, r *http.Request) {
	var faculty models.Faculty
	err := json.NewDecoder(r.Body).Decode(&faculty)
	if err != nil {
		s.respondError(w, http.StatusBadRequest, err)
		return
	}

	err = s.facultySrv.UpdateFaculty(&faculty)
	if err != nil {
		s.respondError(w, http.StatusInternalServerError, err)
		return
	}

	s.respond(w, http.StatusOK, faculty)
}

func (s *HTTPServer) deleteFaculty(w http.ResponseWriter, r *http.Request) {
	facultyID, err := uuid.Parse(mux.Vars(r)["faculty_id"])
	if err != nil {
		s.respondError(w, http.StatusBadRequest, err)
		return
	}

	err = s.facultySrv.DeleteFaculty(facultyID)
	if err != nil {
		s.respondError(w, http.StatusInternalServerError, err)
		return
	}

	s.respond(w, http.StatusNoContent, nil)
}
