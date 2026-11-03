// 079: HTTP Server. Handlers write a status, headers and a body; a
// http.ServeMux routes paths to handlers.
//
// Scenario: a health endpoint reports JSON and can personalise the name.
//
// Rules:
//  1. newHandler registers GET /health on a ServeMux; other paths return 404.
//  2. Any method other than GET returns 405 with an empty body.
//  3. The ?name= query value is used; missing or empty means "world".
//  4. A successful response is 200, Content-Type application/json and the body
//     {"status":"ok","name":"<name>"} followed by exactly one newline.
//
// TODO: implement newHandler and health below. Keep main() as the demo harness.
package main

import (
	"fmt"
	"log"
	"net/http"
)

// healthResponse is the JSON shape of /health.
type healthResponse struct {
	Status string `json:"status"`
	Name   string `json:"name"`
}

// newHandler builds the mux used by the server.
func newHandler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", health)
	return mux
}

// health writes the JSON health payload.
func health(w http.ResponseWriter, r *http.Request) {
	// TODO
}

func main() {
	fmt.Println("listening on http://127.0.0.1:8080")
	log.Fatal(http.ListenAndServe("127.0.0.1:8080", newHandler()))
}
