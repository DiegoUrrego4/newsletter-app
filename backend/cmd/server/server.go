package server

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/DiegoUrrego4/newsletter-app/internal/adapters/http/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type Server struct {
	db *sql.DB
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Run() {
	if err := s.loadConfig(); err != nil {
		log.Fatalf("Error loading configuration: %v", err)
	}

	if err := s.connectDatabase(); err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	defer s.db.Close()

	if err := s.applyMigrations(); err != nil {
		log.Fatalf("Error applying migrations: %v", err)
	}

	r := s.registerRouter()

	log.Println("Starting server at port :8080")

	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("Could not start server: %v", err)
	}
}

func (s *Server) loadConfig() error {
	if err := godotenv.Load(); err != nil {
		return fmt.Errorf("error loading .env file: %w", err)
	}

	return nil
}

func (s *Server) connectDatabase() error {

	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"))

	log.Println("Attempting to connect to database...")
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return fmt.Errorf("error connecting to database: %w", err)
	}

	if err := db.Ping(); err != nil {
		return fmt.Errorf("error pinging database: %w", err)
	}

	s.db = db
	log.Println("Connected to database")
	return nil
}

func (s *Server) registerRouter() chi.Router {
	router := chi.NewRouter()

	router.Use(middleware.Recoverer)

	router.Get("/ping", handlers.PingHandler)

	// TODO: Add repositories
	//newsletterRepository := repository.NewPostgresNewsletterRepository(s.db)

	// TODO: add services
	// TODO: add handlers

	return router
}

func (s *Server) applyMigrations() error {
	driver, err := postgres.WithInstance(s.db, &postgres.Config{})
	if err != nil {
		return fmt.Errorf("error creando el driver de migración: %w", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://migrations",
		"postgres", driver)
	if err != nil {
		return fmt.Errorf("error creando instancia de migración: %w", err)
	}

	if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return fmt.Errorf("error aplicando migraciones: %w", err)
	}

	return nil
}
