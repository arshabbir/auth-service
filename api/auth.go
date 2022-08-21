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

	// Validate the password by checking against DB
	usr, err := s.db.GetByEmail(user.Email)
	if err != nil {
		utils.SendError(w, http.StatusUnauthorized, err.Error())

	}
	log.Println("Validating login")
	if usr.Password != user.Password {
		// Need toimplemnet tocken generation logic
		utils.SendError(w, http.StatusUnauthorized, usr.Password)
		return
	}
	log.Println("Login successful")
	userResponse := models.UserResponse{ID: usr.ID, CreatedAt: usr.CreatedAt, Active: usr.Active, UpdatedAt: usr.UpdatedAt}
	if err := json.NewEncoder(w).Encode(&userResponse); err != nil {
		log.Println("error sending the reponse")
	}
}
