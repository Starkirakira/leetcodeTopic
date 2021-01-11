package lc

import "math"

func LengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	dp := make([]int, len(nums))
	dp[0] = 1
	ans := 1
	for i := 1; i < len(nums); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = int(math.Max(float64(dp[i]),float64(dp[j] + 1)))
			}
		}
		ans = int(math.Max(float64(ans), float64(dp[i])))
	}
	return ans
}
