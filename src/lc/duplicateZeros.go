package lc

func DuplicateZeros(arr []int)  {

	for i := 0; i < len(arr)-1; i++ {
		if arr[i] == 0{
			arr = append(arr[:i+1], arr[i:len(arr)-1]...)
			i++
		}
	}
}
