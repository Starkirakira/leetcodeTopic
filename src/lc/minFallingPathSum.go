package lc

import "math"

func MinFallingPathSum(A [][]int) int {
	n := len(A)
	for r := n - 2; r >= 0; r-- {
		for c := 0; c < n; c++ {
			best := A[r + 1][c]
			if c > 0 {
				best = int(math.Min(float64(best),float64(A[r+1][c-1])))
			}
			if c + 1 < n {
				best = int(math.Min(float64(best),float64(A[r+1][c+1])))
			}
			A[r][c] += best
		}
	}

	ans := math.MaxInt32
	for _,x := range A[0] {
		ans = int(math.Min(float64(ans),float64(x)))
	}
	return ans
}
