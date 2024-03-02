package httpserver

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/google"
	"oss-backend/internal/config"
)

func (s *HTTPServer) newRouter(_ config.Config) *mux.Router {
	router := mux.NewRouter()

	goth.UseProviders(
		google.New(s.googleOAuthCfg.ClientID, s.googleOAuthCfg.ClientSecret, s.googleOAuthCfg.RedirectURL, s.googleOAuthCfg.Scopes...),
	)

	router.Use(corsMiddleware)

	api := router.PathPrefix("/api/v1").Subrouter()

	protected := api.PathPrefix("/").Subrouter()
	protected.Use(s.authMiddleware)

	api.HandleFunc("/auth/{provider}/callback", s.oauthCallback).Methods(http.MethodGet, http.MethodOptions)

	api.HandleFunc("/status", s.getStatus).Methods(http.MethodGet, http.MethodOptions)
	protected.HandleFunc("/profile/me", s.getMe).Methods(http.MethodGet, http.MethodOptions)

	protected.HandleFunc("/media/upload", s.uploadMedia).Methods(http.MethodPost, http.MethodOptions)

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
