package models

import "time"

type Subscription struct {
	Id                   int64     `json:"id"`
	UserId               int64     `json:"userId"`
	StripeSubscriptionId string    `json:"stripeSubscriptionId"`
	StripePriceId        string    `json:"stripePriceId"`
	PlanId               int64     `json:"planId"`
	Status               string    `json:"status"`
	CurrentPeriodStart   time.Time `json:"currentPeriodStart"`
	CurrentPeriodEnd     time.Time `json:"currentPeriodEnd"`
	CreatedAt            time.Time `json:"createdAt"`
	UpdatedAt            time.Time `json:"updatedAt"`
}

type CreateSubscription struct {
	UserId               int64     `json:"userId"`
	StripeSubscriptionId string    `json:"stripeSubscriptionId"`
	StripePriceId        string    `json:"stripePriceId"`
	PlanId               int64     `json:"planId"`
	Status               string    `json:"status"`
	CurrentPeriodStart   time.Time `json:"currentPeriodStart"`
	CurrentPeriodEnd     time.Time `json:"currentPeriodEnd"`
}

type TokenBalance struct {
	UserId       int64     `json:"userId"`
	TokenBalance int       `json:"tokenBalance"`
	PeriodStart  time.Time `json:"periodStart"`
	PeriodEnd    time.Time `json:"periodEnd"`
}
