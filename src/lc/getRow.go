package lc

func GetRow(rowIndex int) []int {
	pre, cur:= []int{},[]int{}
	for i := 0; i <= rowIndex; i++ {
		cur = make([]int, i+1)
		cur[0] , cur[i] = 1,1
		for j :=1; j < i; j++ {
			cur[j] = pre[j - 1] + pre[j]
		}
		pre = cur
	}
	return pre
}
