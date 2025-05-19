package utils

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type ErrorResponse struct {
	Message   string      `json:"message"` // Human-readable
	Code      int         `json:"code"`
	Error     string      `json:"error"` // Machine-readable
	Details   interface{} `json:"details"`
	Timestamp time.Time   `json:"timestamp"`
}

func WriteError(
	w http.ResponseWriter,
	message string,
	code int,
	error string,
	details interface{},
) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(ErrorResponse{
		Message:   message,
		Code:      code,
		Error:     error,
		Details:   details,
		Timestamp: time.Now(),
	}); err != nil {
		log.Printf("error encoding JSON response: %v", err)
	}
}
