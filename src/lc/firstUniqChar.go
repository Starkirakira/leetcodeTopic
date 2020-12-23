package lc

func FirstUniqChar(s string) int {
	list := [26]int{}

	for _,j := range s {
		list[j - 'a']++
	}

	for i,j := range s {
		if list[j-'a'] == 1 {
			return i
		}
	}
	return -1



}
