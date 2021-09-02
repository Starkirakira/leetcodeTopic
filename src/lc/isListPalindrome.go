package lc

func IsListPalindrome(head *ListNode) bool {
	if head == nil {
		return true
	}
	dummy := &ListNode{0, nil}
	dummy.Next = head
	fast := dummy
	slow := dummy
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	r1 := slow.Next
	slow.Next = nil
	r2 := ReverseList(r1)
	l2 := head
	for r2 != nil {
		if r2.Val != l2.Val {
			return false
		}
		r2 = r2.Next
		l2 = l2.Next
	}
	return true

}
