package lc

func RemoveElement(nums []int, val int) int {
	n := len(nums)
	//Change the position of the target number and the last number, then slice nums delete the last number and let length-1 and i-1
	for i := 0; i < n; i++ {
		if nums[i] == val {
			nums[i] = nums[n-1]
			nums = nums[:n-1]
			n--
			i--
		}
	}
	return len(nums)
}
