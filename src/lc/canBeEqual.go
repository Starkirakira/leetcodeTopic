package lc

import "sort"

func CanBeEqual(target []int, arr []int) bool {
	sort.Ints(target)
	sort.Ints(arr)
	for i := 0; i < len(arr); i++ {
		if target[i] != arr[i] {
			return false
		}
	}
	return true
}
