package web

import (
	"context"
	"net/http"

	"github.com/mulib/middleware"
)

type ServiceContextKey string

func WithService(key ServiceContextKey, srv interface{}) middleware.Middleware {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := context.WithValue(r.Context(), key, srv)
			h.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
