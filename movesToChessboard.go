package main

import "os"

func movesToChessboard(board [][]int) int {
	N := len(board)
	count := make(map[int]int)

	for _,row:= range board{
		code := 0
		for x:= range row{
			code = 2 * code + x
		}
		if _,ok:= count[code];!ok{
			
		}
	}


	return 0
}
