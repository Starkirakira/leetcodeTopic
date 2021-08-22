package lc

func NumsSubarrayProductLessThank(nums []int, k int) int {
	n := len(nums)
	left := 0
	product := 1
	amo := 0

	for right := 0; right < n; right++ {
		product *= nums[right]
		for product >= k && left <= right {
			product /= nums[left]
			left++
		}
		if left <= right {
			amo += right - left + 1
		}
	}
	return amo
}
