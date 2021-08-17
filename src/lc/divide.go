package lc

import "math"

func Divide(a int, b int) int {
	flag := 0
	if a > 0 {
		a = -a
		flag += 1
	}
	if b > 0 {
		b = -b
		flag += 1
	}
	ret := calc(a, b)
	if flag != 1 && ret == math.MinInt32 {
		ret++
	}
	if flag == 1 {
		return ret
	}
	return -ret
}

func calc(a int, b int) int {
	ret := 0
	for a <= b {
		cnt := 1
		val := b
		for val >= math.MinInt32>>1 && a <= val<<1 {
			cnt += cnt
			val += val
		}
		ret -= cnt
		a -= val
	}
	return ret
}
