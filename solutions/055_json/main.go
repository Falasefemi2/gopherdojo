// 055: JSON. encoding/json maps exported struct fields to object keys through
// struct tags; omitempty drops zero values from the encoded output.
//
// Scenario: a user profile is stored as JSON and later loaded back, including
// records written by older versions that have no email field.
//
// Rules:
//  1. User fields encode as id, name, email, active, in that order.
//  2. email is omitted when empty; the other fields always appear.
//  3. encodeUser returns compact JSON with no extra whitespace.
//  4. decodeUser ignores unknown fields, leaves missing optional fields at
//     their zero value, and returns an error for malformed JSON.
package main

import (
	"encoding/json"
	"fmt"
)

// User is the profile record stored as JSON.
type User struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email,omitempty"`
	Active bool   `json:"active"`
}

// encodeUser marshals u to compact JSON.
func encodeUser(u User) (string, error) {
	data, err := json.Marshal(u)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// decodeUser unmarshals JSON into a User.
func decodeUser(data string) (User, error) {
	var u User
	err := json.Unmarshal([]byte(data), &u)
	return u, err
}

func main() {
	data, err := encodeUser(User{ID: 1, Name: "Ada", Email: "ada@example.com", Active: true})
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println(data)
}
