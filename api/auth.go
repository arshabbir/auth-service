package api

import "net/http"

func (s *server) HandleAuth(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

	w.Write([]byte("auth handler invoked"))

	// Need to implement the auth logic using DAO
}
