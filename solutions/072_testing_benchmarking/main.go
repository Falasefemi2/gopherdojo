// 072: Reference solution. Read this only after your own attempt passes.
package main

import "fmt"

func SumSquares(n int) int {
	sum := 0
	for i := 0; i < n; i++ {
		sum += i * i
	}
	return sum
}

func main() {
	fmt.Println(SumSquares(4), SumSquares(10))
}
