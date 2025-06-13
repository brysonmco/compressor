package middleware

import (
	"context"
	"errors"
	"github.com/brysonmco/compressor/internal/auth"
	"github.com/brysonmco/compressor/internal/db"
	"github.com/brysonmco/compressor/internal/utils"
	"net/http"
	"strings"
)

type AuthMiddleware struct {
	Auth     *auth.Auth
	Database *db.Database
}

func NewAuthMiddleware(auth *auth.Auth, database *db.Database) *AuthMiddleware {
	return &AuthMiddleware{
		Auth:     auth,
		Database: database,
	}
}

func (m *AuthMiddleware) Protected(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Validate token
		id, err := m.Auth.ValidateAccessToken(strings.TrimPrefix(r.Header.Get("Authorization"), "Bearer "))
		if err != nil && errors.Is(err, errors.New("expired_token")) {
			utils.WriteError(w, r, http.StatusUnauthorized, "token has expired", "expired_token", nil)
			return
		} else if err != nil {
			utils.WriteError(w, r, http.StatusUnauthorized, "invalid token", "invalid_token", nil)
			return
		}

		// Ensure their email is valid
		user, err := m.Database.FindUserByID(r.Context(), id)
		if err != nil {
			utils.WriteError(w, r, http.StatusInternalServerError, "error fetching user", "user_not_found", nil)
		}
		if !user.EmailVerified {
			utils.WriteError(w, r, http.StatusUnauthorized, "email not verified", "email_not_verified", nil)
			return
		}

		ctx := context.WithValue(r.Context(), "userId", id)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (m *AuthMiddleware) ProtectedWithoutEmailVerification(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Validate token
		id, err := m.Auth.ValidateAccessToken(strings.TrimPrefix(r.Header.Get("Authorization"), "Bearer "))
		if err != nil && errors.Is(err, errors.New("expired_token")) {
			utils.WriteError(w, r, http.StatusUnauthorized, "token has expired", "expired_token", nil)
			return
		} else if err != nil {
			utils.WriteError(w, r, http.StatusUnauthorized, "invalid token", "invalid_token", nil)
			return
		}

		ctx := context.WithValue(r.Context(), "userId", id)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
