package lc

func IsIsomorphic(s string, t string) bool {
	listA := map[byte]byte{}
	listT := map[byte]byte{}

	for i := range s {
		x , y := s[i] , t[i]

		if listA[x] > 0 && listA[x] != y || listT[y] > 0 && listT[y] != x {
			return false
		}
		listA[x] = y
		listT[y] = x

	}

	return true



}
