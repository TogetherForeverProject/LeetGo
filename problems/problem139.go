// Word Break
//
// Given a string s and a dictionary of strings wordDict, return true if s can be segmented into a space-separated sequence of one or more dictionary words.
//
// Note that the same word in the dictionary may be reused multiple times in the segmentation.

package problems

import (
    "fmt"
)

func wordBreak(s string, wordDict []string) bool {
	// Create a set to store the words from the wordDict for efficient lookup
	wordSet := make(map[string]bool)
	for _, word := range wordDict {
		wordSet[word] = true
	}

	n := len(s)
	// Create a boolean slice dp to store the segmentation result for each substring
	dp := make([]bool, n+1)
	dp[0] = true // An empty string can be segmented into words from the wordDict

	for i := 1; i <= n; i++ {
		for j := 0; j < i; j++ {
			// Check if the substring s[j:i] can be segmented and the word s[j:i] is in the wordDict
			if dp[j] && wordSet[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}

	return dp[n]
}

// Input: s = "leetcode"
func Problem139() {
    s := "leetcode"
    wordDict := []string{"leet", "code"}
    fmt.Println("Can be segmented:", wordBreak(s, wordDict))
}

func init() {
    RegisterProblem(139, Problem139)
}
