package middleware

import (
	"context"
	"net/http"
)

var userCtxKey = &contextKey{"user"}

type contextKey struct {
	name string
}

type Permission string

const (
	EMPLOYEE     Permission = "EMPLOYEE"
	MANAGER      Permission = "MANAGER"
	SUPERVISOR   Permission = "SUPERVISOR"
	EXECUTIVE    Permission = "EXECUTIVE"
	FINANCE_TEAM Permission = "FINANCE_TEAM"
)

type contextInfo struct {
	id          string
	name        string
	admin       bool
	permissions []Permission
}

func Middleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			header := r.Header.Get("Authorization")
			if header == "" {
				next.ServeHTTP(w, r)
				return
			}
			tokenString := header
			token, err := ParseToken(tokenString)
			if err != nil {
				http.Error(w, "Invalid token", http.StatusForbidden)
				return
			}
			id := token["id"].(string)
			name := token["name"].(string)
			admin := token["admin"].(bool)
			permissions := token["permissions"].([]Permission)
			contextInfo := &contextInfo{id, name, admin, permissions}
			ctx := context.WithValue(r.Context(), userCtxKey, contextInfo)
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}

func LoggedIn(ctx context.Context) bool {
	raw, _ := ctx.Value(userCtxKey).(*contextInfo)
	return raw.id != ""
}

func ForID(ctx context.Context) string {
	raw, _ := ctx.Value(userCtxKey).(*contextInfo)
	return raw.id
}

func ForName(ctx context.Context) string {
	raw, _ := ctx.Value(userCtxKey).(*contextInfo)
	return raw.name
}

func ForAdmin(ctx context.Context) bool {
	raw, _ := ctx.Value(userCtxKey).(*contextInfo)
	return raw.admin
}

func ForPermissions(ctx context.Context) []Permission {
	raw, _ := ctx.Value(userCtxKey).(*contextInfo)
	return raw.permissions
}
