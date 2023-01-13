package middleware

import (
	"net/http"

	"golang.org/x/time/rate"
)

var api_limit = rate.NewLimiter(1, 10)

func LimitApiCalls(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !api_limit.Allow() {
			http.Error(w, http.StatusText(http.StatusTooManyRequests), http.StatusTooManyRequests)
			return
		} else {
			next.ServeHTTP(w, r)
		}
	})
}
