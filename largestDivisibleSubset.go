package main

import (
	"sort"
)

func largestDivisibleSubset(nums []int) []int {
	if len(nums)==0 {
		return []int{}
	}
	sort.Ints(nums)
	dp := make([]int,len(nums))
	maxRetSize := -1
	maxRetIndex := -1

	for i:=0; i<len(nums);i++{
		retSize := 0
		for j:=0;j<i;j++{
			if nums[i]%nums[j] == 0 && retSize < dp[j] {
				retSize = dp[j]
			}
		}
		dp[i] = retSize + 1
		if maxRetSize < dp[i] {
			maxRetSize = dp[i]
			maxRetIndex = i
		}

	}

	res := []int{}
	currSize := maxRetSize
	currTail := nums[maxRetIndex]
	for i:= maxRetIndex;i>=0;i--{
		if currSize == 0{
			break
		}
		if currTail % nums[i] == 0 && currSize == dp[i]{
			res = append([]int{nums[i]},res...)
			currTail = nums[i]
			currSize -= 1
		}
	}
	return res
}