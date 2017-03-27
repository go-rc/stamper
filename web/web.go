package web

import (
	"fmt"
	"log"
	"net/http"

	"github.com/mulib/middleware"

	"github.com/tombell/stamper/services"
)

func Run(host, port string, l *log.Logger) error {
	key := ServiceContextKey("GitHubService")

	http.Handle("/", http.HandlerFunc(RootHandler))
	http.Handle("/github",
		middleware.Use(
			http.HandlerFunc(GitHubHandler),
			WithService(key, services.Service),
			WithLogger(l),
		),
	)

	addr := fmt.Sprintf("%s:%s", host, port)
	l.Printf("HTTP service listening on %s", addr)

	err := http.ListenAndServe(addr, nil)
	return err
}
