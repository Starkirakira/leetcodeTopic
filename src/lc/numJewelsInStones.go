package lc

func NumJewelsInStones(jewels string, stones string) int {
	num := 0
	for i := 0; i < len(jewels); i++ {
		for j := 0; j < len(stones); j++ {
			if jewels[i] == stones[j] {
				num++
			}
		}
	}

	return num
}
