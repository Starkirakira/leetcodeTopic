package lc

// 多级双向链表
type ListDoublePathList struct {
	Val   int
	Prev  *ListDoublePathList
	Next  *ListDoublePathList
	Child *ListDoublePathList
}

// 单向链表
type ListNode struct {
	Val  int
	Next *ListNode
}

// 二叉树
type TreeNodeii struct {
	Val   int
	Left  *TreeNodeii
	Right *TreeNodeii
}

// 二叉树
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 带权并查集
type unionFind1 struct {
	parent, size []int
	setCount     int
}

func newUnionFind1(n int) *unionFind1 {
	parent := make([]int, n)
	size := make([]int, n)
	for i := range parent {
		parent[i] = i
		size[i] = 1
	}
	return &unionFind1{parent, size, n}
}

func (uf *unionFind1) find(x int) int {
	if uf.parent[x] != x {
		uf.parent[x] = uf.find(uf.parent[x])
	}
	return uf.parent[x]
}

func (uf *unionFind1) union(x, y int) bool {
	fx, fy := uf.find(x), uf.find(y)
	if fx == fy {
		return false
	}
	if uf.size[fx] < uf.size[fy] {
		fx, fy = fy, fx
	}
	uf.size[fx] += uf.size[fy]
	uf.parent[fy] = fx
	uf.setCount--
	return true
}

func (uf *unionFind1) isSameSet(x, y int) bool {
	return uf.find(x) == uf.find(y)
}

func maxP(a ...int64) int64 {
	res := a[0]
	for _, v := range a[1:] {
		if v > res {
			res = v
		}
	}
	return res
}
