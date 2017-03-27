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
		// TODO
		l.Fatal(err)
	}

	event := r.Header.Get("X-GitHub-Event")

	err = srv.HandleEvent(event, body)
	if err != nil {
		// TODO
		l.Fatal(err)
	}

	fmt.Fprintf(w, "OK")
}
