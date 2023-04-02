package lc

import "math"

func MaxProfitWithoutTips(k int, prices []int) int64 {
	n := len(prices)
	if n == 0 {
		return 0
	}

	k = int(math.Min(float64(k), float64(n/2)))
	buy := make([][]int64, n)
	sell := make([][]int64, n)

	for i := range buy {
		buy[i] = make([]int64, k+1)
		sell[i] = make([]int64, k+1)
	}
	buy[0][0] = -int64(prices[0])
	for i := 1; i <= k; i++ {
		buy[0][i] = math.MinInt64
		sell[0][i] = math.MinInt64
	}

	for i := 1; i < n; i++ {
		buy[i][0] = int64(int(math.Max(float64(buy[i-1][0]), float64(sell[i-1][0]-int64(prices[i])))))
		for j := 1; j <= k; j++ {
			buy[i][j] = int64(int(math.Max(float64(buy[i-1][j]), float64(sell[i-1][j]-int64(prices[i])))))
			sell[i][j] = int64(int(math.Max(float64(sell[i-1][j]), float64(buy[i-1][j-1]+int64(prices[i])))))
		}
	}

	return maxP(sell[n-1]...)
}
