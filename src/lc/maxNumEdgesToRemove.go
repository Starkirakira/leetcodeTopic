package lc



func MaxNumEdgesToRemove(n int, edges [][]int) int {
	ans := len(edges)
	alice, bob := newUnionFind1(n), newUnionFind1(n)
	for _, e := range edges {
		x, y := e[1] - 1, e[2] - 1
		if e[0] == 3 && (!alice.isSameSet(x, y) || !bob.isSameSet(x, y)) {
			alice.union(x, y)
			bob.union(x, y)
			ans--
		}
	}
	uf := [2] *unionFind1{alice, bob}
	for _, e := range edges {
		if tp := e[0]; tp < 3 && uf[tp - 1].union(e[1] - 1, e[2] - 1) {
			ans--
		}
	}
	if alice.setCount > 1 || bob.setCount > 1 {
		return -1
	}
	return ans
}
