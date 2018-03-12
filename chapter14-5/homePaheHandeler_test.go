package home

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestJsonHandler(t *testing.T) {
	//Test function ทำตัว respond ให้กับ handler
	res := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/json", strings.NewReader(
		`{"first_name": "John",
		"last_name": "Cena",
		"email": "john@test.com"}`))
	home := HomePageHandler{}
	home.ServeHTTP(res, req)

	if res.Code != 201 {
		t.Fatalf("Expect status to = 201, but got %d", res.Code)
	}

	user := new(User)
	json.NewDecoder(res.Body).Decode(user)

	if user.FirstName != "John" {
		t.Fatalf("Expect status to = 201, but got %s", user.FirstName)
	}

}
