package main

import (
	"github.com/awesomebfm/compressor/internal/auth"
	"github.com/awesomebfm/compressor/internal/db"
	"github.com/awesomebfm/compressor/internal/handlers"
	internalmiddleware "github.com/awesomebfm/compressor/internal/middleware"
	"github.com/awesomebfm/compressor/internal/storage"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/stripe/stripe-go/v82"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	// Stripe
	stripe.Key = os.Getenv("STRIPE_SECRET_KEY")

	// Database
	database, err := db.NewDatabase(os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	// Auth
	ath := auth.NewAuth()
	authMiddleware := internalmiddleware.NewAuthMiddleware(ath, database)

	// Storage
	strge, err := storage.NewStorage(
		os.Getenv("S3_UPLOADS_BUCKET"),
		os.Getenv("S3_ENDPOINT"),
		os.Getenv("S3_ACCESS_KEY"),
		os.Getenv("S3_SECRET_KEY"),
		os.Getenv("DEPLOYMENT_TARGET") != "development",
	)
	if err != nil {
		log.Fatalf("failed to connect to object storage: %v", err)
	}

	// Router
	r := chi.NewRouter()

	// CORS
	var allowedOrigins []string
	switch os.Getenv("DEPLOYMENT_TARGET") {
	case "development":
		allowedOrigins = []string{"http://localhost:5173"}
		break
	default:
		allowedOrigins = []string{"https://compressor.brysonmcbreen.dev"}
	}
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   allowedOrigins,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	// Middleware
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	// Handlers
	r.Mount("/v1/auth", handlers.NewAuthHandler(database, ath))
	r.Mount("/v1/subscriptions", handlers.NewSubscriptionHandler(
		database,
		authMiddleware,
		os.Getenv("STRIPE_ENDPOINT_SECRET")))
	r.Mount("/v1/compress", handlers.NewCompressionHandler(
		database,
		authMiddleware,
		strge))

	log.Fatal(http.ListenAndServe(os.Getenv("LISTEN_ADDR"), r))
}
