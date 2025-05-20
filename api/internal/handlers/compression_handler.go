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

type createCompressionJobRequest struct {
	Container string `json:"container"`
}

func (h *CompressionHandler) handleCreateCompressionJob(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value("userID").(int64)

	// Parse request body
	var req createCompressionJobRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.WriteError(w, "error parsing JSON", http.StatusBadRequest, "invalid_json", nil)
		return
	}

	// Ensure valid container
	allowedContainers := []string{"mp4", "mkv", "mov", "avi", "webm", "flv", "ts", "mpg", "ogg", "wav"}
	for i := 0; i < len(allowedContainers); i++ {
		if req.Container == allowedContainers[i] {
			break
		}
		if i == len(allowedContainers)-1 {
			utils.WriteError(w, "invalid container", http.StatusBadRequest, "invalid_container", nil)
			return
		}
	}

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

type uploadCompleteRequest struct {
	JobId int64 `json:"jobId"`
}

func (h *CompressionHandler) handleUploadComplete(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value("userID").(int64)

	// Parse request body
	var req uploadCompleteRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.WriteError(w, "error parsing JSON", http.StatusBadRequest, "invalid_json", nil)
		return
	}

	// Find job
	job, err := h.Database.FindJobById(r.Context(), req.JobId)
	if err != nil {
		utils.WriteError(w, "job not found", http.StatusBadRequest, "job_not_found", nil)
		return
	}

	if job.UserId != id {
		// We don't want to leak information about another user's jobs
		utils.WriteError(w, "job not found", http.StatusBadRequest, "job_not_found", nil)
		return
	}

	// Ensure file has not yet been uploaded
	if job.FileUploaded {
		utils.WriteError(w, "file already uploaded", http.StatusBadRequest, "file_already_uploaded", nil)
		return
	}

	// Check that the file was actually uploaded
	inUploads, err := h.Storage.FileInUploads(r.Context(), job.Id, job.InputContainer)
	if err != nil {
		log.Printf("error checking if file exists: %v", err)
		utils.WriteError(w, "error checking if file exists", http.StatusInternalServerError, "internal_error", nil)
		return
	}
	if !inUploads {
		utils.WriteError(w, "file not found", http.StatusBadRequest, "file_not_found", nil)
		return
	}

	// Update job
	job.FileUploaded = true
	err = h.Database.UpdateJob(r.Context(), job)
	if err != nil {
		log.Printf("error updating job: %v", err)
		utils.WriteError(w, "internal service error", http.StatusInternalServerError, "internal_error", nil)
		return
	}

	// Tell compression-service to provision a VM
	// TODO: Implement

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(map[string]bool{
		"success": true,
	})
	if err != nil {
		log.Printf("error encoding JSON response: %v", err)
	}
}
