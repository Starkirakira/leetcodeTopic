package lc

import "math"

func MinTaps(n int , ranges []int) int {
	//dp
	//prev := make([]int, n+1)//0,1,2,3,4,5,...,n+1
	//for i:=range prev{
	//	prev[i] = i
	//}
	//for i := 0; i <=n; i++{
	//	l := maxN(i - ranges[i] , 0)
	//	r := minN(i + ranges[i] , n)
	//	prev[r] = minN(prev[r] , l)
	//}
	//
	//dp := make([]int,n+1)
	//for i := range dp {
	//	dp[i] = math.MaxInt32
	//}
	//dp[0] = 0
	//for i :=1; i <=n; i++{
	//	for j := prev[i]; j<i; j++{
	//		if dp[j] != math.MaxInt32 {
	//			dp[i] = minN(dp[i] , dp[j] + 1)
	//		}
	//	}
	//}
	//if dp[n] == math.MaxInt32 {
	//	return -1
	//}else {
	//	return dp[n]
	//}

	//贪心
	prev := make([]int , n+1)
	for i := range prev {
		prev[i] = i
	}

	for i:=0; i<=n; i++{
		l := maxN(i - ranges[i] , 0)
		r := minN(i +ranges[i] , n)
		prev[r] = minN(prev[r],l)
	}

	breakpoint := n
	furthest := math.MaxInt32
	ans := 0

	for i:=n; i>0; i--{
		furthest = minN(furthest,prev[i])
		if i == breakpoint{
			if furthest >=i {
				return -1
			}
			breakpoint = furthest
			ans++
		}
	}
	return ans
}
func maxN(a int,b int) int{
	if a < b {
		return b
	}else{
		return a
	}
}
func minN(a int,b int) int{
	if a < b {
		return a
	}else{
		return b
	}
}
