// 079: Reference solution. Read this only after your own attempt passes.
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// healthResponse is the JSON shape of /health.
type healthResponse struct {
	Status string `json:"status"`
	Name   string `json:"name"`
}

func newHandler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", health)
	return mux
}

func health(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	name := r.URL.Query().Get("name")
	if name == "" {
		name = "world"
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(healthResponse{Status: "ok", Name: name})
}

func main() {
	fmt.Println("listening on http://127.0.0.1:8080")
	log.Fatal(http.ListenAndServe("127.0.0.1:8080", newHandler()))
}
