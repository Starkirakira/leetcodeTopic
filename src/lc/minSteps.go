package lc

func MinSteps(n int) int {
	ans := 0
	d := 2
	for n>1 {
		for n%d == 0 {
			ans += d
			n /= d
		}
		d += 1
	}
	return ans
}
