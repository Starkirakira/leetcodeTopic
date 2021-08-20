package lc

import "sort"

func ThreeSum(nums []int) [][]int {
	n := len(nums)
	res := [][]int{}
	sort.Ints(nums)
	for i := 0; i < n; i++ {
		target := -nums[i]
		l1 := i + 1
		l2 := n - 1
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for l1 < l2 {
			if nums[l1]+nums[l2] == target {
				res = append(res, []int{nums[i], nums[l1], nums[l2]})
				for l1 < l2 {
					l1++
					if nums[l1-1] != nums[l1] {
						break
					}
				}
				for l1 < l2 {
					l2--
					if nums[l2] != nums[l2+1] {
						break
					}
				}
			} else if nums[l1]+nums[l2] < target {
				l1++
			} else {
				l2--
			}
		}
	}
	return res
}
