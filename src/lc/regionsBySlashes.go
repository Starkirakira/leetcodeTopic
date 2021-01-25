package lc

func RegionsBySlashes(grid []string) int {
	N := len(grid)
	size := 4* N * N
	uf := unionFind2(size)
	for i := 0; i < N; i++ {
		row := []byte(grid[i])
		for j := 0; j < N; j++ {
			index := 4 * (i * N + j)
			c := row[j]
			if c == '/' {
				uf.union(index, index + 3)
				uf.union(index + 1, index + 2)
			} else if c == '\\' {
				uf.union(index, index + 1)
				uf.union(index + 2, index + 3)
			} else {
				uf.union(index, index + 1)
				uf.union(index + 1, index + 2)
				uf.union(index + 2, index + 3)
			}

			if j + 1 < N {
				uf.union(index + 1, 4*(i * N + j + 1) + 3)
			}
			if i + 1 < N {
				uf.union(index + 2, 4*((i + 1) * N + j))
			}
		}
	}
	return uf.setCount
}

func unionFind2(n int) *unionFind1 {
	parent := make([]int, n)
	size := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
		size[i] = 1
	}
	return &unionFind1{parent,size,n}

}
