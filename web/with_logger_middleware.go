package web

import (
	"context"
	"log"
	"net/http"

	"github.com/mulib/middleware"
)

type LoggerContextKey string

func WithLogger(l *log.Logger) middleware.Middleware {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := context.WithValue(r.Context(), LoggerContextKey("logger"), l)
			h.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
