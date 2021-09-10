package lc

import (
	"math"
)

func IsAlienSorted(words []string, order string) bool {
	word := make(map[string]int)
	for i, j := range order {
		word[string(j)] = i
	}

	for i := 0; i < len(words)-1; i++ {
		s1 := words[i]
		s2 := words[i+1]
		l1 := len(s1)
		l2 := len(s2)
		for j := 0; j < int(math.Max(float64(l1), float64(l2))); j++ {
			idx1 := 0
			idx2 := 0
			if j >= l1 {
				idx1 = -1
			} else {
				idx1 = word[string(s1[j])]
			}
			if j >= l2 {
				idx2 = -1
			} else {
				idx2 = word[string(s2[j])]
			}
			if idx1 > idx2 {
				return false
			}
			if idx1 < idx2 {
				break
			}
		}
	}
	return true
}
