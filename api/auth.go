package api

import (
	"auth-service/models"
	"auth-service/utils"
	"encoding/json"
	"log"
	"net/http"
)

func (s *server) HandleAuth(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

	w.Write([]byte("auth handler invoked"))

	// Need to implement the auth logic using DAO
	user := models.User{}

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		utils.SendError(w, http.StatusBadGateway, err.Error())
		return
	}

	userResponse := models.UserResponse{ID: user.ID, CreatedAt: user.CreatedAt, Active: user.Active, UpdatedAt: user.UpdatedAt}

	if err := json.NewEncoder(w).Encode(&userResponse); err != nil {
		log.Println("error sending the reponse")
	}
}
