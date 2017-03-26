package web

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/tombell/stamper/services"
)

func GitHubHandler(w http.ResponseWriter, r *http.Request) {
	l := r.Context().Value(LoggerContextKey("logger")).(*log.Logger)

	srv := r.Context().Value(ServiceContextKey("GitHubService")).(*services.GitHubService)

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		l.Fatalf("unable to read request body")
	}

	event := r.Header.Get("X-GitHub-Event")

	err = srv.HandleEvent(event, body)
	if err != nil {
		l.Fatalf("unable to handle request body")
	}

	fmt.Fprintf(w, "OK")
}
