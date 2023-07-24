// Minimum Falling Path Sum
//
// Given an n x n array of integers matrix, return the minimum sum of any falling path through matrix.
//
// A falling path starts at any element in the first row and chooses the element in the next row that is either directly below or diagonally left/right.
// Specifically, the next element from position (row, col) will be (row + 1, col - 1), (row + 1, col), or (row + 1, col + 1).

package problems

import (
    "fmt"
)

func minFallingPathSum(matrix [][]int) int {
	n := len(matrix)

	// Start from the second row and update the minimum sum for each cell
	for i := 1; i < n; i++ {
		for j := 0; j < n; j++ {
			// For each cell, find the minimum sum of the current cell and the adjacent cells in the row above
			// Handle boundary cases for the first and last columns
			if j == 0 {
				matrix[i][j] += min(matrix[i-1][j], matrix[i-1][j+1])
			} else if j == n-1 {
				matrix[i][j] += min(matrix[i-1][j-1], matrix[i-1][j])
			} else {
				matrix[i][j] += min(matrix[i-1][j-1], min(matrix[i-1][j], matrix[i-1][j+1]))
			}
		}
	}

	// Find the minimum value in the last row, which represents the minimum sum of any falling path through the matrix
	minSum := matrix[n-1][0]
	for j := 1; j < n; j++ {
		minSum = min(minSum, matrix[n-1][j])
	}

	return minSum
}

func Problem931() {
	matrix := [][]int{
		{2, 1, 3},
		{6, 5, 4},
		{7, 8, 9},
	}

	fmt.Println("Minimum sum of any falling path:", minFallingPathSum(matrix))
}

func init() {
    RegisterProblem(931, Problem931)
}
