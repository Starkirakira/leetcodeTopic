package lc

import (
	"math"
	"sort"
)

func MaximumProduct(nums []int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	n := len(nums)
	return int(math.Max(float64(nums[n - 1] * nums[ n - 2] * nums[n - 3]), float64(nums[0] * nums[1] * nums[n - 1])))

}
