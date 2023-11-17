package lc

func BuildArray(target []int, n int) []string {
	//target数组单调递增
	//因为数组单调递增，所以目标数组的构建应该是目标元素被push进去后，连续push n-i-1个push，然后pop n-i-1次，再push一次目标元素
	var ans []string
	tmp := 0
	for _, tar := range target {
		times := tar - tmp - 1
		for i := 1; i <= times; i++ {
			ans = append(ans, "Push")
			ans = append(ans, "Pop")
		}
		ans = append(ans, "Push")
		tmp = tar
	}
	return ans
}
