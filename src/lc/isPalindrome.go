package lc

import "fmt"

func IsPalindrome(s string) bool {
	sre := []int{}
	for _, s := range s {
		if int(s) >= 65 && int(s) <= 90 {
			sre = append(sre, int(s))
		} else if int(s) >= 97 && int(s) <= 122 {
			sre = append(sre, int(s)-32)
		} else if (int(s) >= 48) && (int(s) <= 57) {
			sre = append(sre, int(s))
		}
	}
	fmt.Println(sre)
	spot := len(sre)
	for m, n := range sre {
		if n != sre[spot-m-1] {
			return false
		}
	}
	return true
}
