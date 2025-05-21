package db

import (
	"context"
	"github.com/awesomebfm/compressor/internal/models"
)

func (d *Database) FindSubscriptionById(
	ctx context.Context,
	id int64,
) (*models.Subscription, error) {
	query := `SELECT id, user_id, stripe_subscription_id, stripe_price_id, plan_id, status, current_period_start, 
       current_period_end, created_at, updated_at 
		FROM subscriptions
		WHERE id = $1`

	row := d.Pool.QueryRow(ctx, query, id)

	var subscription models.Subscription
	if err := row.Scan(
		&subscription.Id,
		&subscription.UserId,
		&subscription.StripeSubscriptionId,
		&subscription.StripePriceId,
		&subscription.PlanId,
		&subscription.Status,
		&subscription.CurrentPeriodStart,
		&subscription.CurrentPeriodEnd,
		&subscription.CreatedAt,
		&subscription.UpdatedAt,
	); err != nil {
		return nil, err
	}

	return &subscription, nil
}

func (d *Database) FindActiveSubscriptionByUserId(
	ctx context.Context,
	userId int64,
) (*models.Subscription, error) {
	query := `SELECT id, user_id, stripe_subscription_id, stripe_price_id, plan_id, status, current_period_start, 
	   current_period_end, created_at, updated_at 
		FROM subscriptions
		WHERE user_id = $1 AND status = 'active'`

	row := d.Pool.QueryRow(ctx, query, userId)

	var subscription models.Subscription
	if err := row.Scan(
		&subscription.Id,
		&subscription.UserId,
		&subscription.StripeSubscriptionId,
		&subscription.StripePriceId,
		&subscription.PlanId,
		&subscription.Status,
		&subscription.CurrentPeriodStart,
		&subscription.CurrentPeriodEnd,
		&subscription.CreatedAt,
		&subscription.UpdatedAt,
	); err != nil {
		return nil, err
	}

	return &subscription, nil
}

func (d *Database) CreateSubscription(
	ctx context.Context,
	subscriptionReq models.CreateSubscription,
) (*models.Subscription, error) {
	query := `INSERT INTO subscriptions (user_id, stripe_subscription_id, stripe_price_id, plan_id, status, 
                           current_period_start, current_period_end) 
			VALUES ($1, $2, $3, $4, $5, $6)
    		RETURNING id, user_id, stripe_subscription_id, stripe_price_id, plan_id, status, current_period_start, 
    		    current_period_end, created_at, updated_at`

	var subscription models.Subscription
	if err := d.Pool.QueryRow(ctx, query,
		subscriptionReq.UserId,
		subscriptionReq.StripeSubscriptionId,
		subscriptionReq.StripePriceId,
		subscriptionReq.PlanId,
		subscriptionReq.Status,
		subscriptionReq.CurrentPeriodStart,
		subscriptionReq.CurrentPeriodEnd,
	).Scan(
		&subscription.Id,
		&subscription.UserId,
		&subscription.StripeSubscriptionId,
		&subscription.StripePriceId,
		&subscription.PlanId,
		&subscription.Status,
		&subscription.CurrentPeriodStart,
		&subscription.CurrentPeriodEnd,
		&subscription.CreatedAt,
		&subscription.UpdatedAt,
	); err != nil {
		return nil, err
	}

	return &subscription, nil
}
