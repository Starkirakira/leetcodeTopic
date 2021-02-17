package lc

func MatrixReshape(nums [][]int, r, c int) [][]int {
	n, m := len(nums), len(nums[0])
	if m*n != r*c {
		return nums
	}
	ans := make([][]int, r)
	for i := range ans {
		ans[i] = make([]int,c)
	}
	for i := 0; i < n*m; i++ {
		ans[i/c][i%c] = nums[i/m][i%m]
	}
	return ans
}
