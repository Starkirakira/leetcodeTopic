package lc

func FindArray(pref []int) []int {
	//[5,2,0,3,1] [5,7,2,3,2]
	arr := make([]int, len(pref))
	arr[0] = pref[0]
	for i := 1; i < len(pref); i++ {
		arr[i] = pref[i-1] ^ pref[i]
	}
	return arr

}
