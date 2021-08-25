package lc

func CheckInclusion(s1, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}
	sRoom := make([]int, 26)
	for i, j := range s1 {
		sRoom[j-'a']++
		sRoom[s2[i]-'a']--
	}
	if allRepeat(sRoom) {
		return true
	}

	for i := len(s1); i < len(s2); i++ {
		sRoom[s2[i-len(s1)]-'a']++
		sRoom[s2[i]-'a']--
		if allRepeat(sRoom) {
			return true
		}
	}
	return false
}

func allRepeat(repeat []int) bool {
	for _, i := range repeat {
		if i != 0 {
			return false
		}
	}
	return true
}
