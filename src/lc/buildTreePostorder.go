package lc

func BuildTreePostorder(inorder []int, postorder []int) *TreeNode {
	inorderList := make(map[int]int)
	for i,j := range inorder {
		inorderList[j] = i
	}
	var helpOrder func(int, int) *TreeNode
	helpOrder = func(inLeft, inRight int) *TreeNode {
		if inLeft > inRight {
			return nil
		}
		val := postorder[len(postorder) - 1]
		postorder = postorder[:len(postorder) - 1]
		root := &TreeNode{Val: val}
		inorderRootIndex := inorderList[val]
		root.Right = helpOrder(inorderRootIndex + 1, inRight)
		root.Left = helpOrder(inLeft, inorderRootIndex - 1)
		return root
	}
	return helpOrder(0, len(inorder) - 1)
}


