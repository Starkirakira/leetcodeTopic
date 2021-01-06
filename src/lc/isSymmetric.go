package lc

func IsSymmetric(root *TreeNode) bool {
	var checkthis func(p,q *TreeNode) bool

	checkthis = func(p, q *TreeNode) bool {
		if p == nil && q == nil {
			return true
		}
		if p == nil || q == nil {
			return false
		}
		return p.Val == q.Val && checkthis(p.Left, q.Right) && checkthis(p.Right, q.Left)
	}
	return checkthis(root, root)
}
