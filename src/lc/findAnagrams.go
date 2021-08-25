package lc

func FindAnagrams(s, p string) []int {
	if len(s) < len(p) {
		return []int{}
	}
	sRoom := make([]int, 26)
	ret := []int{}
	for i := 0; i < len(p); i++ {
		sRoom[p[i]-'a']++
		sRoom[s[i]-'a']--
	}

	if allRepeat(sRoom) {
		ret = append(ret, 0)
	}

	for i := len(p); i < len(s); i++ {
		sRoom[s[i]-'a']--
		sRoom[s[i-len(p)]-'a']++
		if allRepeat(sRoom) {
			ret = append(ret, i-len(p)+1)
		}
	}

	return ret

}
