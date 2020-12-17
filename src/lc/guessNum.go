package lc

func Game(guess []int, answer []int) int {
	n := len(answer)
	ans := 0
	for i := 0; i < n; i++ {
		if guess[i] == answer[i] {
			ans++
		}
	}
	return ans
}
