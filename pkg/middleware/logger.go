package middleware

import (
	"bytes"
	"io"
	"log/slog"
	"net/http"
	"time"
)

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		duration := time.Since(start)

		var bodyBytes []byte
		if r.Body != nil {
			bodyBytes, _ = io.ReadAll(r.Body)
		}

		slog.Info("request handled",
			"method", r.Method,
			"url", r.URL.Path,
			"body", string(bodyBytes),
			"duration", duration,
		)
		r.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
		next.ServeHTTP(w, r)
	})
}
