package lc

type ListNode struct {
	Val int
	Next *ListNode
}

type TreeNodeii struct {
	Val int
	Left *TreeNodeii
	Right *TreeNodeii
}

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func maxP(a ...int) int {
	res := a[0]
	for _, v := range a[1:] {
		if v > res {
			res = v
		}
	}
	return res
}