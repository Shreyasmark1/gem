package middleware

import (
	"context"
	"gem/pkg/response"
	"gem/pkg/utils"
	"net/http"
	"strings"
)

type contextKey string

const SessionContextKey contextKey = "session"

type Session struct {
	UserID string
}

func JWTMiddleware(jwtSecret string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				next.ServeHTTP(w, r)
				return
			}

			parts := strings.Split(authHeader, " ")
			if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
				next.ServeHTTP(w, r)
				return
			}

			tokenString := parts[1]
			claims, err := utils.ValidateJWT(tokenString, jwtSecret)
			if err != nil || claims == nil {
				next.ServeHTTP(w, r)
				return
			}

			if sub, ok := (*claims)["sub"].(string); ok && sub != "" {
				session := &Session{UserID: sub}
				ctx := context.WithValue(r.Context(), SessionContextKey, session)

				r = r.WithContext(ctx)
			}

			next.ServeHTTP(w, r)
		})
	}
}

func RequireAuth() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			session, ok := r.Context().Value(SessionContextKey).(*Session)
			if !ok || session == nil || session.UserID == "" {
				_ = response.Unauthorized("Unauthorized")
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
