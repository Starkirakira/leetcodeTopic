package main

type MedianFinder struct {
	listA []int //保存前半部分的较小值
	listB []int//保存后半部分的较大值
}


/** initialize your data structure here. */
func ConstructorMedian() MedianFinder {
	return MedianFinder{listA: []int{}, listB: []int{}}
}


func (this *MedianFinder) AddNum(num int)  {
	if len(this.listA) != len(this.listB) {

	}
}


func (this *MedianFinder) FindMedian() float64 {

}
