package lc

func MinCostClimbingStairs(cost []int) int {
	n := len(cost)
	dp := make([]int, n+1)
	for i := 2; i <= n; i++ {
		dp[i] = minCost(dp[i-1]+cost[i-1], dp[i-2]+cost[i-2])
	}
	return dp[n]
}

func minCost(a, b int) int {
	if a < b {
		return a
	}
	return b
}

