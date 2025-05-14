package models

import "time"

type Session struct {
	Id        int64     `json:"id"`
	TokenHash string    `json:"tokenHash"`
	UserId    int64     `json:"userId"`
	ExpiresAt time.Time `json:"expiresAt"`
	Revoked   bool      `json:"revoked"`
	CreatedAt time.Time `json:"createdAt"`
}

type CreateSession struct {
	TokenHash string `json:"tokenHash"`
	UserId    int64  `json:"userId"`
	ExpiresAt time.Time
	CreatedAt time.Time
}
