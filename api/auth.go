package api

import "net/http"

func (s *server) HandleAuth(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("auth handler invoked"))
}
