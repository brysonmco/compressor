package handlers

import (
	"github.com/brysonmco/compressor/internal/db"
	"github.com/brysonmco/compressor/internal/middleware"
	"github.com/brysonmco/compressor/internal/utils"
	"github.com/go-chi/chi/v5"
	"net/http"
	"time"
)

type UserHandler struct {
	Database       *db.Database
	AuthMiddleware *middleware.AuthMiddleware
}

func NewUserHandler(
	database *db.Database,
	authMiddleware *middleware.AuthMiddleware,
) http.Handler {
	h := &UserHandler{
		Database:       database,
		AuthMiddleware: authMiddleware,
	}

	r := chi.NewRouter()
	r.With(authMiddleware.Protected).Get("/profile", h.handleGetProfile)

	return r
}

// GET /users/profile
func (h *UserHandler) handleGetProfile(w http.ResponseWriter, r *http.Request) {
	// Grab their ID
	id := r.Context().Value("userId").(int64)

	// Grab user data
	user, err := h.Database.FindUserByID(r.Context(), id)
	if err != nil {
		utils.WriteError(w, r, http.StatusInternalServerError, "error fetching user", "user_not_found", nil)
		return
	}

	// Get subscription
	// TODO: Implement

	// Get tokens
	// TODO: Implement

	// Return user data
	utils.WriteSuccess(w, r, http.StatusOK, "user profile", map[string]interface{}{
		"email":     user.Email,
		"firstName": user.FirstName,
		"lastName":  user.LastName,
		"subscription": map[string]interface{}{
			"active":    false,
			"plan":      0,
			"periodEnd": time.Now().Add(time.Hour * 24),
		},
		"tokens": 1000, // TODO: Implement token fetching
	})
}
