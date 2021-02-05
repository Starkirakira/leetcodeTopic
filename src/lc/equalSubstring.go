package lc

import "sort"

func EqualSubstring(s string, t string, maxCost int) (maxLen int) {
	n := len(s)
	accDiff := make([]int, n+1)
	for i, ch := range s {
		accDiff[i + 1] = accDiff[i] + abs(int(ch) - int(t[i]))
	}
	for i := 1; i <= n; i++ {
		start := sort.SearchInts(accDiff[:i] ,accDiff[i] - maxCost)
		maxLen = max(maxLen, i - start)
	}
	return
}
