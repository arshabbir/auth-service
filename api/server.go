package api

import (
	"auth-service/config"
	"auth-service/dao"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

type server struct {
	cfg *config.Config
	db  dao.AuthDaO
}

type Server interface {
	Start() error
}

func NewServer(cfg *config.Config) Server {
	db := dao.NewAuthDao(cfg.DSN)
	return &server{cfg: cfg, db: db}
}
func (s *server) Start() error {
	//  Add routes

	mux := routes(s)
	srv := http.Server{
		Addr:    s.cfg.AppPort,
		Handler: mux,
	}

	srv.ListenAndServe()

	return nil

}
func routes(s *server) http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"https://*", "http://*"},
	}))

	r.Get("/ping", s.HandlePing)
	r.Post("/", s.HandleAuth)

	return r
}
