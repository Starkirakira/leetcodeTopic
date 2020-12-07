package lc

func XorQueries(arr []int , queries [][]int) []int {
	n := len(arr)
	pre := make([]int,n+1)
	for i:=1;i<=n;i++{
		pre[i] = pre[i - 1] ^ arr[i - 1]
	}

	ans := make([]int , len(queries))
	for i,query := range queries{
		ans[i] = pre[query[0]] ^ pre[query[1] +1]
	}

	return ans


}
