package main

import (
	"fmt"
	"net/http"
)

type HomePage struct{}

func (h *HomePage) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "world"
	}
	fmt.Fprintf(w, "Hello, %s!", name)
}

func main() {

	http.Handle("/", &HomePage{})
	http.ListenAndServe(":3000", nil)
}
