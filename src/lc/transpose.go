package lc

func Transpose(matrix [][]int) [][]int {
	n := len(matrix)
	m := len(matrix[0])
	trans := make([][]int,m)
	for i := range trans {
		trans[i] = make([]int,n)
		for j := range trans[i] {
			trans[i][j] = -1
		}
	}
	for i, row := range matrix {
		for j, v := range row {
			trans[j][i] = v
		}
	}
	return trans
}
