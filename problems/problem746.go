// Min Cost Climb Stairs
//
// You are given an integer array cost where cost[i] is the cost of ith step on a staircase.
// Once you pay the cost, you can either climb one or two steps.
//
// You can either start from the step with index 0, or the step with index 1.
//
// Return the minimum cost to reach the top of the floor.

package problems

import (
    "fmt"
)

func minCostClimbingStairs(cost []int) int {
	first, second := cost[0], cost[1]

	for i := 2; i < len(cost); i++ {
		first, second = second, min(first, second)+cost[i]
	}

	return min(first, second)
}

// Input: cost = [10,15,20]
func Problem746() {
	cost := []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}
	minCost := minCostClimbingStairs(cost) // Use the function directly without package name
	fmt.Println("Minimum cost to reach the top of the floor:", minCost)
}

func init() {
    RegisterProblem(746, Problem746)
}
