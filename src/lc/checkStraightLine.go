package lc

func CheckStraightLine(coordinates [][]int) bool {
	n := len(coordinates)

	for i := 1; i < n - 1; i++ {
		//(x2-x1)x(y3-y2) = (y2-y1)*(x3-x2)  x1y1(i-1)(0,1)  x2y2 i(0,1)  x3y3 (i+1)(0,1)
		if (coordinates[i][0] - coordinates[i-1][0]) * (coordinates[i + 1][1] - coordinates[i][1]) != (coordinates[i][1] - coordinates[i - 1][1]) * (coordinates[i + 1][0] - coordinates[i][0]) {
			return false
		}
	}
	return true
}
