package lc

func IsPathCrossing(path string) bool {
	var cross [2]int
	cross = [2]int{0, 0}
	tmp := make(map[[2]int]struct{})
	tmp[cross] = struct{}{}
	for _, i2 := range path {
		switch i2 {
		case 'N':
			cross[1] += 1
			break
		case 'S':
			cross[1] -= 1
			break
		case 'E':
			cross[0] += 1
			break
		case 'W':
			cross[0] -= 1
			break

		}
		if _, ok := tmp[cross]; !ok {
			tmp[cross] = struct{}{}
		} else {
			return true
		}
	}
	return false
}
