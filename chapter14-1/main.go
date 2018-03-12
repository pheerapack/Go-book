package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc(
		"/",
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintln(w, "Hello")
		})

	barHaandeler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Bar")
	}
	http.HandleFunc("/bar", barHaandeler)
	http.ListenAndServe(":3000", nil)
}
