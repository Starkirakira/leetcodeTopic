package lc

import "math"

func MaxTurbulenceSize(arr []int) int {
	n := len(arr)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	dp1 := make([]int, n + 1)
	dp2 := make([]int, n + 1)
	dp1[1] = 1
	dp2[1] = 1
	result := 1
	for i := 2; i <= n; i++ {
		if arr[i - 1] > arr[i - 2] {
			dp1[i] = dp2[i - 1] + 1
		} else {
			dp1[i] = 1
		}
		if arr[i - 1] < arr[i - 2] {
			dp2[i] = dp1[i - 1] + 1
		} else {
			dp2[i] = 1
		}
		result = int(math.Max(float64(result), float64(dp1[i])))
		result = int(math.Max(float64(result), float64(dp2[i])))

	}
	return result
}
