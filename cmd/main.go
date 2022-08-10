package main

import (
	"auth-service/api"
	"auth-service/config"
	"log"
)

func main() {

	cfg := &config.Config{}
	cfg.LoadConfig()

	log.Printf("Starting Auth server on %s", cfg.AppPort)

	s := api.NewServer(cfg)
	if err := s.Start(); err != nil {
		log.Fatal("error starting the server")
	}

}
