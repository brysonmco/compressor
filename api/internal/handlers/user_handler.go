package handlers

import (
	"errors"
	"github.com/awesomebfm/compressor/internal/db"
	"github.com/awesomebfm/compressor/internal/middleware"
	"github.com/awesomebfm/compressor/internal/utils"
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5"
	"net/http"
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
	r.With(authMiddleware.Protected).Post("/profile", h.handleGetProfile)

	return r
}

// GET /users/profile
func (h *UserHandler) handleGetProfile(w http.ResponseWriter, r *http.Request) {
	// Grab their ID
	id := r.Context().Value("userID").(int64)

	// Grab user data
	user, err := h.Database.FindUserByID(r.Context(), id)
	if err != nil {
		utils.WriteError(w, r, http.StatusInternalServerError, "error fetching user", "user_not_found", nil)
		return
	}

	// Get subscription
	_, err = h.Database.FindActiveSubscriptionByUserId(r.Context(), id)
	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		utils.WriteError(w, r, http.StatusInternalServerError, "error fetching subscription", "subscription_not_found", nil)
		return
	} else if err != nil {
		// IRDK
	}

	// Get tokens
	// TODO: Implement

	// Return user data
	utils.WriteSuccess(w, r, http.StatusOK, "user profile", map[string]interface{}{
		"email": user.Email,
	})
}
