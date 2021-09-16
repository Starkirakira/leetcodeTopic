package lc

import "math"

func LargestRectangleArea(heights []int) int {
	heights = append([]int{-1}, append(heights, 0)...)
	stack := []int{}
	res := 0
	for i, h := range heights {
		for len(stack) > 1 && heights[stack[len(stack)-1]] > h {
			res = int(math.Max(float64(res), float64(heights[stack[len(stack)-1]]*(i-stack[len(stack)-2]-1))))
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
	return res

}
