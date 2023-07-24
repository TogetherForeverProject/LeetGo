package problems

import (
    "fmt"
)

func rob(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	} else if n == 1 {
		return nums[0]
	}

	prev1 := nums[0]
	prev2 := max(nums[0], nums[1])

	for i := 2; i < n; i++ {
		current := max(prev2, prev1+nums[i])
		prev1, prev2 = prev2, current
	}

	return prev2
}

func Problem198() {
	nums := []int{2, 7, 9, 3, 1}
	maxAmount := rob(nums)
	fmt.Println("Maximum amount of money that can be robbed:", maxAmount)
}

func init() {
    RegisterProblem(198, Problem198)
}
