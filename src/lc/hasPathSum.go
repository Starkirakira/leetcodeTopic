package lc

func HasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	var newRoot []*TreeNode
	var sumVal []int
	newRoot = append(newRoot, root)
	sumVal = append(sumVal, root.Val)

	for len(newRoot) != 0 {
		now := newRoot[0]
		newRoot = newRoot[1:]
		temp := sumVal[0]
		sumVal = sumVal[1:]
		if now.Left == nil && now.Right == nil {
			if temp == targetSum {
				return true
			}
			continue
		}

		if now.Left != nil {
			newRoot = append(newRoot, now.Left)
			sumVal = append(sumVal, now.Left.Val+temp)
		}

		if now.Right != nil {
			newRoot = append(newRoot, now.Right)
			sumVal = append(sumVal, now.Right.Val+temp)
		}
	}
	return false

}
