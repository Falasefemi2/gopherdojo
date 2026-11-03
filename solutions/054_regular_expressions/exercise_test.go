package main

import "testing"

func TestIsKeyValue(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want bool
	}{
		{"typical pair", "user=alice", true},
		{"underscore key", "_private=1", true},
		{"digits in value", "port=8080", true},
		{"empty value", "user=", false},
		{"empty key", "=alice", false},
		{"numeric key", "1user=alice", false},
		{"trailing text", "user=alice extra", false},
		{"space in key", "user name=alice", false},
		{"two pairs", "a=1 b=2", false},
		{"empty string", "", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isKeyValue(tt.in); got != tt.want {
				t.Errorf("isKeyValue(%q) = %v, want %v", tt.in, got, tt.want)
			}
		})
	}
}

func TestParseKeyValue(t *testing.T) {
	tests := []struct {
		name      string
		in        string
		wantKey   string
		wantValue string
		wantOK    bool
	}{
		{"typical pair", "user=alice", "user", "alice", true},
		{"underscore key", "_token=abc123", "_token", "abc123", true},
		{"empty value", "user=", "", "", false},
		{"extra equals", "a=b=c", "", "", false},
		{"surrounding spaces", " user=alice", "", "", false},
		{"empty string", "", "", "", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			key, value, ok := parseKeyValue(tt.in)
			if key != tt.wantKey || value != tt.wantValue || ok != tt.wantOK {
				t.Errorf("parseKeyValue(%q) = (%q, %q, %v), want (%q, %q, %v)",
					tt.in, key, value, ok, tt.wantKey, tt.wantValue, tt.wantOK)
			}
		})
	}
}

func TestMaskValues(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want string
	}{
		{"one pair", "user=alice", "user=***"},
		{"two pairs", "user=alice pass=hunter2", "user=*** pass=***"},
		{"surrounding text", "login user=alice now", "login user=*** now"},
		{"query style", "a=1&b=2", "a=***&b=***"},
		{"no pairs", "nothing to mask", "nothing to mask"},
		{"empty string", "", ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maskValues(tt.in); got != tt.want {
				t.Errorf("maskValues(%q) = %q, want %q", tt.in, got, tt.want)
			}
		})
	}
}
