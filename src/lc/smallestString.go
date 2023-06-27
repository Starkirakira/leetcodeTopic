package lc

func SmallestString(s string) string {
	// "cbabc"
	var tmp []byte

	for i := range s {
		tmp = append(tmp, s[i])
	}
	//fmt.Println('a')
	//return "s"
	i := 0
	for i < len(tmp) && tmp[i] == 'a' {
		i++
	}
	if i == len(tmp) {
		tmp[i-1] = 'z'
	}
	for i < len(tmp) && tmp[i] != 'a' {
		tmp[i] -= 1
		i += 1
	}
	return string(tmp)
}
