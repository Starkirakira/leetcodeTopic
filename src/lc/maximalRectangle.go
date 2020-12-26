package lc

import "math"

func MaximalRectangle(matrix [][]byte) (ans int) {
	if len(matrix )== 0 {
		return
	}
	m, n := len(matrix), len(matrix[0])
	left := make([][]int, m)

	for i, row := range matrix {
		left[i] = make([]int, n)
		for j,v := range row {
			if v == '0' {
				continue
			}
			if j == 0 {
				left[i][j] = 1
			} else {
				left[i][j] = left[i][j-1] + 1
			}
		}
	}

	for i,row := range matrix {
		for j,v := range row {
			if v == '0' {
				continue
			}
			width := left[i][j]
			area := width
			for k := i - 1; k >= 0; k-- {
				width = int(math.Min(float64(width),float64(left[k][j])))
				area = int(math.Max(float64(area),float64((i-k+1)*width)))
			}
			ans = int(math.Max(float64(ans),float64(area)))
		}
	}
	return

}
