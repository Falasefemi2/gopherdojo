// 044: Mutexes. A sync.Mutex makes a read-modify-write sequence on shared
// state atomic, which is what a map increment needs.
//
// Scenario: log lines arrive from many goroutines at the same time and a
// counter map must hold the exact number of occurrences of every word.
//
// Rules:
//  1. Spawn one goroutine per key and protect the map with a sync.Mutex.
//  2. Every read and write of the map happens while the mutex is held.
//  3. Join all goroutines before returning the map.
//  4. An empty or nil input returns a non-nil empty map.
//
// TODO: implement countWords. Keep main() as the demo harness.
package main

import "fmt"

// countWords counts how often each key appears, updating the map from one
// goroutine per key.
func countWords(keys []string) map[string]int {
	// TODO
	return nil
}

func main() {
	fmt.Println(countWords([]string{"go", "gopher", "go"}))
}
