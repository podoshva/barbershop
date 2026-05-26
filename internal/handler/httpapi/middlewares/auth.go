// Package middlewares
package middlewares

import (
	"context"
	"main/pkg"
	"net/http"
	"strings"
)

type contextKey string

const (
	ContextProfileID contextKey = "profile_id"
	ContextRole      contextKey = "role"
)

func Auth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		bearerToken := r.Header.Get("Authorization")
		authSlice := strings.Split(bearerToken, " ")
		if len(authSlice) != 2 {
			http.Error(w, "bearer token type invalid", http.StatusUnauthorized)
			return
		}
		claims, err := pkg.ParseToken(authSlice[1])
		if err != nil {
			http.Error(w, "parse token", http.StatusUnauthorized)
			return
		}
		ctx := context.WithValue(r.Context(), ContextProfileID, claims["profile_id"])
		ctx = context.WithValue(ctx, ContextRole, claims["role"])
		next(w, r.WithContext(ctx))
	}
}

func ProfileIDFromContext(ctx context.Context) string {
	v, _ := ctx.Value(ContextProfileID).(string)
	return v
}

func RoleFromContext(ctx context.Context) string {
	v, _ := ctx.Value(ContextProfileID).(string)
	return v
}
