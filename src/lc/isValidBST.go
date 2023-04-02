package lc

import "math"

func IsValidBST(root *TreeNode) bool {
	return isValidBSTHepler(root, math.MinInt64, math.MaxInt64)
}

func isValidBSTHepler(root *TreeNode, l, u int64) bool {
	if root == nil {
		return true
	}

	if int64(root.Val) <= l || int64(root.Val) >= u {
		return false
	}

	return isValidBSTHepler(root.Left, l, int64(root.Val)) && isValidBSTHepler(root.Right, int64(root.Val), u)
}
