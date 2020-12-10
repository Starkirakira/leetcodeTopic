package lc

import "math"

func ClosestDivisors(num int) []int {
	ans := []int{0, int(1e9)}
	for i := num+1; i < num + 3; i++ {
		cur := divide(i)
		if abs(cur[0] - cur[1]) < abs(ans[0] - ans[1]) {
			ans = cur
		}
	}
	return ans
}

func divide(num int) []int {
	for i := int(math.Sqrt(float64(num))); i >= 0; i-- {
		if num % i == 0 {
			return []int{i, num/i}
		}
	}
	return []int{}

}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}