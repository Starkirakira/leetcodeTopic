package lc


func Flatten(root *TreeNode)  {
	if root == nil {
		return
	}
	stack := [] *TreeNode{root}
	var prev *TreeNode
	for len(stack) > 0 {
		curr := stack[len(stack) - 1]
		stack = stack[:len(stack) - 1]
		if prev != nil {
			prev.Left =nil
			prev.Right = curr
		}
		left, right := curr.Left, curr.Right
		if right != nil {
			stack = append(stack, right)
		}
		if left != nil {
			stack = append(stack, left)
		}
		prev = curr
	}
}
