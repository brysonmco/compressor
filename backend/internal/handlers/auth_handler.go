package handlers

import (
	"encoding/json"
	"errors"
	"github.com/awesomebfm/compressor/internal/auth"
	"github.com/awesomebfm/compressor/internal/db"
	"github.com/awesomebfm/compressor/internal/models"
	"github.com/awesomebfm/compressor/internal/subscriptions"
	"github.com/awesomebfm/compressor/internal/utils"
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
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
	r.Post("/login", h.login)
	r.Post("/signup", h.signUp)
	r.Post("/refresh", h.refresh)
	r.Post("/logout", h.logout)

	return r
}

type loginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *AuthHandler) login(w http.ResponseWriter, r *http.Request) {
	var data loginRequest
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		log.Printf("error parsing login request JSON: %v", err)
		utils.WriteError(w, "error parsing JSON", http.StatusBadRequest, "invalid_json", nil)
		return
	}
	defer r.Body.Close()

	details := map[string]interface{}{}

	if data.Email == "" {
		details["email"] = "missing required field"
	}
	if data.Password == "" {
		details["password"] = "missing required field"
	}

	if len(details) > 0 {
		utils.WriteError(w, "missing required fields", http.StatusBadRequest, "missing_fields", details)
		return
	}

	// Fetch user's account from the database
	user, err := h.Database.FindUserByEmail(r.Context(), data.Email)
	if err != nil && errors.Is(err, pgx.ErrNoRows) {
		utils.WriteError(w, "invalid credentials", http.StatusUnauthorized, "invalid_credentials", nil)
		return
	} else if err != nil {
		log.Printf("error finding user by email: %v", err)
		utils.WriteError(w, "error logging in", http.StatusInternalServerError, "internal_error", nil)
		return
	}

	// Check password
	if !h.Auth.CheckPasswordHash(data.Password, user.PasswordHash) {
		utils.WriteError(w, "invalid credentials", http.StatusUnauthorized, "invalid_credentials", nil)
		return
	}

	// Generate access token
	accessToken, err := h.Auth.GenerateAccessToken(user.Id)
	if err != nil {
		log.Printf("error generating access token: %v", err)
		utils.WriteError(w, "error logging in", http.StatusInternalServerError, "internal_error", nil)
		return
	}

	// Generate refresh token
	refreshToken, err := h.Auth.GenerateRefreshToken()
	if err != nil {
		log.Printf("error generating refresh token: %v", err)
		utils.WriteError(w, "error logging in", http.StatusInternalServerError, "internal_error", nil)
		return
	}

	// Set refresh token cookie
	switch os.Getenv("DEPLOYMENT_TARGET") {
	case "development":
		http.SetCookie(w, &http.Cookie{
			Name:     "refreshToken",
			Value:    refreshToken,
			HttpOnly: true,
			Secure:   false,
			Path:     "v1/auth/refresh",
			SameSite: http.SameSiteStrictMode,
		})
	default:
		http.SetCookie(w, &http.Cookie{
			Name:     "refreshToken",
			Value:    refreshToken,
			HttpOnly: true,
			Secure:   true,
			Path:     "v1/auth/refresh",
			SameSite: http.SameSiteStrictMode,
			Domain:   "api-compressor.brysonmcbreen.dev",
		})
	}

	// Hash refresh token
	hashedRefreshToken := h.Auth.HashRefreshToken(refreshToken)

	// Persist session
	now := time.Now()
	_, err = h.Database.CreateSession(r.Context(), &models.CreateSession{
		TokenHash: hashedRefreshToken,
		UserId:    user.Id,
		ExpiresAt: now.Add(time.Hour * 24 * 30),
		CreatedAt: now,
	})
	if err != nil {
		log.Printf("error creating session: %v", err)
		utils.WriteError(w, "error logging in", http.StatusInternalServerError, "internal_error", nil)
		return
	}

	// Return access token
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(map[string]string{
		"accessToken": accessToken,
	})
	if err != nil {
		log.Printf("error encoding JSON response: %v", err)
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
	var data signUpRequest
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		log.Printf("error parsing sign up request JSON: %v", err)
		utils.WriteError(w, "error parsing JSON", http.StatusBadRequest, "invalid_json", nil)
		return
	}
	defer r.Body.Close()

	details := map[string]interface{}{}

	if data.Email == "" {
		details["email"] = "missing required field"
	}
	if data.FirstName == "" {
		details["firstName"] = "missing required field"
	}
	if data.LastName == "" {
		details["lastName"] = "missing required field"
	}
	if data.Password == "" {
		details["password"] = "missing required field"
	}
	if data.ConfirmPassword == "" {
		details["confirmPassword"] = "missing required field"
	}

	if len(details) > 0 {
		utils.WriteError(w, "missing required fields", http.StatusBadRequest, "missing_fields", details)
		return
	}

	if data.Password != data.ConfirmPassword {
		utils.WriteError(w, "passwords do not match", http.StatusBadRequest, "password_mismatch", nil)
		return
	}

	// Ensure account does not already exist
	_, err = h.Database.FindUserByEmail(r.Context(), data.Email)
	if err == nil {
		utils.WriteError(w, "account already exists", http.StatusBadRequest, "account_exists", nil)
		return
	} else if !errors.Is(err, pgx.ErrNoRows) {
		log.Printf("error finding user by email: %v", err)
		utils.WriteError(w, "error creating account", http.StatusInternalServerError, "internal_error", nil)
		return
	}

	// Hash password
	hashedPassword, err := h.Auth.HashPassword(data.Password)
	if err != nil {
		log.Printf("error hashing password: %v", err)
		utils.WriteError(w, "error creating account", http.StatusInternalServerError, "internal_error", nil)
		return
	}

	// Persist account
	user, err := h.Database.CreateUser(r.Context(), &models.CreateUser{
		Email:        data.Email,
		FirstName:    data.FirstName,
		LastName:     data.LastName,
		PasswordHash: hashedPassword,
	})
	if err != nil {
		log.Printf("error creating user: %v", err)
		utils.WriteError(w, "error creating account", http.StatusInternalServerError, "internal_error", nil)
		return
	}

	// Create Stripe customer
	customer, err := subscriptions.CreateStripeCustomer(user)
	if err != nil {
		log.Printf("error creating customer: %v", err)
		utils.WriteError(w, "error creating account", http.StatusInternalServerError, "internal_error", nil)
		return
	}
	user.StripeCustomerId = customer.ID
	err = h.Database.UpdateUser(r.Context(), user)
	if err != nil {
		log.Printf("error updating user: %v", err)
		utils.WriteError(w, "error creating account", http.StatusInternalServerError, "internal_error", nil)
		return
	}

	// Generate access token
	accessToken, err := h.Auth.GenerateAccessToken(user.Id)
	if err != nil {
		log.Printf("error generating access token: %v", err)
		utils.WriteError(w, "error creating account", http.StatusInternalServerError, "internal_error", nil)
		return
	}

	// Generate refresh token
	refreshToken, err := h.Auth.GenerateRefreshToken()
	if err != nil {
		log.Printf("error generating refresh token: %v", err)
		utils.WriteError(w, "error creating account", http.StatusInternalServerError, "internal_error", nil)
		return
	}

	// Set refresh token cookie
	switch os.Getenv("DEPLOYMENT_TARGET") {
	case "development":
		http.SetCookie(w, &http.Cookie{
			Name:     "refreshToken",
			Value:    refreshToken,
			HttpOnly: true,
			Secure:   false,
			Path:     "v1/auth/refresh",
			SameSite: http.SameSiteStrictMode,
		})
		break
	default:
		http.SetCookie(w, &http.Cookie{
			Name:     "refreshToken",
			Value:    refreshToken,
			HttpOnly: true,
			Secure:   true,
			Path:     "v1/auth/refresh",
			SameSite: http.SameSiteStrictMode,
			Domain:   "api-compressor.brysonmcbreen.dev",
		})
	}

	// Hash refresh token
	hashedRefreshToken := h.Auth.HashRefreshToken(refreshToken)

	// Persist session
	now := time.Now()
	_, err = h.Database.CreateSession(r.Context(), &models.CreateSession{
		TokenHash: hashedRefreshToken,
		UserId:    user.Id,
		ExpiresAt: now.Add(time.Hour * 24 * 30),
		CreatedAt: now,
	})
	if err != nil {
		log.Printf("error creating session: %v", err)
		utils.WriteError(w, "error creating account", http.StatusInternalServerError, "internal_error", nil)
		return
	}

	// Return access token
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(map[string]string{
		"accessToken": accessToken,
	})
	if err != nil {
		log.Printf("error encoding JSON response: %v", err)
	}
}

func (h *AuthHandler) refresh(w http.ResponseWriter, r *http.Request) {
	token, err := r.Cookie("refreshToken")
	if err != nil {
		utils.WriteError(w, "missing refresh token", http.StatusBadRequest, "missing_refresh_token", nil)
		return
	}

	hashedRefreshToken := h.Auth.HashRefreshToken(token.Value)

	session, err := h.Database.FindSessionByTokenHash(r.Context(), hashedRefreshToken)
	if err != nil && errors.Is(err, pgx.ErrNoRows) {
		utils.WriteError(w, "invalid refresh token", http.StatusUnauthorized, "invalid_token", nil)
		return
	} else if err != nil {
		log.Printf("error finding session by token hash: %v", err)
		utils.WriteError(w, "error refreshing token", http.StatusInternalServerError, "internal_error", nil)
		return
	}

	// Ensure session is still valid
	if session.ExpiresAt.Before(time.Now()) {
		utils.WriteError(w, "refresh token expired", http.StatusUnauthorized, "expired_token", nil)
		return
	}

	// Ensure session has not been revoked
	if session.Revoked {
		utils.WriteError(w, "refresh token revoked", http.StatusUnauthorized, "revoked_token", nil)
		return
	}

	// Ensure session was not created in the future
	if session.CreatedAt.After(time.Now()) {
		utils.WriteError(w, "refresh token created in the future", http.StatusUnauthorized, "refresh_token_created_in_future", nil)
		return
	}

	// Generate access token
	accessToken, err := h.Auth.GenerateAccessToken(session.UserId)
	if err != nil {
		log.Printf("error generating access token: %v", err)
		utils.WriteError(w, "error refreshing token", http.StatusInternalServerError, "internal_error", nil)
		return
	}

	// Return access token
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(map[string]string{
		"accessToken": accessToken,
	})
	if err != nil {
		log.Printf("error encoding JSON response: %v", err)
	}
}

func (h *AuthHandler) logout(w http.ResponseWriter, r *http.Request) {
	id, err := h.Auth.ValidateAccessToken(strings.TrimPrefix(r.Header.Get("Authorization"), "Bearer "))
	if err != nil && errors.Is(err, errors.New("expired_token")) {
		utils.WriteError(w, "token has expired", http.StatusUnauthorized, "expired_token", nil)
		return
	} else if err != nil {
		utils.WriteError(w, "token is invalid", http.StatusUnauthorized, "invalid_token", nil)
		return
	}

	err = h.Database.RevokeAllSessionsByUserId(r.Context(), id)
	if err != nil {
		log.Printf("error revoking all sessions by user ID: %v", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(map[string]bool{
		"success": true,
	})
	if err != nil {
		log.Printf("error encoding JSON response: %v", err)
	}
}
