package db

import (
	"context"
	"fmt"
)

func (d *Database) FindPlanByStripeProductId(
	ctx context.Context,
	stripeProductId string,
) (int64, error) {
	return -1, fmt.Errorf("not implemented")
}
