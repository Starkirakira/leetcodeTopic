package lc

import "math"

func MaxProfitWithoutTips(k int, prices []int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}

	k = int(math.Min(float64(k),float64(n / 2)))
	buy := make([][]int, n)
	sell := make([][]int, n)

	for i := range buy {
		buy[i] = make([]int, k + 1)
		sell[i] = make([]int, k + 1)
	}
	buy[0][0] = - prices[0]
	for i := 1; i <= k; i++ {
		buy[0][i] = math.MinInt64
		sell[0][i] = math.MinInt64
	}

	for i := 1; i < n; i++ {
		buy[i][0] = int(math.Max(float64(buy[i - 1][0]),float64(sell[i - 1][0] - prices[i])))
		for j := 1; j <= k; j++ {
			buy[i][j] = int(math.Max(float64(buy[i - 1][j]),float64(sell[i - 1][j] - prices[i])))
			sell[i][j] = int(math.Max(float64(sell[i - 1][j]),float64(buy[i - 1][j - 1] + prices[i])))
		}
	}

	return maxP(sell[n - 1]...)
}
func maxP(a ...int) int {
	res := a[0]
	for _, v := range a[1:] {
		if v > res {
			res = v
		}
	}
	return res
}
