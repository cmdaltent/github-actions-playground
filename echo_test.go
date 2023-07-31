package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestEchoHandler(t *testing.T) {
	// Create a new request with a POST method and a body of "hello, world"
	request, err := http.NewRequest(http.MethodPost, "http://localhost:8080", strings.NewReader("hello, world"))
	if err != nil {
		t.Fatal(err)
	}

	// Create a new response recorder
	responseRecorder := httptest.NewRecorder()

	// Call the handler function with the response recorder and the request
	handler(responseRecorder, request)

	// Get the response from the response recorder
	response := responseRecorder.Result()

	// Check that the response status code is 200 OK
	if response.StatusCode != http.StatusOK {
		t.Errorf("expected status OK; got %v", response.Status)
	}

	// Check that the response body is "hello, world"
	body, err := io.ReadAll(response.Body)
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		err := response.Body.Close()
		if err != nil {
			t.Fatal(err)
		}
	}()
	if string(body) != "hello, world" {
		t.Errorf("expected body hello, world; got %v", string(body))
	}
}
