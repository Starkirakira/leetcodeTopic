package lc

func PlusOne(digits []int) []int {
	n := len(digits)
	//1. The last digit unequal 9, then +1
	//2. The array have several 9, so need find the first digit unequal 9 from last, let this num +1, and set the rest of numbers to 0
	for i := n - 1; i >= 0; i-- {
		if digits[i] != 9 {
			digits[i]++
			for j := i + 1; j < n; j++ {
				digits[j] = 0
			}
			return digits
		}
	}

	digits = make([]int, n+1)
	digits[0] = 1
	return digits
}
