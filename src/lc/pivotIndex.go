package lc

func PivotIndex(nums []int) int {
	n := len(nums)
	res := -1
	sums := func(ints []int) (ans int) {
		for i := 0; i < len(ints); i++ {
			ans += ints[i]
		}
		return ans
	}
	//fmt.Println(sums(nums[0:2]))
	
	for i := 0; i < n; i++ {
		if sums(nums[0:i]) == sums(nums[i + 1:]) {
			res = i
			break
		}
	}
	return res
}
