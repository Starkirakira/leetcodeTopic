package lc

func MinSetSize(arr []int) int {
	countmap := make(map[int]int, 0)
	maxcount := 0
	for _,num :=range arr {
		countmap[num]++
		if countmap[num] > maxcount {
			maxcount = countmap[num]
		}
	}
	bucket := make([]int, maxcount + 1)
	for _,count := range countmap {
		bucket[count]++
	}
	n := (len(arr) - 1) / 2
	ret := 0
	for count := maxcount; count > 0; count-- {
		v := bucket[count]
		if v == 0 {
			continue
		}
		if v*count < n {
			ret += v
			n -= v * count
		} else {
			ret += n / count + 1
			return ret
		}
	}
	return ret
}
