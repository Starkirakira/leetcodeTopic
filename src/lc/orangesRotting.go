package lc

func OrangesRotting(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}

	n := 0
	health := 0
	row := len(grid)
	col := len(grid[0])
	beginOrange := []int{}
	var badOrange  [][]int
	dx := []int{0,1,0,-1}
	dy := []int{1,0,-1,0}
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if grid[i][j] == 2 {
				badOrange = append(badOrange,[]int{i,j})
			}else if grid[i][j] == 1{
				health++
			}

		}
	}

	for len(badOrange) != 0 && health > 0 {
		n++

		m := len(badOrange)

		for x:=0;x < m; x++ {
			beginOrange = badOrange[0]
			badOrange = badOrange[1:]
			for i := 0; i < 4; i++ {
				dxx := beginOrange[0]+dx[i]
				dyy := beginOrange[1]+dy[i]

				if dxx >= 0 && dxx < row && dyy < col && dyy >= 0 && grid[dxx][dyy] == 1 {
					grid[dxx][dyy] = 2
					badOrange = append(badOrange,[]int{dxx,dyy})
					health--
				}

			}
		}

	}


	if health > 0 {
		return -1
	}else {
		return n
	}

}
