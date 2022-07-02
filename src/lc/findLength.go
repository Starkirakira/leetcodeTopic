package lc

func FindLength(nums1 []int, nums2 []int) int {
	//[xxxxxx] [yyyyyyy]
	//  i		j
	//dp[i][j] = dp[i+1][j+1] + 1

	n, m := len(nums1), len(nums2)
	dp := make([][]int, n+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, m+1)
	}
	ans := 0
	for i := n - 1; i >= 0; i-- {
		for j := m; j >= 0; j-- {
			if nums1[i] == nums2[j] {
				dp[i][j] = dp[i+1][j+1] + 1
			} else {
				dp[i][j] = 0
			}
			if dp[i][j] > ans {
				ans = dp[i][j]
			}
		}
	}
	return ans

}
