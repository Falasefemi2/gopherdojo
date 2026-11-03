package main

import (
	"strings"
	"testing"
)

func TestRenderAccount(t *testing.T) {
	tests := []struct {
		name string
		a    Account
		want string
	}{
		{
			name: "admin account",
			a:    Account{Name: "Ada", Admin: true},
			want: "Hello Ada!\nAdmin access enabled",
		},
		{
			name: "read only account",
			a:    Account{Name: "Grace Hopper", Admin: false},
			want: "Hello Grace Hopper!\nRead-only access",
		},
		{
			name: "empty name",
			a:    Account{Name: "", Admin: false},
			want: "Hello !\nRead-only access",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := renderAccount(tt.a)
			if err != nil {
				t.Fatalf("renderAccount(%+v) returned error: %v", tt.a, err)
			}
			if got != tt.want {
				t.Errorf("renderAccount(%+v) = %q, want %q", tt.a, got, tt.want)
			}
			if strings.HasSuffix(got, "\n") {
				t.Errorf("renderAccount(%+v) must not end with a newline", tt.a)
			}
		})
	}
}
