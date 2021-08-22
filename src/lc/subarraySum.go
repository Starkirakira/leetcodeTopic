package lc

func SubarraySum(nums []int, k int) int {
	sumMap := make(map[int]int)
	sum := 0
	ret := 0
	sumMap[0] = 1
	for _, j := range nums {
		sum += j
		if _, ok := sumMap[sum-k]; ok {
			ret += sumMap[sum-k]
		} else {
			ret += 0
		}
		if _, ok := sumMap[sum]; ok {
			sumMap[sum]++
		} else {
			sumMap[sum] = 1
		}

	}
	return ret
}
