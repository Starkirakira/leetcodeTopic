package lc

import "math"

func FindRotateSteps(ring string , key string) int {
	const inf = math.MaxInt64 / 2

	n , m := len(ring), len(key)
	pos := [26][]int{}
	for i , c := range ring {
		pos[c - 'a'] = append(pos[c - 'a'] , i)
	}

	dp := make([][]int , m)
	for i := range dp {
		dp[i] = make([]int , n)
		for j := range dp[i] {
			dp[i][j] = inf
		}
	}

	for _ , p := range pos[key[0] - 'a'] {
		dp[0][p] = minRing(p , n - p) + 1
	}
	for i := 1; i < m; i++ {
		for _ , j := range pos[key[i] - 'a'] {
			for _ , k := range pos[key[i - 1] - 'a'] {
				dp[i][j] = minRing(dp[i][j] , dp[i-1][k] + minRing(absRing(j - k), n - absRing(j - k)) + 1)
			}
		}
	}
	return minRing(dp[m-1]...)
}

func minRing(a ...int) int {
	res := a[0]
	for _, v := range a[1:] {
		if v < res {
			res = v
		}
	}
	return res
}

func absRing(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
