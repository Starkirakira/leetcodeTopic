package lc

import (
	"math"
)

func MaxScore01(s string) int {
	n := len(s)
	sum0 := 0
	sum1 := 0
	res := 0
	for _, i2 := range s {
		if i2 == '1' {
			sum1++
		}
	}
	//n-1 make sure that string s could be split into two substrings
	for i := 0; i < n-1; i++ {
		if s[i] == '0' {
			sum0++
		} else {
			sum1--
		}
		res = int(math.Max(float64(res), float64(sum0+sum1)))
	}
	return res
}
