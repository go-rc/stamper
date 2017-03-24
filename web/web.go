package web

import (
	"fmt"
	"log"
	"net/http"

	"github.com/tombell/stamper/web/handlers"
)

func Run(host, port string, l *log.Logger) error {
	http.Handle("/", http.HandlerFunc(handlers.RootHandler))
	http.Handle("/github/", http.HandlerFunc(handlers.GitHubHandler))

	addr := fmt.Sprintf("%s:%s", host, port)
	l.Printf("HTTP service listening on %s", addr)

	err := http.ListenAndServe(addr, nil)
	return err
}
