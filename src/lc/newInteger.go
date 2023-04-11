package lc

import "strconv"

func NewInteger(n int) int {
	ans := ""
	for n > 0 {
		ans = strconv.Itoa(n%9) + ans
		n /= 9
	}
	res, err := strconv.Atoi(ans)
	if err != nil {
		return 0
	}
	return res
}
