package lc

func MinWindow(s, t string) string {
	if len(s) < len(t) {
		return ""
	}

	count := 'z' - 'A' + 1
	nums := make([]int, count)
	used := make([]bool, count)
	status := 0
	res := ""

	for i := 0; i < len(t); i++ {
		k := int(t[i] - 'A')
		if nums[k] == 0 {
			status++
		}
		nums[k]--
		used[k] = true
	}

	left := 0
	for right := 0; right < len(s); right++ {
		k := int(s[right] - 'A')
		if used[k] == false {
			continue
		}
		nums[k]++
		if nums[k] == 0 {
			status--
		}
		if status == 0 {
			for !used[int(s[left]-'A')] || nums[int(s[left])-'A']-1 >= 0 {
				nums[int(s[left])-'A']--
				left++
			}
			if right-left+1 < len(res) || len(res) == 0 {
				res = s[left : right+1]
			}
		}
	}
	return res

}
