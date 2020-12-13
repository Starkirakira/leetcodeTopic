package lc

import (
	"math"
)

func FindPairs(nums []int , k int) map[int]int {
	pairs := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if math.Abs(float64(nums[i]-nums[j])) == float64(k) && i != j  {
				if _,ok := pairs[nums[i]];!ok {
					if nums[i] - nums[j] >= 0 {
						pairs[nums[i]] = nums[j]
					}
				}
			}
		}
	}

	//pair := len(pairs)


	return pairs
}
