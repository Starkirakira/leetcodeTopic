package lc

func ValidPalindrome(s string) bool {
	n := len(s)
	left := 0
	right := n - 1

	for left < right {
		if s[left] != s[right] {
			if isPalindrome(s, left+1, right) || isPalindrome(s, left, right-1) {
				return true
			} else {
				return false
			}
		}
		left++
		right--
	}
	return true
}
func isPalindrome(s string, left, right int) bool {
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}
