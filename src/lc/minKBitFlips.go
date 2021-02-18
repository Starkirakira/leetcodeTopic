package lc

func MinKBitFlips(A []int, k int) int {
	n := len(A)
	diff := make([]int,  n+1)
	revCnt := 0
	ans := 0
	for i, v := range A {
		revCnt += diff[i]
		if (v+revCnt) % 2 == 0 {
			if i + k > n {
				return -1
			}
			ans++
			revCnt++
			diff[i+k]--
		}
	}
	return ans

}
