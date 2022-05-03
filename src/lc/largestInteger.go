package lc

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func LargestInteger(num int) int {
	s := strconv.Itoa(num)
	nums := strings.Split(s, "")
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			a, _ := strconv.Atoi(nums[i])
			b, _ := strconv.Atoi(nums[j])
			if (a-b)%2 == 0 && a < b {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
	fmt.Println(nums)
	res := 0
	for i := 0; i < len(nums); i++ {
		a, _ := strconv.Atoi(nums[i])
		res += a * int(math.Pow(float64(10), float64(len(nums)-1-i)))
	}
	return res
}
