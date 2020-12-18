package lc

func FindTheDifference(s string, t string) byte {
	list := [26]int{}
	var a byte
	for _,i := range s {
		list[i-'a']++
	}

	for i,j := range t {
		list[j - 'a']--
		if list[j - 'a'] < 0 {
			a = t[i]
		}
	}
	return a
}
