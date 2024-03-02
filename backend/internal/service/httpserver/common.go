package httpserver

import (
	"encoding/json"
	"net/http"
)

func (s *HTTPServer) respond(w http.ResponseWriter, status int, payload any) error {
	if payload != nil {
		w.Header().Set("Content-Type", "application/json")
	}

	w.WriteHeader(status)

	if payload == nil {
		return nil
	}

	return json.NewEncoder(w).Encode(payload)
}

func (s *HTTPServer) respondError(w http.ResponseWriter, status int, err error) {
	s.respond(w, status, map[string]string{"error": err.Error()})
}

func (s *HTTPServer) decodeJSON(r *http.Request, v any) error {
	return json.NewDecoder(r.Body).Decode(v)
}
