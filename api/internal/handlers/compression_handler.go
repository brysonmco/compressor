package handlers

import (
	"encoding/json"
	"github.com/brysonmco/compressor/internal/db"
	"github.com/brysonmco/compressor/internal/middleware"
	"github.com/brysonmco/compressor/internal/models"
	"github.com/brysonmco/compressor/internal/storage"
	"github.com/brysonmco/compressor/internal/utils"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
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
	FileName      string `json:"fileName"`
	FileContainer string `json:"fileContainer"`
}

func (h *CompressionHandler) handleCreateCompressionJob(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value("userId").(int64)

	// Parse request body
	var req createCompressionJobRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.WriteError(w, r, http.StatusBadRequest, "error parsing JSON", "invalid_json", nil)
		return
	}

	// Validate request
	if req.FileName == "" || req.FileContainer == "" {
		utils.WriteError(w, r, http.StatusBadRequest, "missing required fields", "missing_fields", nil)
		return
	}

	// Get their subscription
	// TODO: Implement subscription handling
	/*var plan *models.Plan
	subscription, err := h.Database.FindActiveSubscriptionByUserId(r.Context(), id)
	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		log.Printf("error fetching subscription: %v", err)
		utils.WriteError(w, r, http.StatusInternalServerError, "error creating job", "internal_error", nil)
		return
	} else if errors.Is(err, pgx.ErrNoRows) {
		plan, err = h.Database.FindPlanByName(r.Context(), "Free")
		if err != nil {
			log.Printf("error fetching plan: %v", err)
			utils.WriteError(w, r, http.StatusInternalServerError, "error creating job", "internal_error", nil)
			return
		}
	} else {
		plan, err = h.Database.FindPlanById(r.Context(), subscription.PlanId)
		if err != nil {
			log.Printf("error fetching plan: %v", err)
			utils.WriteError(w, r, http.StatusInternalServerError, "error creating job", "internal_error", nil)
			return
		}
	}*/

	// Ensure valid container
	allowedContainers := []string{"mp4", "mkv", "mov", "avi", "webm", "flv", "ts", "mpg", "ogg", "wav"}
	for i := 0; i < len(allowedContainers); i++ {
		if req.FileContainer == allowedContainers[i] {
			break
		}
		if i == len(allowedContainers)-1 {
			utils.WriteError(w, r, http.StatusBadRequest, "invalid container", "invalid_container", nil)
			return
		}
	}

	// Create job
	job, err := h.Database.CreateJob(r.Context(), &models.CreateJob{
		UserId:   id,
		FileName: req.FileName,
	})
	if err != nil {
		log.Printf("error creating job: %v", err)
		utils.WriteError(w, r, http.StatusInternalServerError, "error creating job", "internal_error", nil)
		return
	}

	// Generate upload URL
	uploadURL, formData, err := h.Storage.GenerateUploadURLForUploads(
		r.Context(),
		job.Id,
		req.FileContainer,
		time.Now().Add(time.Hour),
		10240,
	)
	if err != nil {
		log.Printf("error generating upload URL: %v", err)
		utils.WriteError(w, r, http.StatusInternalServerError, "error creating job", "internal_error", nil)
	}

	utils.WriteSuccess(w, r, http.StatusOK, "job created", map[string]interface{}{
		"jobId":     job.Id,
		"uploadUrl": uploadURL,
		"formData":  formData,
	})
}

type uploadCompleteRequest struct {
	JobId int64 `json:"jobId"`
}

func (h *CompressionHandler) handleUploadComplete(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value("userId").(int64)

	// Parse request body
	var req uploadCompleteRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.WriteError(w, r, http.StatusBadRequest, "error parsing JSON", "invalid_json", nil)
		return
	}

	// Find job
	job, err := h.Database.FindJobById(r.Context(), req.JobId)
	if err != nil {
		utils.WriteError(w, r, http.StatusBadRequest, "job not found", "job_not_found", nil)
		return
	}

	if job.UserId != id {
		// We don't want to leak information about another user's jobs
		utils.WriteError(w, r, http.StatusBadRequest, "job not found", "job_not_found", nil)
		return
	}

	// Ensure file has not yet been uploaded
	if job.FileUploaded {
		utils.WriteError(w, r, http.StatusBadRequest, "file already uploaded", "file_already_uploaded", nil)
		return
	}

	// Check that the file was actually uploaded
	inUploads, err := h.Storage.FileInUploads(r.Context(), job.Id, job.InputContainer)
	if err != nil {
		log.Printf("error checking if file exists: %v", err)
		utils.WriteError(w, r, http.StatusInternalServerError, "error checking if file exists", "internal_error", nil)
		return
	}
	if !inUploads {
		utils.WriteError(w, r, http.StatusBadRequest, "file not found", "file_not_found", nil)
		return
	}

	// Update job
	job.FileUploaded = true
	err = h.Database.UpdateJob(r.Context(), job)
	if err != nil {
		log.Printf("error updating job: %v", err)
		utils.WriteError(w, r, http.StatusInternalServerError, "internal service error", "internal_error", nil)
		return
	}

	// Tell compression-service to provision a VM
	// TODO: Implement

	utils.WriteSuccess(w, r, http.StatusOK, "file uploaded", nil)
}
