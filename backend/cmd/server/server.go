package server

import (
	"log"
	"net/http"

	"github.com/DiegoUrrego4/newsletter-app/internal/adapters/http/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Server struct{}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Run() {
	r := s.registerRouter()

	log.Println("Starting server at port :8080")

	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("Could not start server: %v", err)
	}
}

func (s *Server) registerRouter() chi.Router {
	router := chi.NewRouter()

	router.Use(middleware.Recoverer)

	router.Get("/ping", handlers.PingHandler)

	return router
}
