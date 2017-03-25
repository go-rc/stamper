package handlers

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/tombell/stamper/services"
)

// GitHubHandler is the http.Handler for handling incoming GitHub webook
// requests.
func GitHubHandler(w http.ResponseWriter, r *http.Request) {
	var err error

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		// TODO
		panic(err)
	}

	event := r.Header.Get("X-GitHub-Event")

	err = services.HandleEvent(event, body)
	if err != nil {
		panic(err)
	}

	fmt.Fprintf(w, "OK")
}
