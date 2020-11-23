package main

import "sort"

func pondSizes(land [][]int) []int{
	recBakc := []int{}
	for i := 0; i < len(land); i++ {
		for j:=0; j<len(land[i]);j++{
			count := 0
			size := 0
			if land[i][j] == 0 {
				size = dfs(i,j,count,land)
				if size != 0{
						recBakc = append(recBakc,size)
				}
			}
		}
	}
	sort.Ints(recBakc)
	return recBakc
}

func dfs(x int , y int , count int , land [][]int) int {
	land[x][y] = -1
	count++
	dx := []int{1, 1, 0,-1,-1,-1, 0, 1}
	dy := []int{0,-1,-1,-1, 0, 1, 1, 1}
	for i:=0; i<8; i++{
		x1 := x + dx[i]
		y1 := y + dy[i]
		if x1 >=0 && x1 < len(land) && y1 >=0 && y1 <len(land[0]) && land[x1][y1] == 0{
			count = dfs(x1,y1,count,land)
		}

	}
	return count

}