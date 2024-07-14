package middleware

import "net/http"

func Chain(h http.Handler, middlewares ...func(http.Handler) http.Handler) http.Handler {
	// ミドルウェアを逆順に適用
	for i := len(middlewares) - 1; i >= 0; i-- {
		h = middlewares[i](h)
	}
	return h
}
