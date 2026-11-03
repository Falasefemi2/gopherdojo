package main

import "testing"

func TestEnvOr(t *testing.T) {
	t.Run("set", func(t *testing.T) {
		t.Setenv("DOJO_ENV", "production")
		if got := envOr("DOJO_ENV", "development"); got != "production" {
			t.Errorf("envOr(DOJO_ENV) = %q, want %q", got, "production")
		}
	})

	t.Run("empty falls back", func(t *testing.T) {
		t.Setenv("DOJO_ENV", "")
		if got := envOr("DOJO_ENV", "development"); got != "development" {
			t.Errorf("envOr(DOJO_ENV) = %q, want the fallback %q", got, "development")
		}
	})

	t.Run("unset falls back", func(t *testing.T) {
		if got := envOr("DOJO_NOT_SET_076", "fallback"); got != "fallback" {
			t.Errorf("envOr(DOJO_NOT_SET_076) = %q, want %q", got, "fallback")
		}
	})
}

func TestEnvInt(t *testing.T) {
	tests := []struct {
		name     string
		value    string
		fallback int
		want     int
	}{
		{"set", "9090", 8080, 9090},
		{"empty", "", 8080, 8080},
		{"not a number", "eighty", 8080, 8080},
		{"float", "80.5", 8080, 8080},
		{"negative", "-1", 8080, -1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Setenv("DOJO_PORT_076", tt.value)
			if got := envInt("DOJO_PORT_076", tt.fallback); got != tt.want {
				t.Errorf("envInt(DOJO_PORT_076 = %q, %d) = %d, want %d",
					tt.value, tt.fallback, got, tt.want)
			}
		})
	}

	t.Run("unset falls back", func(t *testing.T) {
		if got := envInt("DOJO_PORT_NOT_SET_076", 8080); got != 8080 {
			t.Errorf("envInt(DOJO_PORT_NOT_SET_076) = %d, want 8080", got)
		}
	})
}
