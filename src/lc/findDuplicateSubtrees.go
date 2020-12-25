package lc

import "strconv"

func FindDuplicateSubtrees(root *TreeNode) []*TreeNode {
	ans,_ := seriz(root, make([]*TreeNode, 0), map[string]int{})
	return ans
}

func seriz(root *TreeNode, ans []*TreeNode, m map[string]int) ([]*TreeNode, string) {
	if root == nil {
		return ans, "*"
	}
	var l,r string
	ans,l = seriz(root.Left, ans, m)
	ans,r = seriz(root.Right, ans, m)
	s :=   l + "," + r + "," + strconv.Itoa(root.Val)
	m[s]++
	if m[s] == 2 {
		ans = append(ans, root)
	}
	return ans, s

}