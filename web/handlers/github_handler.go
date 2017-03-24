package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func GitHubHandler(w http.ResponseWriter, r *http.Request) {
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
