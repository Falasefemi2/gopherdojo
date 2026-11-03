// 063: SHA256 Hashes. sha256.Sum256 returns a [32]byte array; %x formats it as
// lowercase hex, which is the usual way to display a digest.
//
// Scenario: an artifact registry records a checksum next to every uploaded
// tarball so downloads can be verified byte for byte.
//
// Rules:
//  1. sha256Hex returns lowercase hex of SHA-256 over the UTF-8 bytes of s.
//  2. Output is always exactly 64 characters; not 32 and not uppercase.
//  3. The empty string still has a hash; "" is not a special case.
//  4. Equal input gives equal output, different input gives different output.
package main

import (
	"crypto/sha256"
	"fmt"
)

// sha256Hex returns the hex-encoded SHA-256 digest of s.
func sha256Hex(s string) string {
	sum := sha256.Sum256([]byte(s))
	return fmt.Sprintf("%x", sum)
}

func main() {
	fmt.Println(sha256Hex("hello"))
	fmt.Println(sha256Hex(""))
}
