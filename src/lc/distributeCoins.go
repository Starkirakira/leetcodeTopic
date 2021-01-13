package lc

import "math"

func DistributeCoins(root *TreeNode) int {
	ans := 0
	var dfs979  func(node *TreeNode) int
	dfs979 = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		l := dfs979(node.Left)
		r := dfs979(node.Right)
		ans += int(math.Abs(float64(l))) + int(math.Abs(float64(r)))
		return node.Val + l + r -1
	}
	dfs979(root)
	return ans
}

