package lc

func DailyTemperatures(temperatures []int) []int {
	stack := make([]int, len(temperatures))
	ans := make([]int, len(temperatures))
	for i, temperature := range temperatures {
		for len(stack) > 0 && temperature > temperatures[stack[len(stack)-1]] {
			pre := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			ans[pre] = i - pre
		}
		stack = append(stack, i)
	}
	return ans
}
