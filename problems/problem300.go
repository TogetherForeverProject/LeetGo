// Longest Increasing Subsequence
//
// Given an integer array nums, return the length of the longest strictly increasing subsequence.

package problems

import (
    "fmt"
)

func lengthOfLIS(nums []int) int {
	tails := make([]int, len(nums))
	size := 0

	for _, num := range nums {
		left, right := 0, size
		for left < right {
			mid := left + (right-left)/2
			if tails[mid] < num {
				left = mid + 1
			} else {
				right = mid
			}
		}

		if left == size {
			tails[size] = num
			size++
		} else {
			tails[left] = num
		}
	}

	return size
}

// Input: nums = [10, 9, 2, 5, 3, 7, 101, 18]
func Problem300() {
    nums := []int{10, 9, 2, 5, 3, 7, 101, 18}
	fmt.Println("Length of the longest increasing subsequence:", lengthOfLIS(nums))
}

func init() {
    RegisterProblem(300, Problem300)
}
