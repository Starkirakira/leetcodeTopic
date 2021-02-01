package lc

func FairCandySwap(A,B []int) []int {
	sumA := 0
	setA := map[int]struct{}{}
	for _, v := range A {
		sumA += v
		setA[v] = struct{}{}
	}
	sumB := 0
	for _,v := range B {
		sumB += v
	}

	delta := (sumA - sumB) / 2
	for i := 0; ; i++ {
		y := B[i]
		x := y + delta
		if _, has := setA[x]; has {
			return []int{x,y}
		}
	}
}
