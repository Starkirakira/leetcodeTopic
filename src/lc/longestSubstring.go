package lc

import "strings"

func LongestSubstring(s string, k int) int {
	ans := 0
	if s == "" {
		return 0
	}
	cnt := [26]int{}
	for _, ch := range s {
		cnt[ch-'a']++
	}
	var split byte
	for i, c := range cnt[:] {
		if 0 < c && c < k {
			split = 'a' + byte(i)
			break
		}
	}
	if split == 0 {
		return len(s)
	}
	for _, subStr := range strings.Split(s, string(split)) {
		ans = max(ans, LongestSubstring(subStr, k))
	}
	return ans
}
