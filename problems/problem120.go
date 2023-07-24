// Triangle
//
// Given a triangle array, return the minimum path sum from top to bottom.
//
// For each step, you may move to an adjacent number of the row below.
// More formally, if you are on index i on the current row, you may move to either index i or index i + 1 on the next row.

package problems

import (
    "fmt"
)

func minimumTotal(triangle [][]int) int {
	n := len(triangle)

	// Start from the second-last row and update the minimum path sum for each cell
	for i := n - 2; i >= 0; i-- {
		for j := 0; j <= i; j++ {
			triangle[i][j] = triangle[i][j] + min(triangle[i+1][j], triangle[i+1][j+1])
		}
	}

	// The minimum path sum will be stored in the top cell of the triangle
	return triangle[0][0]
}

func Problem120() {
    triangle := [][]int{
        {2},
        {3, 4},
        {6, 5, 7},
        {4, 1, 8, 3},
    }

    fmt.Println("Minimum path sum from top to bottom:", minimumTotal(triangle))
}

func init() {
    RegisterProblem(120, Problem120)
}
