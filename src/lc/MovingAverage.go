package lc

type MovingAverage struct {
	size int
	list []int
}

/** Initialize your data structure here. */
func ConstructorMoving(size int) MovingAverage {
	return MovingAverage{size: size, list: []int{}}
}

func (this *MovingAverage) Next(val int) float64 {
	n := len(this.list)
	size := this.size
	if n < size {
		this.list = append(this.list, val)
	} else {
		this.list = this.list[1:]
		this.list = append(this.list, val)
	}
	if len(this.list) == 1 {
		return float64(this.list[0])
	}
	ret := 0
	for _, j := range this.list {
		ret += j
	}

	return float64(ret / len(this.list))
}

/**
 * Your MovingAverage object will be instantiated and called as such:
 * obj := Constructor(size);
 * param_1 := obj.Next(val);
 */
