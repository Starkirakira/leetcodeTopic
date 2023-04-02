package lc

import (
	"math"
)

func MinAreaRect(points [][]int) int64 {
	pointSet := make(map[int]struct{})
	for _, point := range points {
		pointSet[4001*point[0]+point[1]] = struct{}{}
	}

	var maxRes int64 = math.MaxInt64
	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			if points[i][0] != points[j][0] && points[i][1] != points[j][1] {
				_, ok1 := pointSet[4001*points[i][0]+points[j][1]]
				_, ok2 := pointSet[4001*points[j][0]+points[i][1]]
				if ok1 && ok2 {
					maxRes = int64(int(math.Min(float64(maxRes), math.Abs(float64(points[j][0]-points[i][0]))*math.Abs(float64(points[j][1]-points[i][1])))))

				}

			}
		}
	}
	if maxRes < math.MaxInt64 {
		return maxRes
	}
	return 0

}
