package lc

import "math"

func FindLengthOfLCIS(nums []int) int {
	start := 0
	ans := 0
	for i, v := range nums {
		if i > 0 && v <= nums[i - 1] {
			start = i
		}
		ans = int(math.Max(float64(ans), float64(i - start + 1)))
	}
	return ans
}
