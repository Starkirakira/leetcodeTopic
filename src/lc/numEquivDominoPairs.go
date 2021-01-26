package lc

func NumEquivDominoPairs(dominoes [][]int) int {
	cnt := [100]int{}
	ans := 0
	for _, d := range dominoes {
		if d[0] > d[1] {
			d[0], d[1] = d[1], d[0]
		}
		v := d[0] * 10 + d[1]
		ans += cnt[v]
		cnt[v]++
	}
	return ans
}
