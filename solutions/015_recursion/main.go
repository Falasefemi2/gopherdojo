// 015: Recursion. A recursive function needs a base case and a smaller
// subproblem; without them it never stops.
//
// Scenario: a billing report needs factorials, Fibonacci numbers, and the
// digit sum of invoice numbers, all computed without loops.
//
// Rules:
//  1. factorial(n) is 1 for n == 0 and n * factorial(n-1) for n > 0.
//  2. A negative factorial input returns 0, the "invalid" marker.
//  3. fib(n) returns F(0)=0, F(1)=1, and F(n)=F(n-1)+F(n-2) for n > 1.
//  4. sumDigits(n) sums decimal digits recursively and uses the absolute
//     value, so sumDigits(-99) is 18.
package main

import "fmt"

func factorial(n int) int {
	if n < 0 {
		return 0
	}
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func fib(n int) int {
	if n < 0 {
		return 0
	}
	if n < 2 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

func sumDigits(n int) int {
	if n < 0 {
		n = -n
	}
	if n < 10 {
		return n
	}
	return n%10 + sumDigits(n/10)
}

func main() {
	fmt.Println("5! =", factorial(5))
	fmt.Println("fib(10) =", fib(10))
	fmt.Println("sumDigits(1234) =", sumDigits(1234))
	fmt.Println("sumDigits(-99) =", sumDigits(-99))
}
