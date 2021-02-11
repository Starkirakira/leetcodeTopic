package lc

import (
	"container/heap"
	"sort"
)

type KthLargest struct {
	Val int
	sort.IntSlice
}

func ConstructorKth(k int, nums []int) KthLargest {
	kl := KthLargest{Val: k}
	for _, x := range nums {
		kl.Add(x)
	}
	return kl

}

func (this *KthLargest) Push(v interface{}) {
	this.IntSlice = append(this.IntSlice, v.(int))
}
func (this *KthLargest) Pop() interface{} {
	a := this.IntSlice
	v := a[len(a) - 1]
	this.IntSlice = a[:len(a)-1]
	return v
}
func (this *KthLargest) Add(val int) int {
	heap.Push(this, val)
	if this.Len() > this.Val {
		heap.Pop(this)
	}
	return this.IntSlice[0]
}
