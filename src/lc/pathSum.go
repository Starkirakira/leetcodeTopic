package lc

func PathSum(root *TreeNode, sum int) (ans [][]int) {
	path := []int{}

	var dfsii func(*TreeNode, int)

	dfsii = func(node *TreeNode, i int) {
		if node == nil {
			return
		}
		i -= node.Val
		path = append(path, node.Val)
		defer func() {path = path[:len(path) - 1]}()
		if node.Left == nil && node.Right == nil && i == 0 {
			ans = append(ans, append([]int(nil), path...))
			return
		}
		dfsii(node.Left, i)
		dfsii(node.Right, i)
	}
	dfsii(root, sum)
	return
}
