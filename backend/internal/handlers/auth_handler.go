package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/awesomebfm/compressor/internal/auth"
	"github.com/awesomebfm/compressor/internal/db"
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5"
	"io"
	"log"
	"net/http"
	"os"
)

type AuthHandler struct {
	Database *db.Database
	Auth     *auth.Auth
}

func NewAuthHandler(
	database *db.Database,
	ath *auth.Auth,
) http.Handler {
	h := &AuthHandler{
		Database: database,
		Auth:     ath,
	}

	r := chi.NewRouter()
	r.Post("/auth/login", h.login)
	r.Post("/auth/signup", h.signUp)
	r.Post("/auth/refresh", h.refresh)
	r.Post("/auth/logout", h.logout)

	return r
}

type loginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *AuthHandler) login(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("error reading login request body: %v", err)
		http.Error(w, "couldn't read request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	var data loginRequest
	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Printf("error parsing login request JSON: %v", err)
		http.Error(w, "error parsing JSON", http.StatusBadRequest)
		return
	}

	// Fetch user's account from the database
	user, err := h.Database.FindUserByEmail(r.Context(), data.Email)
	if err != nil && errors.Is(err, pgx.ErrNoRows) {
		http.Error(w, "invalid credentials", http.StatusUnauthorized)
		return
	} else if err != nil {
		log.Printf("error finding user by email: %v", err)
		http.Error(w, "error communicating with database", http.StatusInternalServerError)
		return
	}

	// Check password
	if !h.Auth.CheckPasswordHash(data.Password, user.PasswordHash) {
		http.Error(w, "invalid credentials", http.StatusUnauthorized)
		return
	}

	// Generate refresh token
	refreshToken, err := h.Auth.GenerateRefreshToken(user.Id)
	if err != nil {
		log.Printf("error generating refresh token: %v", err)
		http.Error(w, "error generating token", http.StatusInternalServerError)
		return
	}

	switch os.Getenv("DEPLOYMENT_TARGET") {
	case "development":
		http.SetCookie(w, &http.Cookie{
			Name:     "refreshToken",
			Value:    refreshToken,
			HttpOnly: true,
			Secure:   false,
			Path:     "/auth/refresh",
			SameSite: http.SameSiteNoneMode,
			Domain:   "localhost:8080",
		})
	default:
		http.SetCookie(w, &http.Cookie{
			Name:     "refreshToken",
			Value:    refreshToken,
			HttpOnly: true,
			Secure:   true,
			Path:     "/auth/refresh",
			SameSite: http.SameSiteNoneMode,
			Domain:   "api-compressor.brysonmcbreen.dev",
		})
	}

}

type signUpRequest struct {
	Email           string `json:"email"`
	FirstName       string `json:"firstName"`
	LastName        string `json:"lastName"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirmPassword"`
}

func (h *AuthHandler) signUp(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Printf("error parsing sign up request form data: %v", err)
		http.Error(w, "error parsing JSON", http.StatusBadRequest)
		return
	}

	var data signUpRequest
	data.Email = r.FormValue("email")
	data.FirstName = r.FormValue("firstName")
	data.LastName = r.FormValue("lastName")
	data.Password = r.FormValue("password")
	data.ConfirmPassword = r.FormValue("confirmPassword")

	if data.Email == "" || data.FirstName == "" || data.LastName == "" || data.Password == "" || data.ConfirmPassword == "" {
		http.Error(w, "missing required fields", http.StatusBadRequest)
		return
	}

	if data.Password != data.ConfirmPassword {
		http.Error(w, "passwords do not match", http.StatusBadRequest)
		return
	}

	// Ensure account does not already exist
	_, err = h.Database.FindUserByEmail(r.Context(), data.Email)
	if err == nil {
		http.Error(w, "account already exists", http.StatusBadRequest)
		return
	} else if !errors.Is(err, pgx.ErrNoRows) {
		log.Printf("error finding user by email: %v", err)
		http.Error(w, "error communicating with database", http.StatusInternalServerError)
		return
	}

	// Hash password
	hashedPassword, err := h.Auth.HashPassword(data.Password)
	if err != nil {
		log.Printf("error hashing password: %v", err)
		http.Error(w, "error creating account", http.StatusInternalServerError)
		return
	}

	fmt.Println(data)
	fmt.Println(hashedPassword)
}

func (h *AuthHandler) refresh(w http.ResponseWriter, r *http.Request) {

}

func (h *AuthHandler) logout(w http.ResponseWriter, r *http.Request) {

}
