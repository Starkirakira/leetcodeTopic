package lc

import (
	mrand "math/rand"
	"time"
)

type RandomizedSet struct {
	numLocation map[int]int
	nums        []int
	rand        *mrand.Rand
}

/** Initialize your data structure here. */
func ConstructorOFFER30() RandomizedSet {
	location := make(map[int]int)
	nums := []int{}
	return RandomizedSet{location, nums, mrand.New(mrand.NewSource(time.Now().UnixNano()))}
}

/** Inserts a value to the set. Returns true if the set did not already contain the specified element. */
func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.numLocation[val]; ok {
		return false
	}
	this.numLocation[val] = len(this.nums)
	this.nums = append(this.nums, val)
	return true
}

/** Removes a value from the set. Returns true if the set contained the specified element. */
func (this *RandomizedSet) Remove(val int) bool {
	if _, ok := this.numLocation[val]; !ok {
		return false
	}
	location := this.numLocation[val]
	this.nums[location] = this.nums[len(this.nums)-1]
	this.numLocation[this.nums[len(this.nums)-1]] = location
	delete(this.numLocation, val)
	this.nums = this.nums[:len(this.nums)-1]
	return true

}

/** Get a random element from the set. */
func (this *RandomizedSet) GetRandom() int {
	return this.nums[this.rand.Intn(len(this.nums))]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
