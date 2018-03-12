package main

import (
	"fmt"
	"net/http"
)

type HomePage struct{}

func (h *HomePage) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello HOME")
}

func main() {
	http.HandleFunc(
		"/",
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintln(w, "Hello INDEX")
		})

	barHaandeler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello BAR")
	}
	http.HandleFunc("/bar", barHaandeler)
	http.Handle("/home", &HomePage{})
	http.ListenAndServe(":3000", nil)
}
