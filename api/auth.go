package api

import (
	"auth-service/models"
	"auth-service/utils"
	"encoding/json"
	"log"
	"net/http"
)

func (s *server) HandleAuth(w http.ResponseWriter, r *http.Request) {

	user := models.User{}

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		utils.SendError(w, http.StatusBadGateway, err.Error())
		return
	}

	// Validate the password by checking against DB
	usr, err := s.db.GetByEmail(user.Email)
	if err != nil {
		utils.SendError(w, http.StatusUnauthorized, err.Error())
		return

	}
	log.Println("Validating login")
	log.Println("Recived user object : ", usr)
	if usr.Password != user.Password {
		// Need toimplemnet tocken generation logic
		utils.SendError(w, http.StatusUnauthorized, usr.Password)
		return
	}
	log.Println("Login successful")
	userResponse := models.UserResponse{ID: usr.ID, CreatedAt: usr.CreatedAt, Active: usr.Active, UpdatedAt: usr.UpdatedAt}
	log.Println("Sending user response  ", userResponse)
	if err := json.NewEncoder(w).Encode(&userResponse); err != nil {
		log.Println("error sending the reponse")
	}
}
