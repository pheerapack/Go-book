package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHomePageHandler(t *testing.T) {
	//Test function ทำตัว respond ให้กับ handler
	res := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	HomePageHandler(res, req)
	//Struct ชื่อฟังก์ชั่น

	if res.Code != 200 {
		t.Fatalf("Expect status to = 200, but got %d", res.Code)
	}

}

func TestExpectHelloWorld(t *testing.T) {
	//Test function ทำตัว respond ให้กับ handler
	res := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	HomePageHandler(res, req)
	//Struct ชื่อฟังก์ชั่น

	//เรียก body เอา string
	if res.Body.String() != "Hello world" {
		t.Fatalf("Expect status %s", res.Code)
	}

}
