package lc

func CountSubstrings(s string) int {
	n := len(s)
	ans := 0
	for i := 0; i < 2*n-1; i++ {
		left := i / 2
		right := i/2 + i%2
		for left >= 0 && right < n && s[left] == s[right] {
			left--
			right++
			ans++
		}
	}
	return ans
}
