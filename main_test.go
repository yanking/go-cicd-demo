package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthEndpoint(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	rr := httptest.NewRecorder()

	healthHandler(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("expected 200, got %d", rr.Code)
	}

	var resp Response
	if err := json.NewDecoder(rr.Body).Decode(&resp); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if resp.Status != "ok" {
		t.Errorf("expected status ok, got %s", resp.Status)
	}
	if resp.Message != "alive" {
		t.Errorf("expected message alive, got %s", resp.Message)
	}
}

func TestApiEndpoint(t *testing.T) {
	tests := []struct {
		name     string
		query    string
		wantMsg  string
	}{
		{"with name", "?name=Alice", "Hello, Alice!"},
		{"without name", "", "Hello, World!"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/api"+tt.query, nil)
			rr := httptest.NewRecorder()

			apiHandler(rr, req)

			if rr.Code != http.StatusOK {
				t.Errorf("expected 200, got %d", rr.Code)
			}

			var resp Response
			if err := json.NewDecoder(rr.Body).Decode(&resp); err != nil {
				t.Fatalf("failed to decode response: %v", err)
			}

			if resp.Message != tt.wantMsg {
				t.Errorf("expected %s, got %s", tt.wantMsg, resp.Message)
			}
		})
	}
}
