package lc

import "sort"

func MaxSumDivThree(nums []int) int {

	v := make([][]int, 3)
	for _, num := range nums {
		v[num%3] = append(v[num%3], num)
	}
	sort.Slice(v[1], func(i, j int) bool {
		return v[1][i] > v[1][j]
	})
	sort.Slice(v[2], func(i, j int) bool {
		return v[2][i] > v[2][j]
	})

	ans, lb, lc := 0, len(v[1]), len(v[2])
	for cntb := max(lb-2, 0); cntb <= lb; cntb++ {
		for cntc := max(lc-2, 0); cntc <= lc; cntc++ {
			if (cntb-cntc)%3 == 0 {
				ans = max(ans, accumulate(v[1][:cntb])+accumulate(v[2][:cntc]))
			}
		}
	}
	return ans + accumulate(v[0])

}
func accumulate(v []int) int {
	ans := 0
	for _, x := range v {
		ans += x
	}
	return ans
}
