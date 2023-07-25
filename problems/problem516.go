// Longest Palindromic Subsequence
//
// Given a string s, find the longest palindromic subsequence's length in s.
//
// A subsequence is a sequence that can be derived from another sequence by deleting some or no elements without changing the order of the remaining elements.

package problems

import (
    "fmt"
)

func longestPalindromeSubseq(s string) int {
	length := len(s)
	dp := make([][]int, length)
	for i := range dp {
		dp[i] = make([]int, length)
	}

	for start := length - 1; start >= 0; start-- {
		dp[start][start] = 1

		for end := start + 1; end < length; end++ {
			if s[start] == s[end] {
				dp[start][end] = dp[start+1][end-1] + 2
			} else {
				dp[start][end] = max(dp[start+1][end], dp[start][end-1])
			}
		}
	}

	return dp[0][length-1]
}

// Input: s = "bbbab"
func Problem516() {
	s := "bbbab"
	fmt.Println("Longest palindromic subsequence length:", longestPalindromeSubseq(s))
}

func init() {
    RegisterProblem(516,Problem516)
}
