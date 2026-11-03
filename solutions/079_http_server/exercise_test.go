package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealth(t *testing.T) {
	tests := []struct {
		name   string
		target string
		want   string
	}{
		{"default name", "/health", `{"status":"ok","name":"world"}` + "\n"},
		{"query name", "/health?name=dojo", `{"status":"ok","name":"dojo"}` + "\n"},
		{"empty name", "/health?name=", `{"status":"ok","name":"world"}` + "\n"},
		{"unicode name", "/health?name=caf%C3%A9", `{"status":"ok","name":"café"}` + "\n"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, tt.target, nil)
			rec := httptest.NewRecorder()

			newHandler().ServeHTTP(rec, req)

			if rec.Code != http.StatusOK {
				t.Errorf("GET %s status = %d, want %d", tt.target, rec.Code, http.StatusOK)
			}
			if ct := rec.Header().Get("Content-Type"); ct != "application/json" {
				t.Errorf("GET %s Content-Type = %q, want %q", tt.target, ct, "application/json")
			}
			if got := rec.Body.String(); got != tt.want {
				t.Errorf("GET %s body = %q, want %q", tt.target, got, tt.want)
			}
		})
	}
}

func TestHealthRejectsNonGET(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/health", nil)
	rec := httptest.NewRecorder()

	newHandler().ServeHTTP(rec, req)

	if rec.Code != http.StatusMethodNotAllowed {
		t.Errorf("POST /health status = %d, want %d", rec.Code, http.StatusMethodNotAllowed)
	}
	if rec.Body.Len() != 0 {
		t.Errorf("POST /health body = %q, want empty", rec.Body.String())
	}
}

func TestUnknownPath(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/nope", nil)
	rec := httptest.NewRecorder()

	newHandler().ServeHTTP(rec, req)

	if rec.Code != http.StatusNotFound {
		t.Errorf("GET /nope status = %d, want %d", rec.Code, http.StatusNotFound)
	}
}
