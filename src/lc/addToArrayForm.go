package lc

func AddToArrayForm(A []int, K int) []int {

	n := len(A)
	i := n - 1
	res := []int{}
	sum := 0
	carry := 0
	for i >=0 || K != 0 {
		x,y := 0, 0
		if i >= 0 {
			x = A[i]
		}
		if K != 0 {
			y = K % 10
		}
		sum = x + y + carry
		carry = sum / 10
		K = K / 10
		i--
		res = append(res,sum % 10)
	}
	if carry != 0 {
		res = append(res,carry)
	}
	reverse(res)
	return res
}

func reverse(A []int) {
	for i, n := 0, len(A); i < n/2; i++ {
		A[i], A[n-1-i] = A[n-1-i], A[i]
	}
}

