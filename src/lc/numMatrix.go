package lc

type NumMatrix2110 struct {
	matrix [][]int
}

func Constructor2110(matrix [][]int) NumMatrix2110 {
	perSum := make([][]int, len(matrix)+1)
	for i := range perSum {
		perSum[i] = make([]int, len(matrix[0])+1)
	}
	for i := range matrix {
		for j := 0; j < len(matrix[0]); j++ {

			perSum[i+1][j+1] = matrix[i][j] + perSum[i+1][j] + perSum[i][j+1] - perSum[i][j]
		}
	}
	return NumMatrix2110{matrix: perSum}

}

func (this *NumMatrix2110) SumRegion(row1, col1, row2, col2 int) int {
	return this.matrix[row2+1][col2+1] - this.matrix[row1][col2+1] - this.matrix[row2+1][col1] + this.matrix[row1][col1]
}
