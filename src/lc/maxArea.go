package lc

func MaxArea(height []int) int {
	n := len(height)
	l,r := 0,n-1
	ans := 0
	for l <= r {
		area := minArea(height[l], height[r]) * (r - l)
		ans = maxArea(area, ans)
		if height[l] <= height[r] {
			l++
		} else {
			r--
		}
	}
	return ans

}

func maxArea(a int, b int) int {
	if a > b {
		return a
	}else {
		return b
	}
}

func minArea(a int, b int) int {
	if a > b {
		return b
	}else {
		return a
	}
}