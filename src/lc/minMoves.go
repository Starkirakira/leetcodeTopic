package lc

import "math"

func MinMoves(nums []int, limit int) int {
	diff := make([]int, 2 * limit + 2)
	n := len(nums)

	for i := 0; i < n / 2; i++ {
		A,B := nums[i],nums[n - i - 1]


		l := 1 + int(math.Min(float64(A),float64(B)))
		r := limit + int(math.Max(float64(A),float64(B)))
		sum := A + B
		diff[l]--
		diff[sum]--
		diff[sum + 1]++
		diff[r + 1]++

	}

	res := n
	for i := 2; i <= 2 * limit; i++ {
		n += diff[i]
		res = int(math.Min(float64(res),float64(n)))

	}
	return res

}
