package middleware

import (
	"context"
	"net/http"
)

func Context(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.Background()

		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}
