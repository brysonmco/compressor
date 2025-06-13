package subscriptions

import (
	"github.com/brysonmco/compressor/internal/models"
	"github.com/stripe/stripe-go/v82"
	"github.com/stripe/stripe-go/v82/customer"
	"github.com/stripe/stripe-go/v82/price"
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

func GetProductIdFromPrice(priceId string) (string, error) {
	p, err := price.Get(priceId, nil)
	if err != nil {
		return "", err
	}

	return p.Product.ID, nil
}
