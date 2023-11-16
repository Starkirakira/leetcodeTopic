package lc

func FindSmallestSetOfVertices(n int, edges [][]int) []int {
	//需要找到所有入度为0的点集
	// edges表示从x->y，这样的情况下，y的入度就一定不是0
	tmp := make(map[int]struct{})
	var ans []int
	for _, edge := range edges {
		if _, ok := tmp[edge[1]]; !ok {
			tmp[edge[1]] = struct{}{}
		}
	}

	for i := 0; i < n; i++ {
		if _, ok := tmp[i]; !ok {
			ans = append(ans, i)
		}
	}
	return ans
}
