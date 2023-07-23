package main

import (
    "fmt"
)

// Problem 70
// func main() {
//     n := 5
//     ways := climbStairs(n)
//     fmt.Printf("Number of distinct ways to climb %d steps: %d\n", n, ways)
// }

// Problem 509
// func main() {
//     n := 10
// 	result := fib(n)
// 	fmt.Printf("Fibonacci number at index %d is: %d\n", n, result)
// }

// Problem 1137
func main() {
	n := 25
	result := tribonacci(n)
	fmt.Printf("Tribonacci number at index %d is: %d\n", n, result)
}
