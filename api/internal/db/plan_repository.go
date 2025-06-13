package db

import (
	"context"
	"fmt"
	"github.com/brysonmco/compressor/internal/models"
)

func (d *Database) FindAllPlans(
	ctx context.Context,
) ([]*models.Plan, error) {
	return []*models.Plan{}, fmt.Errorf("not implemented")
}

func (d *Database) FindPlanByStripeProductId(
	ctx context.Context,
	stripeProductId string,
) (*models.Plan, error) {
	return &models.Plan{}, fmt.Errorf("not implemented")
}

func (d *Database) FindPlanById(
	ctx context.Context,
	id int64,
) (*models.Plan, error) {
	return &models.Plan{}, fmt.Errorf("not implemented")
}

func (d *Database) FindPlanByName(
	ctx context.Context,
	name string,
) (*models.Plan, error) {
	return &models.Plan{}, fmt.Errorf("not implemented")
}
