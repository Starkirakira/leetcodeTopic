package lc

func SingleNumber(nums []int) int {
	res := int32(0)
	for i := 0; i < 32; i++ {
		bitSum := int32(0)
		for _, j := range nums {
			bitSum += (int32(j) >> i) & 1
		}
		res |= (bitSum % 3) << i
	}
	return int(res)

}
