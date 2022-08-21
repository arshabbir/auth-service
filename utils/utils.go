package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

type apiError struct {
	Statuscode int    `json:"statuscode"`
	Message    string `json:"message"`
}

func SendError(w http.ResponseWriter, statuscode int, msg string) error {
	w.WriteHeader(statuscode)
	w.Header().Add("Content-type", "application/json")
	payload := apiError{Statuscode: statuscode, Message: msg}
	if err := json.NewEncoder(w).Encode(&payload); err != nil {
		log.Println("error sending the message ")
		return err
	}
	return nil

}
