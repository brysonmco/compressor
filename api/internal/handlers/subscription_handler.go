package handlers

import (
	"encoding/json"
	"github.com/awesomebfm/compressor/internal/db"
	"github.com/awesomebfm/compressor/internal/middleware"
	"github.com/awesomebfm/compressor/internal/models"
	"github.com/awesomebfm/compressor/internal/utils"
	"github.com/go-chi/chi/v5"
	"github.com/stripe/stripe-go/v82"
	"github.com/stripe/stripe-go/v82/checkout/session"
	"github.com/stripe/stripe-go/v82/webhook"
	"io"
	"log"
	"net/http"
	"time"
)

type SubscriptionHandler struct {
	Database       *db.Database
	AuthMiddleware *middleware.AuthMiddleware
	EndpointSecret string
}

func NewSubscriptionHandler(
	database *db.Database,
	authMiddleware *middleware.AuthMiddleware,
	endpointSecret string,
) http.Handler {
	h := &SubscriptionHandler{
		Database:       database,
		AuthMiddleware: authMiddleware,
		EndpointSecret: endpointSecret,
	}

	r := chi.NewRouter()
	r.With(authMiddleware.Protected).Post("/checkout", h.handleCreateCheckoutSession)
	r.With(authMiddleware.Protected).Post("/cancel", h.handleCancelSubscription)
	r.Post("/webhook", h.handleStripeWebhook)

	return r
}

type createCheckoutSessionRequest struct {
	PriceId string `json:"priceId"`
}

func (h *SubscriptionHandler) handleCreateCheckoutSession(w http.ResponseWriter, r *http.Request) {
	// Grab their ID
	id := r.Context().Value("userId").(int64)

	// Grab request data
	var data createCheckoutSessionRequest
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		log.Printf("error parsing login request JSON: %v", err)
		utils.WriteError(w, r, http.StatusBadRequest, "error parsing JSON", "invalid_json", nil)
		return
	}

	// Get the user object
	user, err := h.Database.FindUserByID(r.Context(), id)
	if err != nil {
		log.Printf("error finding user: %v", err)
		utils.WriteError(w, r, http.StatusInternalServerError, "error creating checkout session", "internal_error", nil)
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
		SuccessURL: stripe.String("http://localhost:8080/dashboard"),
		CancelURL:  stripe.String("http://localhost:8080/pricing"),
	}
	sess, err := session.New(params)
	if err != nil {
		log.Printf("error creating checkout session: %v", err)
		utils.WriteError(w, r, http.StatusInternalServerError, "error creating checkout session", "internal_error", nil)
	}

	// Return checkout url
	utils.WriteSuccess(w, r, http.StatusOK, "checkout session created", map[string]string{
		"checkoutUrl": sess.URL,
	})
}

func (h *SubscriptionHandler) handleStripeWebhook(w http.ResponseWriter, r *http.Request) {
	const maxBodyBytes = int64(65536)
	r.Body = http.MaxBytesReader(w, r.Body, maxBodyBytes)
	payload, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("error reading webhook payload: %v", err)
		utils.WriteError(w, r, http.StatusBadRequest, "could not read webhook payload", "bad_payload", nil)
		return
	}

	evnt, err := webhook.ConstructEvent(payload, r.Header.Get("Stripe-Signature"), h.EndpointSecret)
	if err != nil {
		log.Printf("error constructing webhook event: %v", err)
		utils.WriteError(w, r, http.StatusBadRequest, "could not construct webhook event", "bad_payload", nil)
		return
	}

	switch evnt.Type {
	case "customer.subscription.created":
		h.handleSubscriptionCreated(w, r, evnt)
		break
	case "customer.subscription.updated":
	case "customer.subscription.deleted":
	default:
		log.Printf("unhandled webhook event type: %s", evnt.Type)
		utils.WriteError(w, r, http.StatusBadRequest, "unhandled webhook event type", "bad_payload", nil)
	}

}

func (h *SubscriptionHandler) handleCancelSubscription(w http.ResponseWriter, r *http.Request) {

}

func (h *SubscriptionHandler) handleSubscriptionCreated(w http.ResponseWriter, r *http.Request, evnt stripe.Event) {
	var subscription stripe.Subscription
	err := json.Unmarshal(evnt.Data.Raw, &subscription)
	if err != nil {
		log.Printf("error unmarshalling webhook event: %v", err)
		utils.WriteError(w, r, http.StatusBadRequest, "could not unmarshall webhook event", "bad_payload", nil)
		return
	}

	customerId := subscription.Customer.ID
	subscriptionId := subscription.ID
	priceId := subscription.Items.Data[0].Price.ID

	user, err := h.Database.FindUserByStripeCustomerID(r.Context(), customerId)
	if err != nil {
		log.Printf("error finding user by stripe customer id: %v", err)
		utils.WriteError(w, r, http.StatusInternalServerError, "could not handle subscription creation", "internal_error", nil)
		return
	}

	// Get the plan
	productId := subscription.Items.Data[0].Price.Product.ID
	plan, err := h.Database.FindPlanByStripeProductId(r.Context(), productId)
	if err != nil {
		log.Printf("error finding plan by product id %v: %v", productId, err)
		utils.WriteError(w, r, http.StatusInternalServerError, "could not handle subscription creation", "internal_error", nil)
		return
	}

	subReq := models.CreateSubscription{
		UserId:               user.Id,
		StripeSubscriptionId: subscriptionId,
		StripePriceId:        priceId,
		PlanId:               plan.Id,
		Status:               string(subscription.Status),
		CurrentPeriodStart:   time.Unix(subscription.Items.Data[0].CurrentPeriodStart, 0),
		CurrentPeriodEnd:     time.Unix(subscription.Items.Data[0].CurrentPeriodEnd, 0),
	}

	_, err = h.Database.CreateSubscription(r.Context(), subReq)
	if err != nil {
		log.Printf("error creating subscription in database: %v", err)
		utils.WriteError(w, r, http.StatusInternalServerError, "could not handle subscription creation", "internal_error", nil)
		return
	}
}
