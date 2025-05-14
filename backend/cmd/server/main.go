package main

import (
	"github.com/awesomebfm/compressor/internal/auth"
	"github.com/awesomebfm/compressor/internal/db"
	"github.com/awesomebfm/compressor/internal/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	// Database
	database, err := db.NewDatabase(os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	// Auth
	ath := auth.NewAuth()

	// Router
	r := chi.NewRouter()

	// Middleware
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	// Handlers
	r.Mount("/v1", handlers.NewAuthHandler(database, ath))

	log.Fatal(http.ListenAndServe(os.Getenv("LISTEN_ADDR"), r))
}
