// 078: Reference solution. Read this only after your own attempt passes.
package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
)

func fetch(client *http.Client, url string) (int, string, error) {
	resp, err := client.Get(url)
	if err != nil {
		return 0, "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, "", err
	}
	return resp.StatusCode, string(body), nil
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
