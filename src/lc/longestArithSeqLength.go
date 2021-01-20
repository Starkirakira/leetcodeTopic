package lc

import "math"

func LongestArithSeqLength(A []int) int {
	n := len(A)
	dp := make([][]int,n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int,20002)
		for j := 0; j < len(dp[0]); j++ {
			dp[i][j] = 1
		}
	}

	ans := 2
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			sub := A[i] - A[j]
			dp[i][sub + 10000] = int(math.Max(float64(dp[i][sub + 10000]), float64(dp[j][sub + 10000] + 1)))
			ans = int(math.Max(float64(ans), float64(dp[i][sub + 10000])))
		}
}

return ans

}
