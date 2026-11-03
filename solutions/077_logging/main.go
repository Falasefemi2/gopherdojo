// 077: Reference solution. Read this only after your own attempt passes.
package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func logInfo(w io.Writer, msg string) {
	log.New(w, "", 0).Print("INFO: " + msg)
}

func logError(w io.Writer, msg string) {
	log.New(w, "", 0).Print("ERROR: " + msg)
}

func main() {
	fmt.Println("events:")
	logInfo(os.Stdout, "worker started")
	logError(os.Stdout, "queue empty")
}
