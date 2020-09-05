package server

import (
	"fmt"
	"net/http"
)

func RegisterHandler(uc UseCases) {
	// Init handler struct

	// HTTP Handler
	http.HandleFunc("/", indexHandler)
}

// indexHandler responds to requests with our greeting.
func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	fmt.Fprint(w, "Hello, SMMOB!!! Visit us at https://github.com/kben19/smmob")
}
