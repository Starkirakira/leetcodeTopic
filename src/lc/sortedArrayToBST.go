package lc

func SortedArrayToBST(nums []int) *TreeNode {
 	var aToBST func([]int,int,int)*TreeNode

 	aToBST = func(ints []int, i int, i2 int) *TreeNode {
		if i > i2 {
			return nil
		}
		mid := (i + i2) / 2
		root := &TreeNode{Val: ints[mid]}
		root.Left = aToBST(ints, i, mid - 1)
		root.Right = aToBST(ints, mid + 1, i2)
		return root
	}
	return aToBST(nums, 0, len(nums)-1)
}
