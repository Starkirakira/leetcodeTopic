package lc

import "math"

func MaxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return int(math.Max(float64(MaxDepth(root.Left)),float64(MaxDepth(root.Right)))) + 1
}
