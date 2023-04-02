package lc

import (
	"math"
	"sort"
)

func FindCriticalAndPseudoCriticalEdges(n int, edges [][]int) [][]int {
	for i, e := range edges {
		edges[i] = append(e, i)
	}
	sort.Slice(edges, func(i, j int) bool {
		return edges[i][2] < edges[j][2]
	})

	calcMST := func(uf *unionFind1, ignoreID int) (mstValue int64) {
		for i, e := range edges {
			if i != ignoreID && uf.union(e[0], e[1]) {
				mstValue += int64(e[2])
			}
		}
		if uf.setCount > 1 {
			return math.MaxInt64
		}
		return
	}

	mstValue := calcMST(newUnionFind1(n), -1)

	var keyEdges, pseudokeyEdges []int
	for i, e := range edges {
		if calcMST(newUnionFind1(n), i) > mstValue {
			keyEdges = append(keyEdges, e[3])
			continue
		}

		uf := newUnionFind1(n)
		uf.union(e[0], e[1])
		if int64(e[2])+calcMST(uf, i) == mstValue {
			pseudokeyEdges = append(pseudokeyEdges, e[3])
		}
	}
	return [][]int{keyEdges, pseudokeyEdges}
}
