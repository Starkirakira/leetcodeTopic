package lc

import "math"

func GetMoneyAmount(n int) int {
	if(n == 1){
		return 0
	}
	 dp := make([][]int,n+1)
	for i := range dp{
		dp[i] = make([]int,n+1)
	}

	for i :=2;i<=n;i++{
		for j:=1;j<=n-i+1;j++{
			minres := math.MaxInt32
			for v:= j + (i-1)/2;v<j+i-1;v++{
				res := v + max(dp[j][v-1],dp[v+1][i+j-1])
				minres = min(res,minres)
			}
			dp[j][j+i-1]=minres
		}


	}

	return dp[1][n]
}

func min(a int,b int) int {
	if a < b{
		return a
	}else{
		return b
	}
}

func max(a int,b int) int {
	if a < b{
		return b
	}else{
		return a
	}
}