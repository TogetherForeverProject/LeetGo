package problems

import (
    "fmt"
)

func climbStairs(n int) int {
    if n <= 2 {
        return n
    }

    dp := make([]int, n+1)
    dp[1] = 1
    dp[2] = 2

    for i := 3; i <= n; i++ {
        dp[i] = dp[i-1] + dp[i-2]
    }

    return dp[n]
}

func Problem70() {
    n := 5
    ways := climbStairs(n)
    fmt.Printf("Number of distinct ways to climb %d steps: %d\n", n, ways)
}

func init() {
    RegisterProblem(70, Problem70)
}
