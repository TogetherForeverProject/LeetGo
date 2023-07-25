// Minimum ASCII Delete Sum for Two Strings
//
// Given two strings s1 and s2, return the lowest ASCII sum of deleted characters to make two strings equal.

package problems

import (
    "fmt"
)

func minimumDeleteSum(s1 string, s2 string) int {
	l1, l2 := len(s1), len(s2)
	if l1 < l2 {
		l1, l2 = l2, l1
		s1, s2 = s2, s1
	}

	dp := [2][]int{}
	dp[0], dp[1] = make([]int, l2+1), make([]int, l2+1)

	for i := 0; i <= l1; i++ {
		dp[0], dp[1] = dp[1], dp[0]
		for j := 0; j <= l2; j++ {
			if i == 0 && j == 0 {
				continue
			}
			dp[1][j] = 1e9
			if i > 0 {
				dp[1][j] = min(dp[1][j], dp[0][j]+int(s1[i-1]))
			}
			if j > 0 {
				dp[1][j] = min(dp[1][j], dp[1][j-1]+int(s2[j-1]))
			}
			if i > 0 && j > 0 && s1[i-1] == s2[j-1] {
				dp[1][j] = min(dp[1][j], dp[0][j-1])
			}
		}
	}
	return dp[1][l2]
}

// Input: s1 = "sea", s2 = "eat"
func Problem712() {
    s1 := "sea"
    s2 := "eat"
    result := minimumDeleteSum(s1, s2)
    fmt.Println("Minimum ascii delete sum for two strings:", result)
}

func init() {
    RegisterProblem(712, Problem712)
}
