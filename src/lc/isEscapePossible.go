package lc

type pairOne struct {
	x,y int
}
var dirsOne = []pairOne{{-1,0},{1,0},{0,-1},{0,1}}

func IsEscapePossible(block [][]int, source []int, target []int) bool {
	const (
		blocked = -1
		valid = 0
		found = 1
		boundary int = 1e6
	)
	n := len(block)
	if n < 2 {
		return true
	}

	blockSet := map[pairOne]bool{}
	for _,b := range block {
		blockSet[pairOne{b[0],b[1]}] = true
	}

	check := func(start, finish []int) int {
		sx, sy := start[0], start[1]
		fx, fy := finish[0], finish[1]
		countdown := n * (n - 1) / 2

		q := []pairOne{{sx, sy}}
		vis := map[pairOne]bool{{sx, sy} : true}
		for len(q) > 0 && countdown > 0 {
			p := q[0]
			q = q[1:]
			for _,d := range dirsOne {
				x, y := p.x+d.x, p.y+d.y
				np := pairOne{x,y}
				if 0 <= x && x < boundary && 0 <= y && y < boundary && !blockSet[np] && !vis[np] {
					if x == fx && y == fy {
						return found
					}
					countdown--
					vis[np] = true
					q = append(q,np)
				}
			}
		}
		if countdown > 0 {
			return blocked
		}
		return valid
	}
	res := check(source, target)
	return res == found || res == valid && check(target,source) != blocked
}
