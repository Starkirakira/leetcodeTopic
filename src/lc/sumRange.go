package lc

type NumArray struct {
	sums []int
}


func ConstructorRange(nums []int) NumArray {
	sums := make([]int, len(nums) + 1)
	for i,v := range nums {
		sums[i+1] = sums[i] + v
	}
	return NumArray{sums}
}


func (this *NumArray) SumRange(i int, j int) int {
	return this.sums[j+1]-this.sums[i]
}
