package lc

func TwoNumberSum(numbers []int, target int) []int {
	l1 := 0
	l2 := len(numbers) - 1
	for l1 < l2 {
		if numbers[l1]+numbers[l2] == target {
			return []int{l1, l2}
		} else if numbers[l1]+numbers[l2] < target {
			l1++
		} else {
			l2--
		}

	}
	return []int{}
}
