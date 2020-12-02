package main

import "container/heap"

type MinHeap []int
type MaxHeap []int



func (h MinHeap) Len() int {
	return len(h)
}

func (h MinHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h MinHeap) Swap(i, j int)  {
	h[i],h[j] = h[j],h[i]
}

func (h *MinHeap) Push(x interface{})  {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() interface{}  {
	val := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return val
}

func (h MaxHeap) Len() int {
	return len(h)
}

func (h MaxHeap) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h MaxHeap) Swap(i, j int)  {
	h[i],h[j] = h[j],h[i]
}

func (h *MaxHeap) Push(x interface{})  {
	*h = append(*h,x.(int))
}

func (h *MaxHeap) Pop() interface{} {
	val := (*h)[len(*h) - 1]
	*h = (*h)[:len(*h)-1]
	return val
}

type MedianFinder struct {
	minHeap *MinHeap //保存前半部分的较小值
	maxHeap *MaxHeap//保存后半部分的较大值
}







/** initialize your data structure here. */
func ConstructorMedian() MedianFinder {
	minHeap := &MinHeap{}
	maxHeap := &MaxHeap{}
	heap.Init(&MinHeap{})
	heap.Init(&MaxHeap{})

	return MedianFinder{
		minHeap: minHeap,
		maxHeap: maxHeap,
	}
}


func (this *MedianFinder) AddNum(num int)  {
	if this.minHeap.Len() == 0 && this.maxHeap.Len() == 0 {
		heap.Push(this.maxHeap, num)
		return
	}

	if num >= (*this.maxHeap)[0] {
		if this.minHeap.Len() == this.maxHeap.Len() {
			heap.Push(this.minHeap, num)
			heap.Push(this.maxHeap, heap.Pop(this.minHeap))
		} else {
			heap.Push(this.minHeap, num)
		}
	} else if num < (*this.maxHeap)[0] {
		if this.minHeap.Len() == this.maxHeap.Len() {
			heap.Push(this.maxHeap, num)
		} else {
			heap.Push(this.maxHeap, num)
			heap.Push(this.minHeap, heap.Pop(this.maxHeap))
		}
	}
}


func (this *MedianFinder) FindMedian() float64 {
	if this.minHeap.Len() == this.maxHeap.Len() {
		return float64((*this.minHeap)[0] + (*this.maxHeap)[0]) / 2.0
	} else {
		return float64((*this.maxHeap)[0])
	}
}


