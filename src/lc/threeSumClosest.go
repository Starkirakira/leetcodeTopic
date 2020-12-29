package lc

import (
	"math"
	"sort"
)

func ThreeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	n := len(nums)
	ans := math.MaxInt32
	update := func(cur int) {
		if int(math.Abs(float64(cur - target))) < int(math.Abs(float64(ans - target))) {
			ans = cur
		}
	}
	for i := 0; i < n; i++ {
		if i > 0 && nums[i] == nums[i - 1] {
			continue
		}

		j,k := i + 1, n - 1

		for j < k {
			sum := nums[i] + nums[j] + nums[k]
			if sum == target {
				return target
			}
			update(sum)
			if sum > target {
				k0 := k - 1
				for j < k0 && nums[k0] == nums[k] {
					k0--
				}
				k = k0
			} else {
				j0 := j + 1
				for j0 < k && nums[j0] == nums[j] {
					j0++
				}
				j = j0
			}

		}
	}
	return ans

}
