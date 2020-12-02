package main



func movesToChessboard(board [][]int) int {
	n := len(board)
	row := board[0]

	sameColCount := 0

	errorColCount := 0

	for i, num := range row {
		if num == row[0] {
			sameColCount++

			if i%2 == 1 {
				errorColCount++
			}
		} else {

			if i%2 == 0 {
				errorColCount++
			}
		}
	}

	if n%2 == 0 {
		if sameColCount != n/2 {
			return -1
		}
	} else {

		if sameColCount != n/2 && sameColCount != n/2+1 {
			return -1
		}
	}

	sameRowCount := 0

	errorRowCount := 0
	for i, line := range board {
		if isSameSlice(row, line) {
			sameRowCount++

			if i%2 == 1 {
				errorRowCount++
			}
		} else if isOppoSlice(row, line) {

			if i%2 == 0 {
				errorRowCount++
			}
		} else {
			return -1
		}
	}

	if n%2 == 0 {
		if sameRowCount != n/2 {
			return -1
		}
	} else {

		if sameRowCount != n/2 && sameRowCount != n/2+1 {
			return -1
		}
	}
	if n%2 == 0 {

		if n-errorColCount < errorColCount {
			errorColCount = n - errorColCount
		}
		if n-errorRowCount < errorRowCount {
			errorRowCount = n - errorRowCount
		}
	} else {

		if sameColCount*2 < n {
			errorColCount = n - errorColCount
		}


		if sameRowCount*2 < n {
			errorRowCount = n - errorRowCount
		}
	}
	return (errorColCount + errorRowCount) / 2
}

func isSameSlice(line1 []int, line2 []int) bool {
	for i, _ := range line1 {
		if line1[i] != line2[i] {
			return false
		}
	}
	return true

}

func isOppoSlice(line1 []int, line2 []int) bool {
	for i, _ := range line1 {
		if line1[i] == line2[i] {
			return false
		}
	}
	return true
}