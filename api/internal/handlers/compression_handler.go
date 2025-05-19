package handlers

import (
	"encoding/json"
	"github.com/awesomebfm/compressor/internal/db"
	"github.com/awesomebfm/compressor/internal/middleware"
	"github.com/awesomebfm/compressor/internal/models"
	"github.com/awesomebfm/compressor/internal/storage"
	"github.com/awesomebfm/compressor/internal/utils"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"strconv"
	"time"
)

type CompressionHandler struct {
	Database       *db.Database
	AuthMiddleware *middleware.AuthMiddleware
	Storage        *storage.Storage
}

func NewCompressionHandler(
	database *db.Database,
	authMiddleware *middleware.AuthMiddleware,
	strge *storage.Storage,
) http.Handler {
	h := &CompressionHandler{
		Database:       database,
		AuthMiddleware: authMiddleware,
		Storage:        strge,
	}

	r := chi.NewRouter()
	r.With(authMiddleware.Protected).Post("/new", h.handleCreateCompressionJob)
	r.With(authMiddleware.Protected).Post("/upload-complete", h.handleUploadComplete)

	return r
}

func (h *CompressionHandler) handleCreateCompressionJob(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value("userID").(int64)

	// Create job
	job, err := h.Database.CreateJob(r.Context(), &models.CreateJob{
		UserId: id,
	})
	if err != nil {
		log.Printf("error creating job: %v", err)
		utils.WriteError(w, "error creating job", http.StatusInternalServerError, "internal_error", nil)
		return
	}

	// Generate upload URL
	uploadURL, err := h.Storage.GenerateUploadURLForUploads(r.Context(), job.Id, "mp4", time.Now().Add(time.Hour))
	if err != nil {
		log.Printf("error generating upload URL: %v", err)
		utils.WriteError(w, "error creating job", http.StatusInternalServerError, "internal_error", nil)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(map[string]string{
		"uploadURL": uploadURL,
		"JobId":     strconv.FormatInt(job.Id, 10),
	})
	if err != nil {
		log.Printf("error encoding JSON response: %v", err)
	}
}

func (h *CompressionHandler) handleUploadComplete(w http.ResponseWriter, r *http.Request) {

}
