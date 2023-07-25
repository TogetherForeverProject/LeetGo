// Number of Longest Increasing Subsequence
//
// Given an integer array nums, return the number of longest increasing subsequences.
//
// Notice that the sequence has to be strictly increasing.

package problems

import (
    "fmt"
)

func findNumberOfLIS(nums []int) int {
	n := len(nums)

	// Create two arrays to store the lengths and counts of increasing subsequences
	lengths := make([]int, n)
	counts := make([]int, n)

	// Initialize the lengths and counts arrays with base cases
	for i := 0; i < n; i++ {
		lengths[i] = 1
		counts[i] = 1
	}

	maxLength := 1
	totalCount := 1

	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				if lengths[i] == lengths[j]+1 {
					counts[i] += counts[j]
				} else if lengths[i] < lengths[j]+1 {
					lengths[i] = lengths[j] + 1
					counts[i] = counts[j]
				}
			}
		}

		if lengths[i] > maxLength {
			maxLength = lengths[i]
			totalCount = counts[i]
		} else if lengths[i] == maxLength {
			totalCount += counts[i]
		}
	}

	return totalCount
}

// Input: nums = [1, 3, 5, 4, 7]
func Problem673() {
	nums := []int{1, 3, 5, 4, 7}
	fmt.Println("Number of longest increasing subsequences:", findNumberOfLIS(nums))
}

func init() {
    RegisterProblem(673, Problem673)
}
