package lc

import "math"

func FindRotateSteps(ring string, key string) int64 {
	const inf = math.MaxInt64 / 2

	n, m := len(ring), len(key)
	pos := [26][]int{}
	for i, c := range ring {
		pos[c-'a'] = append(pos[c-'a'], i)
	}

	dp := make([][]int64, m)
	for i := range dp {
		dp[i] = make([]int64, n)
		for j := range dp[i] {
			dp[i][j] = inf
		}
	}

	for _, p := range pos[key[0]-'a'] {
		dp[0][p] = minRing(int64(p), int64(n-p)) + 1
	}
	for i := 1; i < m; i++ {
		for _, j := range pos[key[i]-'a'] {
			for _, k := range pos[key[i-1]-'a'] {
				dp[i][j] = minRing(dp[i][j], dp[i-1][k]+minRing(int64(absRing(j-k)), int64(n-absRing(j-k))+1))
			}
		}
	}
	return int64(minRing(dp[m-1]...))
}

func minRing(a ...int64) int64 {
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
