package middleware

import (
	"context"
	"net/http"

	mw "github.com/mulib/middleware"

	"github.com/tombell/stamper/services"
)

const GitHubServiceContextKey = contextKey("ctxGitHubService")

func WithGitHubService(srv *services.GitHubService) mw.Middleware {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := context.WithValue(r.Context(), GitHubServiceContextKey, srv)
			h.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
