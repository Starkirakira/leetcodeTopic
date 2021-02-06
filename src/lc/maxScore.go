package lc

func MaxScore(cardPoints []int, k int) int {
	n := len(cardPoints)
	windowSize := n - k
	sum := 0
	for _,pt := range cardPoints[:windowSize] {
		sum += pt
	}
	minSum := sum
	for i := windowSize; i < n; i++ {
		sum +=cardPoints[i] - cardPoints[i - windowSize]
		minSum = min(minSum, sum)
	}
	total := 0
	for _, pt := range cardPoints {
		total += pt
	}
	return total - minSum
}
