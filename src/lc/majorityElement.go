package lc

func MajorityElement(nums []int) int {
	//Also can use the sort array to judge it, just find the median value of array
	var count = make(map[int]int)
	cnt, maj := 0, 0
	for i := range nums {
		if _, ok := count[nums[i]]; ok {
			count[nums[i]]++
		} else {
			count[nums[i]] = 1
		}
		if count[nums[i]] > cnt {
			maj = nums[i]
			cnt = count[nums[i]]
		}
	}
	return maj
}
