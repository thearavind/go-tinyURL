package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestFindTinyURL(test *testing.T) {
	test.Run("Visit the short link of google.com", func(t *testing.T) {
		request, _ := http.NewRequest("GET", "/wWAOlq", nil)
		request.Header.Set("Content-Type", "application/json")
		response := httptest.NewRecorder()
		router.ServeHTTP(response, request)
		if response.Code != http.StatusFound {
			t.Errorf("Invalid response code: %d", response.Code)
		}
	})

	test.Run("Visit a wong short link", func(t *testing.T) {
		request, _ := http.NewRequest("GET", "/wWAOlq", nil)
		request.Header.Set("Content-Type", "application/json")
		response := httptest.NewRecorder()
		router.ServeHTTP(response, request)
		if response.Code != http.StatusFound {
			t.Errorf("Invalid response code: %d", response.Code)
		}
	})
}
