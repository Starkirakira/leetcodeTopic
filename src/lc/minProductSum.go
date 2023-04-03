package lc

import "sort"

func MinProductSum(nums1 []int, nums2 []int) int {
	//nums1 can be sorted, nums2 is fixed
	//顺序和不小于乱序和，乱序和不小于逆序和
	sort.Ints(nums1)
	sort.Sort(sort.Reverse(sort.IntSlice(nums2)))
	var product int
	for i := range nums1 {
		product += nums1[i] * nums2[i]
	}
	return product
}
