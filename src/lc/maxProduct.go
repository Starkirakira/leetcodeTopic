package lc

func MaxProduct(words []string) int {
	if words == nil || len(words) < 2 {
		return 0
	}

	length := len(words)
	masks := make([]int, length)
	for i, word := range words[:] {
		for _, c := range word {
			//'a' - 'a' === 97 - 97 so this range is 0-25
			masks[i] |= 1 << uint(c-'a')
		}
	}

	maxs := 0

	for i := 0; i < length; i++ {
		for j := 0; j < length; j++ {
			if masks[i]&masks[j] == 0 {
				prods := len(words[i]) * len(words[j])
				if prods > maxs {
					maxs = prods
				}
			}
		}
	}
	return maxs
}
