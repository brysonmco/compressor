package handlers

import (
	"github.com/awesomebfm/compressor/internal/db"
	"github.com/go-chi/chi/v5"
	"net/http"
)

type UserHandler struct {
	Database *db.Database
}

func NewUserHandler(
	database *db.Database,
) http.Handler {
	h := &UserHandler{
		Database: database,
	}

	r := chi.NewRouter()
	r.Get("/users/:id/profile", h.getProfile)

	return r
}

// GET /users/:id/profile
func (h *UserHandler) getProfile(w http.ResponseWriter, r *http.Request) {

}
