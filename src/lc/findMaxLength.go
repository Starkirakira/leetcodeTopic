package lc

func FindMaxLength(nums []int) int {
	n := len(nums)
	mapLen := make(map[int]int)
	sum := 0
	mapLen[0] = -1
	maxLen := 0

	for i := 0; i < n; i++ {
		if nums[i] == 0 {
			nums[i] = -1
		}
		sum = sum + nums[i]
		if _, ok := mapLen[sum]; !ok {
			mapLen[sum] = i
		} else {
			maxLen = max(maxLen, i-mapLen[sum])
		}
	}
	return maxLen
}
