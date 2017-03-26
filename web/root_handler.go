package web

import (
	"fmt"
	"net/http"
)

// RootHandler is the main http.Handler for the application.
func RootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Stamper 🏷")
}
