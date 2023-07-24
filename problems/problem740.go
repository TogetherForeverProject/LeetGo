// Delete And Earn
//
// You are given an integer array nums.
// You want to maximize the number of points you get by performing the following operation any number of times:
//
//     Pick any nums[i] and delete it to earn nums[i] points.
//     Afterwards, you must delete every element equal to nums[i] - 1 and every element equal to nums[i] + 1.
//
// Return the maximum number of points you can earn by applying the above operation some number of times.

package problems

import (
    "fmt"
)

func deleteAndEarn(vc []int) int {
    dp := make([]int, 10003)
    for _, v := range vc {
        dp[v] += v
    }
    for i := 3; i <= 10001; i++ {
        dp[i] = max(dp[i - 1], max(dp[i - 2], dp[i - 3]) + dp[i])
    }
    return dp[10001]
}

func Problem740() {
	nums := []int{3, 4, 2}
	maxPoints := deleteAndEarn(nums)
	fmt.Println("Maximum number of points earned:", maxPoints)
}

func init() {
    RegisterProblem(740, Problem740)
}
