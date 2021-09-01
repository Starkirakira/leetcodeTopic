package lc

func AddTwoNumbers(l1, l2 *ListNode) *ListNode {
	r1 := ReverseList(l1)
	r2 := ReverseList(l2)

	var nx, ny *ListNode
	carry := 0

	for r1 != nil || r2 != nil || carry > 0 {
		n1, n2 := 0, 0
		if r1 != nil {
			n1 = r1.Val
			r1 = r1.Next
		}
		if r2 != nil {
			n2 = r2.Val
			r2 = r2.Next
		}
		sum := n1 + n2 + carry
		ny = &ListNode{sum % 10, nil}
		ny.Next = nx
		nx = ny
		carry = sum / 10
	}
	return ny
}
