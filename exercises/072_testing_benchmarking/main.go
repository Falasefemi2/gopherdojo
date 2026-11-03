// 072: Testing and Benchmarking. Go tests are ordinary functions named TestXxx,
// and BenchmarkXxx functions measure the same code under test.
//
// Scenario: a dashboard needs the sum of the squares of the first n integers;
// the tests grade correctness while a benchmark watches speed.
//
// Rules:
//  1. SumSquares(n) returns 0*0 + 1*1 + ... + (n-1)*(n-1) for n > 0.
//  2. SumSquares(0) and any negative n return 0.
//  3. The work lives in SumSquares, not in the test or the benchmark.
//  4. The benchmark loops b.N times over a fixed input.
//
// TODO: implement SumSquares below. Keep main() as the demo harness.
package main

import "fmt"

// SumSquares returns the sum of i*i for i in [0, n).
func SumSquares(n int) int {
	// TODO
	return 0
}

func main() {
	fmt.Println(SumSquares(4), SumSquares(10))
}
