package httpserver

import "net/http"

func (s *HTTPServer) getStatus(w http.ResponseWriter, _ *http.Request) {
	s.respond(w, http.StatusOK, map[string]string{"status": "ok"})
}
