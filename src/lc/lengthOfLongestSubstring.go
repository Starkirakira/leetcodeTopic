package lc

func LengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	sRoom := make([]int, 128)
	left := -1
	maxLen := 0
	countBit := 0
	for i := 0; i < len(s); i++ {
		sRoom[s[i]]++
		if sRoom[s[i]] == 2 {
			countBit++
		}
		for countBit > 0 {
			left++
			sRoom[s[left]]--
			if sRoom[s[left]] == 1 {
				countBit--
			}
		}
		maxLen = max(maxLen, i-left)
	}

	return maxLen
}
