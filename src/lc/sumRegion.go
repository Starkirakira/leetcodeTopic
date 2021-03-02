package lc

type NumMatrix struct {
	sums [][]int
}


func ConstructorRegion(matrix [][]int) NumMatrix {
	sums := make([][]int, len(matrix))
	for i,row := range matrix {
		sums[i] = make([]int, len(row) + 1)
		for j,v := range row {
			sums[i][j+1] = sums[i][j] + v
		}
	}
	return NumMatrix{sums}
}


func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	sum := 0
	for i := row1; i <= row2; i++ {
		sum += this.sums[i][col2+1] - this.sums[i][col1]
	}
	return sum
}
