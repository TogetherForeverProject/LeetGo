// Distinct Subsequences
//
// Given two strings s and t, return the number of distinct subsequences of s which equals t.
//
// The test cases are generated so that the answer fits on a 32-bit signed integer.

package problems

import (
    "fmt"
)

func numDistinct(s string, t string) int {
	m, n := len(s), len(t)

	// Create a 2D dp array with one extra row and column for base cases
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	// Initialize the dp array with base cases
	for i := 0; i <= m; i++ {
		dp[i][0] = 1 // Empty string t is a subsequence of any string s
	}

	// Fill the dp array using the recurrence relation
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if s[i-1] == t[j-1] {
				dp[i][j] = dp[i-1][j-1] + dp[i-1][j]
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}

	return dp[m][n]
}

// Input: s1 = "rabbbit", t1 = "rabbit"
func Problem115() {
	s1 := "rabbbit"
	t1 := "rabbit"
	fmt.Printf("Number of distinct subsequences of %q that equal %q: %d\n", s1, t1, numDistinct(s1, t1))
}

func init() {
    RegisterProblem(115, Problem115)
}
