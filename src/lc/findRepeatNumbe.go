package lc

func FindRepeatNumbe(nums []int) int {
	chMap := make(map[int]bool)
	ans := -1
	for _,j := range nums {
		if ok := chMap[j];!ok {
			chMap[j] = true
		} else {
			ans = j
			break
		}
	}
	return ans
}
