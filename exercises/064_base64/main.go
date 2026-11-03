// 064: Base64 Encoding. base64.StdEncoding uses "+" and "/", while
// base64.URLEncoding uses "-" and "_" so the text can live in a URL.
//
// Scenario: a webhook signs a short binary token and must send it both in a
// JSON field (standard alphabet) and in a callback URL (URL alphabet).
//
// Rules:
//  1. encodeStd and encodeURL return padded base64 for the bytes of s.
//  2. decodeStd and decodeURL return the original bytes as a string.
//  3. Invalid or wrongly-encoded input returns "" and a non-nil error.
//  4. The two alphabets differ only where "+" and "/" would appear.
//
// TODO: implement the four functions below. Keep main() as the demo harness.
package main

import (
	"fmt"
)

// encodeStd encodes s with the standard base64 alphabet.
func encodeStd(s string) string {
	// TODO
	return ""
}

// decodeStd decodes standard base64 text.
func decodeStd(s string) (string, error) {
	// TODO
	return "", nil
}

// encodeURL encodes s with the URL-safe base64 alphabet.
func encodeURL(s string) string {
	// TODO
	return ""
}

// decodeURL decodes URL-safe base64 text.
func decodeURL(s string) (string, error) {
	// TODO
	return "", nil
}

func main() {
	token := "\xfb\xff"
	fmt.Println("std:", encodeStd(token))
	fmt.Println("url:", encodeURL(token))

	back, err := decodeStd(encodeStd("hello"))
	fmt.Println("back:", back, "err:", err)
}
