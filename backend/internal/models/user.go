package models

type User struct {
	Id           int64  `json:"id"`
	Email        string `json:"email"`
	FirstName    string `json:"firstName"`
	LastName     string `json:"lastName"`
	PasswordHash string `json:"passwordHash"`
}
