package models

import "time"

type User struct {
	Id               int64     `json:"id"`
	Email            string    `json:"email"`
	FirstName        string    `json:"firstName"`
	LastName         string    `json:"lastName"`
	PasswordHash     string    `json:"passwordHash"`
	StripeCustomerId string    `json:"stripeCustomerId"`
	EmailVerified    bool      `json:"emailVerified"`
	CreatedAt        time.Time `json:"createdAt"`
	UpdatedAt        time.Time `json:"updatedAt"`
	LastLogin        time.Time `json:"lastLogin"`
}

type CreateUser struct {
	Email        string `json:"email"`
	FirstName    string `json:"firstName"`
	LastName     string `json:"lastName"`
	PasswordHash string `json:"passwordHash"`
}
