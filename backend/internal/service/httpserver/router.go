package httpserver

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/google"
	"oss-backend/internal/config"
	"oss-backend/internal/models"
)

func (s *HTTPServer) newRouter(_ config.Config) *mux.Router {
	var (
		router     = mux.NewRouter()
		api        = router.PathPrefix("/api/v1").Subrouter()
		authorized = api.PathPrefix("/").Subrouter()
		admin      = authorized.PathPrefix("/").Subrouter()
	)

	goth.UseProviders(
		google.New(s.googleOAuthCfg.ClientID, s.googleOAuthCfg.ClientSecret, s.googleOAuthCfg.RedirectURL, s.googleOAuthCfg.Scopes...),
	)

	router.Use(corsMiddleware)
	authorized.Use(s.authMiddleware)

	admin.Use(s.roleMiddleware(models.RoleAdmin))

	api.HandleFunc("/auth/{provider}/callback", s.oauthCallback).Methods(http.MethodGet, http.MethodOptions)

	api.HandleFunc("/status", s.getStatus).Methods(http.MethodGet, http.MethodOptions)
	authorized.HandleFunc("/profile/me", s.getMe).Methods(http.MethodGet, http.MethodOptions)

	admin.HandleFunc("/teacher/invite", s.inviteTeacher).Methods(http.MethodPost, http.MethodOptions)
	admin.HandleFunc("/teacher", s.getTeachers).Methods(http.MethodGet, http.MethodOptions)

	admin.HandleFunc("/student/invite", s.inviteStudent).Methods(http.MethodPost, http.MethodOptions)
	admin.HandleFunc("/student", s.getStudents).Methods(http.MethodGet, http.MethodOptions)

	authorized.HandleFunc("/group/{group_id}/students", s.getStudentsByGroupID).Methods(http.MethodGet, http.MethodOptions)
	authorized.HandleFunc("/group", s.getGroups).Methods(http.MethodGet, http.MethodOptions)

	admin.HandleFunc("/group", s.createGroup).Methods(http.MethodPost, http.MethodOptions)
	admin.HandleFunc("/group", s.updateGroup).Methods(http.MethodPut, http.MethodOptions)
	admin.HandleFunc("/group/{group_id}", s.deleteGroup).Methods(http.MethodDelete, http.MethodOptions)

	authorized.HandleFunc("/notification/send", s.sendNotification).Methods(http.MethodPost, http.MethodOptions)
	authorized.HandleFunc("/notification/email/template", s.getNotificationTemplates).Methods(http.MethodGet, http.MethodOptions)
	authorized.HandleFunc("/notification/email/template", s.createNotificationTemplate).Methods(http.MethodPost, http.MethodOptions)
	authorized.HandleFunc("/notification/email/template/{template_id}", s.updateNotificationTemplate).Methods(http.MethodPut, http.MethodOptions)
	authorized.HandleFunc("/notification/email/template/{template_id}", s.deleteNotificationTemplates).Methods(http.MethodDelete, http.MethodOptions)

	authorized.HandleFunc("/media/upload", s.uploadMedia).Methods(http.MethodPost, http.MethodOptions)

	authorized.HandleFunc("/subject", s.listSubjects).Methods(http.MethodGet, http.MethodOptions)
	authorized.HandleFunc("/subject", s.createSubject).Methods(http.MethodPost, http.MethodOptions)
	admin.HandleFunc("/subject", s.updateSubject).Methods(http.MethodPut, http.MethodOptions)
	admin.HandleFunc("/subject/{subject_id}", s.deleteSubject).Methods(http.MethodDelete, http.MethodOptions)

	admin.HandleFunc("/faculty/{faculty_id}/groups", s.getGroupsByFacultyID).Methods(http.MethodGet, http.MethodOptions)
	authorized.HandleFunc("/faculty/{faculty_id}", s.getFacultyByID).Methods(http.MethodGet, http.MethodOptions)
	authorized.HandleFunc("/faculty", s.listFaculties).Methods(http.MethodGet, http.MethodOptions)
	admin.HandleFunc("/faculty", s.createFaculty).Methods(http.MethodPost, http.MethodOptions)
	admin.HandleFunc("/faculty", s.updateFaculty).Methods(http.MethodPut, http.MethodOptions)
	admin.HandleFunc("/faculty/{faculty_id}", s.deleteFaculty).Methods(http.MethodDelete, http.MethodOptions)

	authorized.HandleFunc("/event", s.getEvents).Methods(http.MethodGet, http.MethodOptions)
	authorized.HandleFunc("/event", s.createEvent).Methods(http.MethodPost, http.MethodOptions)
	authorized.HandleFunc("/event", s.updateEvent).Methods(http.MethodPut, http.MethodOptions)
	authorized.HandleFunc("/event/{event_id}", s.deleteEvent).Methods(http.MethodDelete, http.MethodOptions)

	authorized.HandleFunc("/activity/student/{student_id}", s.createActivity).Methods(http.MethodPost, http.MethodOptions)
	authorized.HandleFunc("/activity/student/{student_id}", s.getActivitiesByStudentID).Methods(http.MethodGet, http.MethodOptions)
	authorized.HandleFunc("/activity/{activity_id}", s.updateActivity).Methods(http.MethodPut, http.MethodOptions)
	authorized.HandleFunc("/activity/{activity_id}", s.deleteActivity).Methods(http.MethodDelete, http.MethodOptions)

	authorized.HandleFunc("/assignment/subject/{subject_id}", s.createAssignment).Methods(http.MethodPost, http.MethodOptions)
	authorized.HandleFunc("/assignment/subject/{subject_id}", s.getAssignmentBySubjectID).Methods(http.MethodGet, http.MethodOptions)
	authorized.HandleFunc("/assignment/{assignment_id}", s.updateAssignment).Methods(http.MethodPut, http.MethodOptions)
	authorized.HandleFunc("/assignment/{assignment_id}", s.deleteAssignment).Methods(http.MethodDelete, http.MethodOptions)

	authorized.HandleFunc("/assignment/{assignment_id}/submit", s.submitAssignment).Methods(http.MethodPost, http.MethodOptions)

	return router
}

func corsMiddleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Access-Control-Allow-Headers, Authorization, X-Requested-With")

		if req.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		handler.ServeHTTP(w, req)
	})
}

func (s *HTTPServer) Start() error {
	return http.ListenAndServe(":8080", s.router)
}
