package web

import (
	"fmt"
	"log"
	"net/http"

	mw "github.com/mulib/middleware"

	"github.com/tombell/stamper/services"
	"github.com/tombell/stamper/web/handlers"
	"github.com/tombell/stamper/web/middleware"
)

// Run sets up the http.Handlers and binds the server to a host/port.
func Run(host, port string, l *log.Logger) error {
	http.Handle("/", http.HandlerFunc(handlers.RootHandler))
	http.Handle("/github/",
		mw.Use(
			http.HandlerFunc(handlers.GitHubHandler),
			middleware.WithGitHubService(services.Service),
		),
	)

	addr := fmt.Sprintf("%s:%s", host, port)
	l.Printf("HTTP service listening on %s", addr)

	err := http.ListenAndServe(addr, nil)
	return err
}
