package main

import (
	"testing"
)

func TestParseURL(t *testing.T) {
	raw := "https://user:secret@api.example.com:8443/v1/items?tag=go&tag=dojo&limit=10&empty="

	u, err := parseURL(raw)
	if err != nil {
		t.Fatalf("parseURL(%q) returned unexpected error: %v", raw, err)
	}
	if u == nil {
		t.Fatalf("parseURL(%q) returned a nil URL", raw)
	}

	if u.Scheme != "https" {
		t.Errorf("Scheme = %q, want %q", u.Scheme, "https")
	}
	if u.Host != "api.example.com:8443" {
		t.Errorf("Host = %q, want %q", u.Host, "api.example.com:8443")
	}
	if u.Path != "/v1/items" {
		t.Errorf("Path = %q, want %q", u.Path, "/v1/items")
	}
	if got := u.Query().Get("tag"); got != "go" {
		t.Errorf("Query().Get(tag) = %q, want %q", got, "go")
	}
	if got := u.Query().Get("limit"); got != "10" {
		t.Errorf("Query().Get(limit) = %q, want %q", got, "10")
	}
	if got := u.Query().Get("empty"); got != "" {
		t.Errorf("Query().Get(empty) = %q, want %q", got, "")
	}
	if got := u.Query().Get("missing"); got != "" {
		t.Errorf("Query().Get(missing) = %q, want %q", got, "")
	}

	if u.User == nil {
		t.Fatalf("User = nil, want user info")
	}
	if got := u.User.Username(); got != "user" {
		t.Errorf("Username = %q, want %q", got, "user")
	}
	if got, ok := u.User.Password(); !ok || got != "secret" {
		t.Errorf("Password = (%q, %v), want (%q, true)", got, ok, "secret")
	}
}

func TestParseURLWithoutQuery(t *testing.T) {
	u, err := parseURL("https://example.com/health")
	if err != nil {
		t.Fatalf("parseURL returned unexpected error: %v", err)
	}
	if u == nil {
		t.Fatal("parseURL returned a nil URL")
	}
	if u.Scheme != "https" || u.Host != "example.com" || u.Path != "/health" {
		t.Errorf("got scheme=%q host=%q path=%q, want https example.com /health", u.Scheme, u.Host, u.Path)
	}
	if got := u.Query().Get("utm_source"); got != "" {
		t.Errorf("Query().Get(utm_source) = %q, want %q", got, "")
	}
}

func TestParseURLErrors(t *testing.T) {
	bad := []string{"://bad", "http://[::1"}

	for _, raw := range bad {
		t.Run(raw, func(t *testing.T) {
			u, err := parseURL(raw)
			if err == nil {
				t.Errorf("parseURL(%q) error = nil, want an error", raw)
			}
			if u != nil {
				t.Errorf("parseURL(%q) = %v, want a nil URL on error", raw, u)
			}
		})
	}
}

func TestQueryParam(t *testing.T) {
	const raw = "https://example.com/search?q=hello%20world&tag=go&empty=&plus=a+b&tag=dojo"

	tests := []struct {
		name string
		key  string
		want string
	}{
		{"decoded space", "q", "hello world"},
		{"present", "tag", "go"},
		{"empty value", "empty", ""},
		{"missing key", "page", ""},
		{"plus decodes to space", "plus", "a b"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := queryParam(raw, tt.key); got != tt.want {
				t.Errorf("queryParam(raw, %q) = %q, want %q", tt.key, got, tt.want)
			}
		})
	}
}

func TestQueryParamBadURL(t *testing.T) {
	if got := queryParam("://bad", "q"); got != "" {
		t.Errorf("queryParam on bad URL = %q, want %q", got, "")
	}
}
