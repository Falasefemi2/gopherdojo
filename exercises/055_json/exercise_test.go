package main

import (
	"reflect"
	"testing"
)

func TestEncodeUser(t *testing.T) {
	tests := []struct {
		name string
		u    User
		want string
	}{
		{
			name: "full record",
			u:    User{ID: 1, Name: "Ada", Email: "ada@example.com", Active: true},
			want: `{"id":1,"name":"Ada","email":"ada@example.com","active":true}`,
		},
		{
			name: "email omitted when empty",
			u:    User{ID: 2, Name: "Bob", Active: false},
			want: `{"id":2,"name":"Bob","active":false}`,
		},
		{
			name: "zero value",
			u:    User{},
			want: `{"id":0,"name":"","active":false}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := encodeUser(tt.u)
			if err != nil {
				t.Fatalf("encodeUser(%+v) returned error: %v", tt.u, err)
			}
			if got != tt.want {
				t.Errorf("encodeUser(%+v) = %s, want %s", tt.u, got, tt.want)
			}
		})
	}
}

func TestDecodeUser(t *testing.T) {
	tests := []struct {
		name string
		data string
		want User
	}{
		{
			name: "full record",
			data: `{"id":1,"name":"Ada","email":"ada@example.com","active":true}`,
			want: User{ID: 1, Name: "Ada", Email: "ada@example.com", Active: true},
		},
		{
			name: "missing optional email",
			data: `{"id":3,"name":"Cy","active":true}`,
			want: User{ID: 3, Name: "Cy", Active: true},
		},
		{
			name: "unknown field is ignored",
			data: `{"id":4,"name":"Dee","legacy":true,"active":false}`,
			want: User{ID: 4, Name: "Dee", Active: false},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := decodeUser(tt.data)
			if err != nil {
				t.Fatalf("decodeUser(%s) returned error: %v", tt.data, err)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("decodeUser(%s) = %+v, want %+v", tt.data, got, tt.want)
			}
		})
	}
}

func TestDecodeUserRejectsBadJSON(t *testing.T) {
	tests := []struct {
		name string
		data string
	}{
		{"truncated object", `{"id":`},
		{"wrong field type", `{"id":"not-a-number","name":"Eve"}`},
		{"trailing garbage", `{"id":5} extra`},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, err := decodeUser(tt.data); err == nil {
				t.Errorf("decodeUser(%s) returned nil error, want an error", tt.data)
			}
		})
	}
}
