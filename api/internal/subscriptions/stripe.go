package subscriptions

import (
	"github.com/awesomebfm/compressor/internal/models"
	"github.com/stripe/stripe-go/v82"
	"github.com/stripe/stripe-go/v82/customer"
	"strconv"
)

func CreateStripeCustomer(user *models.User) (*stripe.Customer, error) {
	params := &stripe.CustomerParams{
		Email: stripe.String(user.Email),
		Name:  stripe.String(user.FirstName + " " + user.LastName),
		Metadata: map[string]string{
			"user_id": strconv.FormatInt(user.Id, 10),
		},
	}
	return customer.New(params)
}
