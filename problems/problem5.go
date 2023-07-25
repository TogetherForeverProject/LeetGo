// Longest Palindromic Substring
//
// Given a string s, return the longest palindromic substring in s.

package problems

import "fmt"

func longestPalindrome(s string) string {
    start, end := 0, 0

    for i := 0; i < len(s); i++ {
        // For odd length palindrome
        len1 := expandAroundCenter(s, i, i)
        // For even length palindrome
        len2 := expandAroundCenter(s, i, i + 1)

        maxLen := max(len1, len2)

        // Update the start and end indices of the longest palindrome if needed
        if maxLen > end - start {
            start = i - (maxLen - 1) / 2
            end = i + maxLen / 2
        }
    }

    return s[start : end + 1]
}

func expandAroundCenter(s string, left, right int) int {
    for left >= 0 && right < len(s) && s[left] == s[right] {
        left--
        right++
    }

    return right - left - 1
}

// Input: s = "babad"
func Problem5() {
    s := "babad"
    fmt.Println("Longest palindromic substring:", longestPalindrome(s))
}

func init() {
    RegisterProblem(5, Problem5)
}
