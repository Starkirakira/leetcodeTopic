package lc

import (
	"math"
	"sort"
	"strconv"
	"strings"
)

func FindMinDifference(timePoints []string) int {
	if len(timePoints) > 24*60 {
		return 0
	}
	mins := []int{}
	for _, t := range timePoints {
		time := strings.Split(t, ":")
		h, _ := strconv.Atoi(time[0])
		m, _ := strconv.Atoi(time[1])
		mins = append(mins, h*60+m)
	}
	sort.Ints(mins)
	mins = append(mins, mins[0]+24*60)
	res := 24 * 60
	for i := 1; i < len(mins); i++ {
		res = int(math.Min(float64(res), float64(mins[i]-mins[i-1])))
	}
	return res
}
