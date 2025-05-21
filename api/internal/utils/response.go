package utils

import (
	"encoding/json"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"net/http"
	"time"
)

type ErrorResponse struct { // Human-readable
	Error   string      `json:"error"` // Machine-readable
	Details interface{} `json:"details"`
}

func WriteError(
	w http.ResponseWriter,
	r *http.Request,
	code int,
	message string,
	error string,
	details interface{},
) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(map[string]interface{}{
		"success":   false,
		"status":    code,
		"requestId": middleware.GetReqID(r.Context()),
		"timestamp": time.Now(),
		"message":   message,
		"error": ErrorResponse{
			Error:   error,
			Details: details,
		},
	}); err != nil {
		log.Printf("error encoding JSON response: %v", err)
	}
}

type SuccessResponse struct {
	Data interface{} `json:"data"`
}

func WriteSuccess(
	w http.ResponseWriter,
	r *http.Request,
	code int,
	message string,
	data interface{},
) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(map[string]interface{}{
		"success":   true,
		"status":    code,
		"requestId": middleware.GetReqID(r.Context()),
		"timestamp": time.Now(),
		"message":   message,
		"data":      data,
	}); err != nil {
		log.Printf("error encoding JSON response: %v", err)
	}
}
