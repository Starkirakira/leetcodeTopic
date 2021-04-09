package lc

func FindMin(nums []int) int {
	l,r := 0, len(nums) - 1
	for l < r {
		p := l + (r - l) / 2
		if nums[p] < nums[r] {
			r = p
		} else if nums[p] > nums[r] {
			l = p + 1
		} else {
			r--
		}
	}
	return nums[l]
}
