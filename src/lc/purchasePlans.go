package lc

import "sort"

func PurchasePlans(nums []int, target int) int {
	ans := 0
	sort.Ints(nums)
	j := len(nums) - 1
	for i := 0; i < len(nums); i++ {
		for ; j > i; j-- {
			if nums[i]+nums[j] <= target {
				break
			}
		}
		if j > i {
			ans += j - i
		}
	}
	return ans % (1e9 + 7)
}
