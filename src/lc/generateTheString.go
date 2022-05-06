package lc

func GenerateTheString(n int) string {
	var sRes string
	if n%2 == 0 {
		for i := 0; i < n-1; i++ {
			sRes += "a"
		}
		sRes += "b"
	} else {
		for i := 0; i < n; i++ {
			sRes += "a"
		}
	}
	return sRes
}
