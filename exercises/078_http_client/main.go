// 078: HTTP Client. http.Client.Get performs a request; the response body is a
// stream that must be closed, and a non-2xx status is still a response.
//
// Scenario: a smoke check asks a status page for its health text.
//
// Rules:
//  1. fetch sends GET to url with the provided client and returns the status
//     code and the full body as a string.
//  2. A 4xx or 5xx status is NOT an error: status and body are returned as-is.
//  3. A transport or URL error returns status 0, an empty body and a non-nil
//     error.
//  4. The response body is always closed before fetch returns.
//
// TODO: implement fetch below. Keep main() as the demo harness.
package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
)

// fetch returns the status code and body of a GET request.
func fetch(client *http.Client, url string) (int, string, error) {
	// TODO
	return 0, "", nil
}

func main() {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "hello, dojo")
	}))
	defer srv.Close()

	status, body, err := fetch(srv.Client(), srv.URL)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Printf("%d %s\n", status, body)
}
