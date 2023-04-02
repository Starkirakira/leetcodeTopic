package lc

import "math"

func ThirdMax(nums []int64) int64 {
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) < 3 {
		return int64(maxNums(nums[0], nums[1]))
	}
	var firstLargeNum int64 = math.MinInt64
	var secondLargeNum int64 = math.MinInt64
	var thirdLargeNum int64 = math.MinInt64
	for i := 0; i < len(nums); i++ { //1,1,2 f 112   s -1-11   t -1-1-1

		if nums[i] > firstLargeNum {
			thirdLargeNum = secondLargeNum
			secondLargeNum = firstLargeNum
			firstLargeNum = nums[i]
		} else if nums[i] < firstLargeNum && (nums[i] > secondLargeNum) {
			thirdLargeNum = secondLargeNum
			secondLargeNum = nums[i]
		} else if nums[i] < secondLargeNum && (nums[i] > thirdLargeNum) {
			thirdLargeNum = nums[i]
		}
	}
	if thirdLargeNum == math.MinInt64 {
		return firstLargeNum
	} else {
		return thirdLargeNum
	}

}

func maxNums(a int64, b int64) int64 {
	if a < b {
		return b
	} else {
		return a
	}
}
