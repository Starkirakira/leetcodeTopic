package lc

import "sort"

func NumSubseq(nums []int, target int) int {
	sort.Ints(nums)
	i, j, ans, pow := 0, len(nums) - 1, 0, map[int]int{0:1}

	for n := 1; n < len(nums); n++ {
		pow[n] = (pow[n - 1] * 2) % (1e9 + 7)
	}
	for j >= i {
		for nums[i] + nums[j] > target {
			if j == 0 {
				return ans
			}
			j--
		}
		ans += pow[j - i]
		ans %= 1e9 + 7
		i++
	}
	return ans

}
