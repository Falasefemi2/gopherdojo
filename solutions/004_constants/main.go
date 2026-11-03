// 004: Constants. Constants are fixed at compile time. An iota block numbers
// a list, and a typed constant carries its type into every expression.
//
// Scenario: a file-upload service checks permissions (read/write/execute as
// bit flags) and refuses anything larger than its 5 MiB upload limit.
//
// Rules:
//  1. Read, Write and Execute are bit flags worth 1, 2 and 4, declared in one
//     iota block.
//  2. KB is an untyped constant worth 1024; MaxUploadBytes is a typed int64
//     constant worth 5 * KB * KB.
//  3. hasPermission reports whether every bit in perm is set in mask.
//  4. permissionString renders three characters in Read, Write, Execute order,
//     using 'r', 'w', 'x' and '-' for an unset bit.
//  5. fitsUpload accepts sizes from 0 through MaxUploadBytes inclusive and
//     rejects anything negative.
package main

import "fmt"

// Permission bits.
const (
	Read = 1 << iota
	Write
	Execute
)

// KB is one kibibyte.
const KB = 1024

// MaxUploadBytes is the service upload limit.
const MaxUploadBytes int64 = 5 * KB * KB

// hasPermission reports whether perm is fully contained in mask.
func hasPermission(mask, perm int) bool {
	return mask&perm == perm
}

// permissionString renders mask as three permission characters.
func permissionString(mask int) string {
	chars := []byte{'-', '-', '-'}
	if mask&Read != 0 {
		chars[0] = 'r'
	}
	if mask&Write != 0 {
		chars[1] = 'w'
	}
	if mask&Execute != 0 {
		chars[2] = 'x'
	}
	return string(chars)
}

// fitsUpload reports whether size is within the upload limit.
func fitsUpload(size int64) bool {
	return size >= 0 && size <= MaxUploadBytes
}

func main() {
	fmt.Println(permissionString(Read | Write))
	fmt.Println(hasPermission(Read|Write, Write))
	fmt.Println(fitsUpload(MaxUploadBytes))
}
