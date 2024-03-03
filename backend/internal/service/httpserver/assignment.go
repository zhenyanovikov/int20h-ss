package httpserver

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"oss-backend/internal/models"
)

func (s *HTTPServer) createAssignment(w http.ResponseWriter, r *http.Request) {
	var assignment models.Assignment

	if err := s.decodeJSON(r, &assignment); err != nil {
		s.respondError(w, http.StatusBadRequest, err)
		return
	}

	subjectID, err := uuid.Parse(mux.Vars(r)["subject_id"])
	if err != nil {
		s.respondError(w, http.StatusBadRequest, err)
		return
	}

	if err = s.assignmentSrv.CreateAssignment(r.Context(), &assignment, subjectID); err != nil {
		s.respondError(w, http.StatusInternalServerError, err)
		return
	}

	s.respond(w, http.StatusCreated, assignment)
}

func (s *HTTPServer) submitAssignment(w http.ResponseWriter, r *http.Request) {
	var submission models.SubmittedAssigment

	if err := s.decodeJSON(r, &submission); err != nil {
		s.respondError(w, http.StatusBadRequest, err)
		return
	}

	assignmentID, err := uuid.Parse(mux.Vars(r)["assignment_id"])
	if err != nil {
		s.respondError(w, http.StatusBadRequest, err)
		return
	}

	submission.AssignmentID = assignmentID

	userID := r.Context().Value("user_id").(uuid.UUID)

	if err = s.assignmentSrv.SubmitAssignment(r.Context(), userID, &submission); err != nil {
		s.respondError(w, http.StatusInternalServerError, err)
		return
	}

	s.respond(w, http.StatusCreated, submission)
}

func (s *HTTPServer) getAssignmentBySubjectID(w http.ResponseWriter, r *http.Request) {
	subjectID, err := uuid.Parse(mux.Vars(r)["subject_id"])
	if err != nil {
		s.respondError(w, http.StatusBadRequest, err)
		return
	}

	assignment, err := s.assignmentSrv.ListAssignmentsBySubject(r.Context(), subjectID)
	if err != nil {
		s.respondError(w, http.StatusInternalServerError, err)
		return
	}

	s.respond(w, http.StatusOK, assignment)
}

func (s *HTTPServer) updateAssignment(w http.ResponseWriter, r *http.Request) {
	var assignment models.Assignment

	if err := s.decodeJSON(r, &assignment); err != nil {
		s.respondError(w, http.StatusBadRequest, err)
		return
	}

	if err := s.assignmentSrv.UpdateAssignment(r.Context(), &assignment); err != nil {
		s.respondError(w, http.StatusInternalServerError, err)
		return
	}

	s.respond(w, http.StatusOK, assignment)
}

func (s *HTTPServer) deleteAssignment(w http.ResponseWriter, r *http.Request) {
	assignmentID, err := uuid.Parse(mux.Vars(r)["assignment_id"])
	if err != nil {
		s.respondError(w, http.StatusBadRequest, err)
		return
	}

	if err = s.assignmentSrv.DeleteAssignment(r.Context(), assignmentID); err != nil {
		s.respondError(w, http.StatusInternalServerError, err)
		return
	}

	s.respond(w, http.StatusNoContent, nil)
}
