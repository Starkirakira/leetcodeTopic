package lc

func MaxSatisfied(customers []int, grumpy []int, x int) int {
	total := 0
	n := len(customers)
	for i := 0; i < n; i++ {
		if grumpy[i] == 0 {
			total += customers[i]
		}
	}
	increase := 0
	for i := 0; i < x; i++ {
		increase += customers[i] * grumpy[i]
	}
	maxIncrease := increase
	for i := x; i < n; i++ {
		increase = increase - customers[i-x]*grumpy[i-x] + customers[i]*grumpy[i]
		maxIncrease = max(maxIncrease, increase)
	}
	return total + maxIncrease
}
