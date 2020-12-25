package lc

import "sort"

func FindContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	n := len(g)
	m := len(s)
	ans := 0
	for i,j := 0,0; i < n && j < m; i++ {
		for j < m && g[i] > s[j] {
			j++
		}
		if j < m {
			ans++
			j++
		}
	}
	return ans
}
