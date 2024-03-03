package httpserver

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"net/http"
	"oss-backend/internal/models"
)

func (s *HTTPServer) listSubjects(w http.ResponseWriter, r *http.Request) {
	vars := r.URL.Query()
	request := &models.SubjectFilterDTO{}
	if teacherID, ok := vars["teacher_id"]; ok {
		parsedUUID, err := uuid.Parse(teacherID[0])
		if err != nil {
			s.respondError(w, http.StatusBadRequest, err)
			return
		}

		request.TeacherID = &parsedUUID
	}

	if groupID, ok := vars["group_id"]; ok {
		parsedUUID, err := uuid.Parse(groupID[0])
		if err != nil {
			s.respondError(w, http.StatusBadRequest, err)
			return
		}

		request.GroupID = &parsedUUID
	}

	if name, ok := vars["name"]; ok {
		request.Name = &name[0]
	}

	subjects, err := s.subjectSrv.ListSubjects(r.Context(), request)
	if err != nil {
		s.respondError(w, http.StatusInternalServerError, err)
		return
	}

	s.respond(w, http.StatusOK, subjects)
}

func (s *HTTPServer) createSubject(w http.ResponseWriter, r *http.Request) {
	var subject models.Subject
	err := json.NewDecoder(r.Body).Decode(&subject)
	if err != nil {
		s.respondError(w, http.StatusBadRequest, err)
		return
	}

	if err := s.subjectSrv.CreateSubject(r.Context(), &subject); err != nil {
		s.respondError(w, http.StatusInternalServerError, err)
		return
	}

	s.respond(w, http.StatusCreated, subject)
}

func (s *HTTPServer) updateSubject(w http.ResponseWriter, r *http.Request) {
	var subject models.Subject
	err := json.NewDecoder(r.Body).Decode(&subject)
	if err != nil {
		s.respondError(w, http.StatusBadRequest, err)
		return
	}

	if err := s.subjectSrv.UpdateSubject(r.Context(), &subject); err != nil {
		s.respondError(w, http.StatusInternalServerError, err)
		return
	}

	s.respond(w, http.StatusOK, subject)
}

func (s *HTTPServer) deleteSubject(w http.ResponseWriter, r *http.Request) {
	subjectID, err := uuid.Parse(mux.Vars(r)["subject_id"])
	if err != nil {
		s.respondError(w, http.StatusBadRequest, err)
		return
	}

	if err := s.subjectSrv.DeleteSubject(r.Context(), subjectID); err != nil {
		s.respondError(w, http.StatusInternalServerError, err)
		return
	}

	s.respond(w, http.StatusNoContent, nil)
}
