package lc

func isSimilar(s, t string) bool  {
	diff := 0
	for i := range s {
		if s[i] != t[i] {
			diff++
			if diff > 2 {
				return false
			}
		}
	}
	return true
}

func NumSimilarGroups(strs []string) int {
	n := len(strs)
	uf := newUnionFind1(n)
	for i, s := range strs {
		for j := i + 1; j < n; j++ {
			if !uf.isSameSet(i, j) && isSimilar(s, strs[j]) {
				uf.union(i ,j)
			}
		}
	}
	return uf.setCount
}
