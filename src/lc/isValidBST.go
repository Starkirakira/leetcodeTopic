package lc

import "math"

func IsValidBST(root *TreeNode) bool {
	return isValidBSTHepler(root, math.MinInt64, math.MaxInt64)
}

func isValidBSTHepler(root *TreeNode, l,u int) bool {
	if root == nil {
		return true
	}

	if root.Val <= l || root.Val >= u {
		return false
	}

	return isValidBSTHepler(root.Left, l, root.Val) && isValidBSTHepler(root.Right, root.Val, u)
}