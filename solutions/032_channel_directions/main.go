// 032: Channel Directions. A chan<- parameter can only send; a <-chan
// parameter can only receive. The compiler enforces the direction.
//
// Scenario: an order producer feeds a fulfillment consumer through one queue.
//
// Rules:
//  1. produce sends 1..n to out in order, then closes out. n <= 0 closes
//     immediately without sending.
//  2. consume receives every value from in until it is closed.
//  3. sumProduced wires produce to consume through one channel and returns the
//     sum of 1..n.
//  4. No function may leave a goroutine blocked on a channel.
package main

import "fmt"

// produce sends 1..n on out and closes it.
func produce(n int, out chan<- int) {
	for i := 1; i <= n; i++ {
		out <- i
	}
	close(out)
}

// consume reads every value from in until in is closed.
func consume(in <-chan int) []int {
	out := make([]int, 0)
	for v := range in {
		out = append(out, v)
	}
	return out
}

// sumProduced builds a channel, runs produce in a goroutine, and returns the
// sum of everything consume reads.
func sumProduced(n int) int {
	ch := make(chan int)
	go produce(n, ch)

	total := 0
	for _, v := range consume(ch) {
		total += v
	}
	return total
}

func main() {
	ch := make(chan int, 5)
	produce(5, ch)
	fmt.Println(consume(ch))
	fmt.Println(sumProduced(5))
}
