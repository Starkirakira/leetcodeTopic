package lc

func GenerateMatrix(n int) [][]int {
	order := make([][]int, n)
	for i := range order {
		order[i] = make([]int, n)
	}
	left := 0
	right := n - 1
	top := 0
	bottom := n - 1
	index := 1

	for left <= right && top <= bottom {
		for i := left; i <= right; i++ {
			order[top][i] = index
			index++
		}
		for i := top + 1; i <= bottom; i++ {
			order[i][right] = index
			index++
		}

		if left < right && top < bottom {
			for i := right - 1; i > left; i-- {
				order[bottom][i] = index
				index++
			}

			for i := bottom; i > top; i-- {
				order[i][left] = index
				index++
			}
		}

		left++
		right--
		top++
		bottom--

	}

	return order

}
