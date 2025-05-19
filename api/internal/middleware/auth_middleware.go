package middleware

import (
	"context"
	"errors"
	"github.com/awesomebfm/compressor/internal/auth"
	"github.com/awesomebfm/compressor/internal/db"
	"github.com/awesomebfm/compressor/internal/utils"
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
		id, err := m.Auth.ValidateAccessToken(strings.TrimPrefix(r.Header.Get("Authorization"), "Bearer "))
		if err != nil && errors.Is(err, errors.New("expired_token")) {
			utils.WriteError(w, "token has expired", http.StatusUnauthorized, "expired_token", nil)
			return
		} else if err != nil {
			utils.WriteError(w, "invalid token", http.StatusUnauthorized, "invalid_token", nil)
			return
		}

		ctx := context.WithValue(r.Context(), "userID", id)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
