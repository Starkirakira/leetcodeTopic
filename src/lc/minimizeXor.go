package lc

import "math/bits"

func MinimizeXor(num1 int, num2 int) int {
	count1 := bits.OnesCount(uint(num1))
	count2 := bits.OnesCount(uint(num2))

	for ; count2 < count1; count2++ {
		num1 &= num1 - 1
	}
	for ; count2 > count1; count2-- {
		num1 |= num1 + 1
	}
	return num1
}
