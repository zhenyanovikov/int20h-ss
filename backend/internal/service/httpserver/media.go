package httpserver

import (
	"net/http"
	"path/filepath"
)

func (s *HTTPServer) uploadMedia(w http.ResponseWriter, r *http.Request) {
	file, fileHeaders, err := r.FormFile("image")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	defer file.Close()

	ext := filepath.Ext(fileHeaders.Filename)

	url, err := s.mediaSrv.Upload(r.Context(), file, ext)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	s.respond(w, http.StatusOK, map[string]string{"url": url})
}
