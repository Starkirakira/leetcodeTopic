package lc

import (
	"sort"
	"strings"
)

func ReorderLogFiles(logs []string) []string {
	sort.SliceStable(logs, func(i, j int) bool {
		l1 := strings.Split(logs[i], " ")
		l2 := strings.Split(logs[j], " ")
		t1, t2 := 0, 0

		if l1[1][0] < '0' || l1[1][0] > '9' {
			t1 = 1
		}

		if l2[1][0] < '0' || l2[1][0] > '9' {
			t2 = 1
		}

		if t1 == 1 && t2 == 0 {
			return true
		}
		if t1 == 0 && t2 == 1 {
			return false
		}

		if t1 == 0 && t2 == 0 {
			return i < j
		}
		for i := 1; i < len(l1) || i < len(l2); i++ {
			if i < len(l1) && i < len(l2) {
				if l1[i] != l2[i] {
					return l1[i] < l2[i]
				}
				continue
			}
			if i < len(l1) {
				return true
			}
			if i < len(l2) {
				return false
			}
		}

		return l1[0] < l2[0]
	})
	return logs
}
