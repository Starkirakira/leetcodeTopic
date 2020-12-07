package lc

func MinDays(bloomDay []int,m,k int) int {
	if m*k > len(bloomDay){
		return -1
	}
	r := 0
	for _,num := range bloomDay {
		r =	maxF(r , num)
	}
	l := 0

	for l<r {
		mid := l + (r - l ) / 2
		count:= check(bloomDay, m, k, mid)
		if count {
			r = mid
		}else{
			l = mid + 1
		}
	}
	return l
}

func check(fl []int, m,k,mid int) bool {
	n := len(fl)
	cnt := 0
	lenF := 0
	for i:=0; i<n; i++{
		if fl[i] <= mid {
			lenF++
			if lenF == k {
				cnt++
				lenF-=k
			}
		}else{
			lenF = 0
		}
	}
	if cnt>=m {
		return true
	}else{
		return false
	}

}

func maxF(a int, b int) int {
	if a>b{
		return a
	}else{
		return b
	}
}