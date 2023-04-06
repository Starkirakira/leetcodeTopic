package lc

func MaxPalindromicProduct(s string) (ans int) {
	var a, b []byte
	var f func(int)
	f = func(i int) {
		if i == len(s) {
			if len(a)*len(b) > ans && isPalindromic(a) && isPalindromic(b) {
				ans = len(a) * len(b)
			}
			return
		}
		f(i + 1)
		a = append(a, s[i])
		f(i + 1)
		a = a[:len(a)-1]

		b = append(b, s[i])
		f(i + 1)
		b = b[:len(b)-1]
	}
	f(0)
	return
}

func isPalindromic(s []byte) bool {
	for i, n := 0, len(s); i < n/2; i++ {
		if s[i] != s[n-1-i] {
			return false
		}
	}
	return true
}
