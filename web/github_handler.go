package web

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/tombell/stamper/services"
)

// GitHubHandler is the http.Handler for incoming GitHub webhook requests.
func GitHubHandler(w http.ResponseWriter, r *http.Request) {
	key := ServiceContextKey("GitHubService")

	srv, ok := r.Context().Value(key).(*services.GitHubService)
	if !ok {
		// TODO
		panic("could not get github service from request context")
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		// TODO
		panic(err)
	}

	event := r.Header.Get("X-GitHub-Event")

	err = srv.HandleEvent(event, body)
	if err != nil {
		// TODO
		panic(err)
	}

	fmt.Fprintf(w, "OK")
}
