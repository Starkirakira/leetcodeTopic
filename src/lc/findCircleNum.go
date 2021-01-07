package lc

func FindCircleNum(isConnected [][]int) (ans int) {
	vis := make([]bool, len(isConnected))
	var dfso func(int)
	dfso = func(i int) {
		vis[i] = true
		for to, conn := range isConnected[i] {
			if conn == 1 && !vis[to] {
				dfso(to)
			}
		}
	}
	for i,j := range vis {
		if !j {
			ans ++
			dfso(i)
		}
	}
	return
}
