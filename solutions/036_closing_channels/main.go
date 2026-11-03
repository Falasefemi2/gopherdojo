// 036: Closing Channels. Closing means "no more sends": receivers drain what
// is buffered and then get the zero value with ok == false.
//
// Scenario: a log shipper finishes a batch and must signal end-of-stream to the
// writer.
//
// Rules:
//  1. closeAfter sends vals in order, then closes the channel exactly once.
//  2. readAll receives until ok is false and returns the values it saw.
//  3. After a channel is closed, every receive returns the zero value and
//     ok == false.
//  4. Only the sender closes, and only after the last send.
package main

import "fmt"

// closeAfter returns a buffered channel holding vals, already closed.
func closeAfter(vals []int) <-chan int {
	ch := make(chan int, len(vals))
	for _, v := range vals {
		ch <- v
	}
	close(ch)
	return ch
}

// readAll drains ch until it reports closed.
func readAll(ch <-chan int) []int {
	out := make([]int, 0)
	for {
		v, ok := <-ch
		if !ok {
			return out
		}
		out = append(out, v)
	}
}

func main() {
	for v := range closeAfter([]int{1, 2, 3}) {
		fmt.Println(v)
	}
	fmt.Println(readAll(closeAfter([]int{4, 5})))
}
