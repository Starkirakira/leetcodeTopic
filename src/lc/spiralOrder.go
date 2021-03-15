package lc

func SpiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}

	X, Y := len(matrix), len(matrix[0])
	order := make([]int, X * Y)
	index := 0
	left, right, top, bottom := 0, Y - 1, 0 ,X - 1

	for left <= right && top <= bottom {
		for i := left; i <= right; i++ {
			order[index] = matrix[top][i]
			index++
		}
		for j := top + 1; j <= bottom; j++ {
			order[index] = matrix[j][right]
			index++
		}

		if left < right && top < bottom {
			for i := right - 1; i > left; i-- {
				order[index] = matrix[bottom][i]
				index++
			}
			for j := bottom; j > top; j-- {
				order[index] = matrix[j][left]
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
