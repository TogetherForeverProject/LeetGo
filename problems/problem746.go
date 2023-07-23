package problems

func MinCostClimbingStairs(cost []int) int {
	first, second := cost[0], cost[1]

	for i := 2; i < len(cost); i++ {
		first, second = second, min(first, second)+cost[i]
	}

	return min(first, second)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
