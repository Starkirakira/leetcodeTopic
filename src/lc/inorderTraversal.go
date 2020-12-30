package lc

func InorderTraversal(root *TreeNode) []int {
	res := []int{}
	for root != nil {
		if root.Left != nil {
			predecessor := root.Left
			for predecessor.Right != nil && predecessor.Right != root {
				predecessor = predecessor.Right
			}
			if predecessor.Right == nil {
				predecessor.Right = root
				root = root.Left
			} else {
				res = append(res, root.Val)
				predecessor.Right = nil
				root = root.Right
			}
		} else {
			res = append(res, root.Val)
			root = root.Right
		}
	}
	return res
}
