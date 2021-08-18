package lc

import "strconv"

func AddBinary(a, b string) string {
	ret := []int{}
	La := len(a) - 1
	Lb := len(b) - 1
	count := 0

	x := 0
	y := 0
	for La >= 0 || Lb >= 0 {
		if La >= 0 {
			x = int(a[La] - '0')
		} else {
			x = 0
		}

		if Lb >= 0 {
			y = int(b[Lb] - '0')
		} else {
			y = 0
		}

		sum := x + y + count
		ret = append([]int{sum % 2}, ret...)
		count = sum / 2

		La--
		Lb--
	}
	if count != 0 {
		ret = append([]int{count}, ret...)
	}
	var res string

	for _, i := range ret {
		res += strconv.Itoa(i)
	}
	return res
}
