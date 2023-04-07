package lc

func SubsequenceSumOr(nums []int) int64 {
	n := len(nums)
	sums := make([]int, n)
	sums[0] = nums[0]
	for i := 1; i < n; i++ {
		sums[i] = sums[i-1] + nums[i]
	}
	var ans int
	for i := range nums {
		ans |= nums[i]
	}
	for i := range sums {
		ans |= sums[i]
	}
	return int64(ans)
}
