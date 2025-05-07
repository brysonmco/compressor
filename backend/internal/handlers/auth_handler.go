package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/awesomebfm/compressor/internal/db"
	"github.com/go-chi/chi/v5"
	"io"
	"log"
	"net/http"
)

type AuthHandler struct {
	Database *db.Database
}

func NewAuthHandler(database *db.Database) http.Handler {
	h := &AuthHandler{
		Database: database,
	}

	r := chi.NewRouter()
	r.Post("/login", h.login)
	r.Post("/signup", h.signUp)

	return r
}

type loginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *AuthHandler) login(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("error reading login request body: %v", err)
		http.Error(w, "couldn't read request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	var data loginRequest
	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Printf("error parsing login request JSON: %v", err)
		http.Error(w, "error parsing JSON", http.StatusBadRequest)
		return
	}

	fmt.Println(data)
}

type signUpRequest struct {
	Email           string `json:"email"`
	FirstName       string `json:"firstName"`
	LastName        string `json:"lastName"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirmPassword"`
}

func (h *AuthHandler) signUp(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Printf("error parsing sign up request form data: %v", err)
		http.Error(w, "error parsing JSON", http.StatusBadRequest)
		return
	}

	var data signUpRequest
	data.Email = r.FormValue("email")
	data.FirstName = r.FormValue("firstName")
	data.LastName = r.FormValue("lastName")
	data.Password = r.FormValue("password")
	data.ConfirmPassword = r.FormValue("confirmPassword")

	if data.Email == "" || data.FirstName == "" || data.LastName == "" || data.Password == "" || data.ConfirmPassword == "" {
		http.Error(w, "missing required fields", http.StatusBadRequest)
		return
	}

	if data.Password != data.ConfirmPassword {
		http.Error(w, "passwords do not match", http.StatusBadRequest)
		return
	}

	fmt.Println(data)
}
