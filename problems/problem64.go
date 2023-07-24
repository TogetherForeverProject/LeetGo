// Minimum Path Sum
//
// Given a m x n grid filled with non-negative numbers, find a path from top left to bottom right, which minimizes the sum of all numbers along its path.
//
// Note: You can only move either down or right at any point in time.

package problems

import (
    "fmt"
)

func minPathSum(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])

	// Create a 2D slice to store the minimum path sum for each cell
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}

	// Initialize the first row and first column of dp with the cumulative sums
	dp[0][0] = grid[0][0]
	for i := 1; i < m; i++ {
		dp[i][0] = dp[i-1][0] + grid[i][0]
	}
	for j := 1; j < n; j++ {
		dp[0][j] = dp[0][j-1] + grid[0][j]
	}

	// Calculate the minimum path sum for each cell in a bottom-up manner
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = grid[i][j] + min(dp[i-1][j], dp[i][j-1])
		}
	}

	// The minimum path sum to reach the bottom-right corner (grid[m-1][n-1]) is stored in dp[m-1][n-1]
	return dp[m-1][n-1]
}

func Problem64() {
    grid := [][]int{
        {1, 3, 1},
        {1, 5, 1},
        {4, 2, 1},
    }

    fmt.Println("Minimum path sum: ", minPathSum(grid))
}

func init() {
    RegisterProblem(64, Problem64)
}
