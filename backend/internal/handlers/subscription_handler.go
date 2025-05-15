package handlers

import (
	"encoding/json"
	"github.com/awesomebfm/compressor/internal/auth"
	"github.com/awesomebfm/compressor/internal/db"
	"github.com/awesomebfm/compressor/internal/utils"
	"github.com/go-chi/chi/v5"
	"github.com/stripe/stripe-go/v82"
	"github.com/stripe/stripe-go/v82/checkout/session"
	"log"
	"net/http"
	"strings"
)

type SubscriptionHandler struct {
	Database *db.Database
	Auth     *auth.Auth
}

func NewSubscriptionHandler(
	database *db.Database,
	auth *auth.Auth,
) http.Handler {
	h := &SubscriptionHandler{
		Database: database,
		Auth:     auth,
	}

	r := chi.NewRouter()
	r.Post("/create", h.createCheckoutSession)

	return r
}

type createCheckoutSessionRequest struct {
	PriceId string `json:"priceId"`
}

func (h *SubscriptionHandler) createCheckoutSession(w http.ResponseWriter, r *http.Request) {
	// Grab their ID
	id, err := h.Auth.ValidateAccessToken(strings.TrimPrefix(r.Header.Get("Authorization"), "Bearer "))
	if err != nil {
		utils.WriteError(w, "invalid token", http.StatusUnauthorized, "invalid_token", nil)
		return
	}

	// Grab request data
	var data createCheckoutSessionRequest
	err = json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		log.Printf("error parsing login request JSON: %v", err)
		utils.WriteError(w, "error parsing JSON", http.StatusBadRequest, "invalid_json", nil)
		return
	}

	// Get the user object
	user, err := h.Database.FindUserByID(r.Context(), id)
	if err != nil {
		log.Printf("error finding user: %v", err)
		utils.WriteError(w, "error creating checkout session", http.StatusInternalServerError, "internal_error", nil)
		return
	}

	// Create checkout session
	params := &stripe.CheckoutSessionParams{
		Customer: stripe.String(user.StripeCustomerId),
		Mode:     stripe.String(stripe.CheckoutSessionModeSubscription),
		LineItems: []*stripe.CheckoutSessionLineItemParams{
			{
				Price:    stripe.String(data.PriceId),
				Quantity: stripe.Int64(1),
			},
		},
		SuccessURL: stripe.String("http://localhost:3000/success"),
		CancelURL:  stripe.String("http://localhost:3000/pricing"),
	}
	sess, err := session.New(params)
	if err != nil {
		log.Printf("error creating checkout session: %v", err)
		utils.WriteError(w, "error creating checkout session", http.StatusInternalServerError, "internal_error", nil)
	}

	// Return checkout url
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(map[string]string{
		"checkoutUrl": sess.URL,
	})
	if err != nil {
		log.Printf("error encoding JSON response: %v", err)
	}
}

func (h *SubscriptionHandler) stripeWebhook(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement
}
