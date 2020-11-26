package main

import "math"

func thirdMax(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) < 3 {
		return maxNums(nums[0],nums[1])
	}
	firstLargeNum := math.MinInt64
	secondLargeNum := math.MinInt64
	thirdLargeNum := math.MinInt64
	for i:=0;i<len(nums);i++{//1,1,2 f 12   s -11   t -1
		if nums[i]==firstLargeNum || nums[i]==secondLargeNum || nums[i] == thirdLargeNum{
			continue
		}
		if nums[i] > firstLargeNum {
			thirdLargeNum = secondLargeNum
			secondLargeNum = firstLargeNum
			firstLargeNum = nums[i]
		}else if nums[i] > secondLargeNum {
			thirdLargeNum = secondLargeNum
			secondLargeNum = nums[i]
		}else if nums[i] > thirdLargeNum{
			thirdLargeNum = nums[i]
		}
	}
	if thirdLargeNum == math.MaxInt64{
		thirdLargeNum = firstLargeNum
	}
	return thirdLargeNum
}

func maxNums(a int,b int) int{
	if a < b {
		return b
	}else{
		return a
	}
}

