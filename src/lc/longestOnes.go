package lc

import "sort"

func LongestOnes(A []int, K int) int {
	n := len(A)
	P := make([]int, n+1)
	ans := 0
	for i,v := range A {
		P[i+1] = P[i] + 1 - v
	}
	for right,v := range P {
		left := sort.SearchInts(P, v-K)
		ans = max(ans, right-left)
	}
	return ans
}
