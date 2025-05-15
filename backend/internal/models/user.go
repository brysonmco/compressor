package models

type User struct {
	Id               int64  `json:"id"`
	Email            string `json:"email"`
	FirstName        string `json:"firstName"`
	LastName         string `json:"lastName"`
	PasswordHash     string `json:"passwordHash"`
	StripeCustomerId string `json:"stripeCustomerId"`
}

type CreateUser struct {
	Email        string `json:"email"`
	FirstName    string `json:"firstName"`
	LastName     string `json:"lastName"`
	PasswordHash string `json:"passwordHash"`
}

type UserClaims struct {
	Id        int64  `json:"id"`
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}
