package lc

func AsteroidCollision(asteroids []int) []int {
	stack := make([]int, 0, len(asteroids)/2)
	for _, v := range asteroids {
		stat := -1
		for len(stack) != 0 && v < 0 && stack[len(stack)-1] > 0 {
			stat = stack[len(stack)-1] + v
			if stat < 0 {
				stack = stack[:len(stack)-1]
			} else if stat == 0 {
				stack = stack[:len(stack)-1]
				break
			} else {
				break
			}
		}
		if stat < 0 {
			stack = append(stack, v)
		}

	}
	return stack
}
