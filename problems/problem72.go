// Edit Distance
//
// Given two strings word1 and word2, return the minimum number of operations required to convert word1 to word2.
//
// You have the following three operations permitted on a word:
//
//     Insert a character
//     Delete a character
//     Replace a character

package problems

import (
    "fmt"
)

func minDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)

	// Create a 2D array dp to store the minimum number of operations required
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	// Initialize the first row and column with the indexes (base cases)
	for i := 0; i <= m; i++ {
		dp[i][0] = i
	}
	for j := 0; j <= n; j++ {
		dp[0][j] = j
	}

	// Calculate the minimum number of operations required for each substring
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(min(dp[i-1][j], dp[i][j-1]), dp[i-1][j-1]) + 1
			}
		}
	}

	return dp[m][n]
}

// Input: word1 = "horse", word2 = "ros"
func Problem72() {
	word1 := "horse"
	word2 := "ros"
	fmt.Println("Minimum number of operations required:", minDistance(word1, word2))
}

func init() {
    RegisterProblem(72, Problem72)
}
