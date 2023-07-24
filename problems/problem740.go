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
