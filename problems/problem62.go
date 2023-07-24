// Unique Paths
//
// There is a robot on an m x n grid. The robot is initially located at the top-left corner (i.e., grid[0][0]).
// The robot tries to move to the bottom-right corner (i.e., grid[m - 1][n - 1]).
// The robot can only move either down or right at any point in time.
//
// Given the two integers m and n, return the number of possible unique paths that the robot can take to reach the bottom-right corner.
//
// The test cases are generated so that the answer will be less than or equal to 2 * 109.

package problems

import (
	"fmt"
)

func uniquePaths(m int, n int) int {
	// Create a 2D slice to store the number of unique paths for each cell
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}

	// Initialize the first row and first column with 1, as there is only one way to reach each cell in the first row and first column
	for i := 0; i < m; i++ {
		dp[i][0] = 1
	}
	for j := 0; j < n; j++ {
		dp[0][j] = 1
	}

	// Calculate the number of unique paths for each cell in a bottom-up manner
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}

	// The number of unique paths to reach the bottom-right corner (grid[m-1][n-1]) is stored in dp[m-1][n-1]
	return dp[m-1][n-1]
}

// Input: m = 3, n = 7
func Problem62() {
	m := 3
	n := 7
	fmt.Println("Number of unique paths:", uniquePaths(m, n))
}

func init() {
	RegisterProblem(62, Problem62)
}
