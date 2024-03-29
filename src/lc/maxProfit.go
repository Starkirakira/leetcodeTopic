package lc

func MaxProfit(prices []int, fee int) int {
	n := len(prices)
	dp := make([][2]int, n)
	dp[0][1] = -prices[0]
	for i := 1; i < n; i++ {
		dp[i][0] = maxPrices(dp[i-1][0], dp[i-1][1]+prices[i]-fee)
		dp[i][1] = maxPrices(dp[i-1][1], dp[i-1][0]-prices[i])
	}
	return dp[n-1][0]
}

func maxPrices(a, b int) int {
	if a > b {
		return a
	}

	return b

}