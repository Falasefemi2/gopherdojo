package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func testServer(t *testing.T) *httptest.Server {
	t.Helper()

	mux := http.NewServeMux()
	mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "hello, dojo")
	})
	mux.HandleFunc("/missing", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "not here")
	})

	srv := httptest.NewServer(mux)
	t.Cleanup(srv.Close)
	return srv
}

func TestFetchOK(t *testing.T) {
	srv := testServer(t)

	status, body, err := fetch(srv.Client(), srv.URL+"/hello")
	if err != nil {
		t.Fatalf("fetch(/hello) unexpected error: %v", err)
	}
	if status != http.StatusOK {
		t.Errorf("fetch(/hello) status = %d, want %d", status, http.StatusOK)
	}
	if body != "hello, dojo" {
		t.Errorf("fetch(/hello) body = %q, want %q", body, "hello, dojo")
	}
}

func TestFetchNon200IsNotAnError(t *testing.T) {
	srv := testServer(t)

	status, body, err := fetch(srv.Client(), srv.URL+"/missing")
	if err != nil {
		t.Fatalf("fetch(/missing) unexpected error: %v", err)
	}
	if status != http.StatusNotFound {
		t.Errorf("fetch(/missing) status = %d, want %d", status, http.StatusNotFound)
	}
	if body != "not here" {
		t.Errorf("fetch(/missing) body = %q, want %q", body, "not here")
	}
}

func TestFetchBadURL(t *testing.T) {
	srv := testServer(t)

	status, body, err := fetch(srv.Client(), "http://[::1")
	if err == nil {
		t.Fatal("fetch(bad url) error = nil, want an error")
	}
	if status != 0 {
		t.Errorf("fetch(bad url) status = %d, want 0", status)
	}
	if body != "" {
		t.Errorf("fetch(bad url) body = %q, want empty", body)
	}
}
