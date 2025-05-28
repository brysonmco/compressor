package handlers

import (
	"errors"
	"github.com/awesomebfm/compressor/internal/db"
	"github.com/awesomebfm/compressor/internal/middleware"
	"github.com/awesomebfm/compressor/internal/models"
	"github.com/awesomebfm/compressor/internal/utils"
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5"
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
	// TODO: This still needs work
	subscription, err := h.Database.FindActiveSubscriptionByUserId(r.Context(), id)
	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		utils.WriteError(w, r, http.StatusInternalServerError, "internal error", "internal_error", nil)
		return
	} else if err != nil {
		subscription = &models.Subscription{
			PlanId:           0,
			Status:           "active",
			CurrentPeriodEnd: time.Now().AddDate(0, 1, 0), // Default to one month from now
		}
	}

	// Get tokens
	// TODO: Implement

	// Return user data
	utils.WriteSuccess(w, r, http.StatusOK, "user profile", map[string]interface{}{
		"email":     user.Email,
		"firstName": user.FirstName,
		"lastName":  user.LastName,
		"subscription": map[string]interface{}{
			"active":    true,
			"plan":      subscription.PlanId,
			"periodEnd": subscription.CurrentPeriodEnd,
		},
		"tokens": 1000, // TODO: Implement token fetching
	})
}
