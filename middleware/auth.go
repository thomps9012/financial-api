package middleware

import (
	"context"
	"net/http"
)

var userCtxKey = &contextKey{"user"}

type contextKey struct {
	name string
}

type contextInfo struct {
	id   string `json:"id" bson:"role"`
	role string `json:"role" bson:"role"`
}

func Middleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			token, tokenErr := r.Cookie("Authorization")
			if tokenErr != nil {
				next.ServeHTTP(w, r)
				return
			}
			tokenString := token.Value
			tokenParsed, err := ParseToken(tokenString)
			if err != nil {
				http.Error(w, "Invalid token", http.StatusForbidden)
				return
			}
			id := tokenParsed["id"].(string)
			role := tokenParsed["role"].(string)
			contextInfo := &contextInfo{id, role}
			ctx := context.WithValue(r.Context(), userCtxKey, contextInfo)
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}

func ForID(ctx context.Context) string {
	raw, _ := ctx.Value(userCtxKey).(*contextInfo)
	return raw.id
}

func ForRole(ctx context.Context) string {
	raw, _ := ctx.Value(userCtxKey).(*contextInfo)
	return raw.role
}
