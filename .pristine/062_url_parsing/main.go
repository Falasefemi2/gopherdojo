// 062: URL Parsing. net/url splits a raw URL into Scheme, Host, Path and a
// Query map; Query().Get returns "" for a missing key.
//
// Scenario: an analytics worker receives campaign links and needs the landing
// path plus the utm_source tag, tolerating links that omit the tag entirely.
//
// Rules:
//  1. parseURL returns a non-nil *url.URL for valid input, else a non-nil error.
//  2. queryParam parses raw and returns Query().Get(key), or "" on any failure.
//  3. A missing or empty query parameter both yield the empty string.
//  4. Query values are percent-decoded; repeated keys return the first value.
//
// TODO: implement parseURL and queryParam below. Keep main() as the demo.
package main

import (
	"fmt"
	"net/url"
)

// parseURL parses raw into a URL.
func parseURL(raw string) (*url.URL, error) {
	// TODO
	return nil, nil
}

// queryParam returns the first value of key in raw's query string.
func queryParam(raw, key string) string {
	// TODO
	return ""
}

func main() {
	u, err := parseURL("https://example.com/search?q=go&page=2")
	if err != nil {
		fmt.Println("parse error:", err)
		return
	}
	fmt.Println("host:", u.Host, "path:", u.Path)
	fmt.Println("q:", queryParam("https://example.com/search?q=go&page=2", "q"))
}
