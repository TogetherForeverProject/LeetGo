// Fibonacci Number
//
// The Fibonacci numbers, commonly denoted F(n) form a sequence, called the Fibonacci sequence,
// such that each number is the sum of the two preceding ones, starting from 0 and 1. That is,
// F(0) = 0, F(1) = 1
// F(n) = F(n - 1) + F(n - 2), for n > 1.
// Given n, calculate F(n).

package problems

import (
    "fmt"
)

func fibonacciMemo(n int, memo map[int]int) int {
	if n <= 0 {
		return 0
	} else if n == 1 {
		return 1
	}

	if val, ok := memo[n]; ok {
		return val
	}

	memo[n] = fibonacciMemo(n-1, memo) + fibonacciMemo(n-2, memo)
	return memo[n]
}

func fib(n int) int {
	memo := make(map[int]int)
	return fibonacciMemo(n, memo)
}

func Problem509() {
    n := 10
	result := fib(n)
	fmt.Printf("Fibonacci number at index %d is: %d\n", n, result)
}

func init() {
    RegisterProblem(509, Problem509)
}
