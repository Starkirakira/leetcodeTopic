package lc

func CountBits(num int) []int {
	bits := make([]int, num+1)
	hightBit := 0
	for i := 1; i <= num; i++ {
		if i & (i - 1) == 0 {
			hightBit = 1
		}
		bits[i] = bits[i - hightBit] + 1
	}
	return bits
}
