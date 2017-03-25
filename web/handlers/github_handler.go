package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// GitHubHandler is the http.Handler for handling incoming GitHub webook
// requests.
func GitHubHandler(w http.ResponseWriter, r *http.Request) {
	event := r.Header.Get("X-GitHub-Event")

	fmt.Printf("event: %s\n\n", event)

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		// TODO
		panic(err)
	}

	var payload bytes.Buffer

	err = json.Indent(&payload, body, "", "  ")
	if err != nil {
		// TODO
		panic(err)
	}

	fmt.Println(payload.String())

	fmt.Fprintf(w, "OK")
}
