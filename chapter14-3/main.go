package main

import (
	"encoding/json"
	"net/http"
	"time"
)

type HomePage struct{}

func (h *HomePage) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	user := new(User)
	json.NewDecoder(r.Body).Decode(user)
	user.CreatedAt = time.Now()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	data, _ := json.Marshal(user)
	w.Write(data)
}

func main() {

	http.Handle("/", &HomePage{})
	http.ListenAndServe(":3000", nil)
}

type User struct {
	Firstname string    `json:"first_name"`
	Lastname  string    `json:"last_name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}
