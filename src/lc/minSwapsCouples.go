package lc



func MinSwapsCouples(row []int) int {
	n := len(row)
	uf := newUnionFind1(n / 2)
	for i := 0; i < n; i += 2 {
		uf.union(row[i] / 2, row[i + 1] / 2)

	}
	return n / 2 - uf.setCount
}
