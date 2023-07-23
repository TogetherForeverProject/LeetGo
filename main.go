// Leet Code Dynamic Programming

package main

import (
	"fmt"
	"github.com/TogetherForeverProject/LeetGo/problems" // Import the problems package with dot notation
)

// Problem 70
// func main() {
//     n := 5
//     ways := problems.ClimbStairs(n)
//     fmt.Printf("Number of distinct ways to climb %d steps: %d\n", n, ways)
// }

// Problem 509
// func main() {
//     n := 10
// 	result := problems.Fib(n)
// 	fmt.Printf("Fibonacci number at index %d is: %d\n", n, result)
// }

// Problem 1137
// func main() {
// 	n := 25
// 	result := problems.Tribonacci(n)
// 	fmt.Printf("Tribonacci number at index %d is: %d\n", n, result)
// }

// Problem 746
// func main() {
// 	cost := []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}
// 	minCost := problems.MinCostClimbingStairs(cost) // Use the function directly without package name
// 	fmt.Println("Minimum cost to reach the top of the floor:", minCost)
// }

// Problem 198
func main() {
	nums := []int{2, 7, 9, 3, 1}
	maxAmount := problems.Rob(nums)
	fmt.Println("Maximum amount of money that can be robbed:", maxAmount)
}
