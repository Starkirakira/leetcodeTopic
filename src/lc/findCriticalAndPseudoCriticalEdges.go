package lc

import (
	"math"
	"sort"
)

type unionFind1 struct {
	parent, size []int
	setCount int
}

func newUnionFind1(n int) *unionFind1 {
	parent := make([]int, n)
	size := make([]int, n)
	for i := range parent {
		parent[i] = i
		size[i] = 1
	}
	return &unionFind1{parent, size, n}
}

func (uf *unionFind1) find(x int) int {
	if uf.parent[x] != x {
		uf.parent[x] = uf.find(uf.parent[x])
	}
	return uf.parent[x]
}

func (uf *unionFind1) union(x,y int) bool {
	fx, fy := uf.find(x), uf.find(y)
	if fx == fy {
		return false
	}
	if uf.size[fx] < uf.size[fy] {
		fx, fy = fy, fx
	}
	uf.size[fx] += uf.size[fy]
	uf.parent[fy] = fx
	uf.setCount--
	return true
}

func FindCriticalAndPseudoCriticalEdges(n int, edges [][]int) [][]int {
	for i, e := range edges {
		edges[i] = append(e, i)
	}
	sort.Slice(edges, func(i, j int) bool {
		return edges[i][2] < edges[j][2]
	})

	calcMST := func(uf *unionFind1, ignoreID int)(mstValue int) {
		for i, e := range edges {
			if i != ignoreID && uf.union(e[0], e[1]) {
				mstValue += e[2]
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
		if e[2]+calcMST(uf, i) == mstValue {
			pseudokeyEdges = append(pseudokeyEdges, e[3])
		}
	}
	return [][]int{keyEdges, pseudokeyEdges}
}





