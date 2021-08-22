package lc

import "math"

func MinSubArrayLen(target int, nums []int) int {
	n := len(nums)
	left := 0
	sum := 0
	MAX_INT := math.MaxInt16
	for right := 0; right < n; right++ {
		sum += nums[right]
		for sum >= target && left <= right {
			MAX_INT = min(MAX_INT, right-left+1)
			if MAX_INT == 1 {
				return MAX_INT
			}
			sum -= nums[left]
			left++
		}
	}
	if MAX_INT == math.MaxInt16 {
		return 0
	}
	return MAX_INT

}
