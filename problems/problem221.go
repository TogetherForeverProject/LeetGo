// Maximal Square
//
// Given an m x n binary matrix filled with 0's and 1's, find the largest square containing only 1's and return its area.

package problems

import (
    "fmt"
)

func maximalSquare(matrix [][]byte) int {
	m := len(matrix)
	n := len(matrix[0])

	// Create a 2D slice to store the side length of the largest square ending at each cell
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}

	maxSide := 0

	// Initialize the first row and first column of dp
	for i := 0; i < m; i++ {
		dp[i][0] = int(matrix[i][0] - '0')
		maxSide = max(maxSide, dp[i][0])
	}
	for j := 0; j < n; j++ {
		dp[0][j] = int(matrix[0][j] - '0')
		maxSide = max(maxSide, dp[0][j])
	}

	// Calculate the side length of the largest square ending at each cell in a bottom-up manner
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][j] == '1' {
				dp[i][j] = mins(dp[i-1][j-1], dp[i][j-1], dp[i-1][j]) + 1
				maxSide = max(maxSide, dp[i][j])
			}
		}
	}

	// The area of the largest square is maxSide^2
	return maxSide * maxSide
}

// Input: matrix = [["1","0","1","0","0"],["1","0","1","1","1"],["1","1","1","1","1"],["1","0","0","1","0"]]
func Problem221() {
	matrix := [][]byte{
		{'1', '0', '1', '0', '0'},
		{'1', '0', '1', '1', '1'},
		{'1', '1', '1', '1', '1'},
		{'1', '0', '0', '1', '0'},
	}

	fmt.Println("Area of the largest square containing only 1's:", maximalSquare(matrix))
}

func init() {
    RegisterProblem(221,Problem221)
}
