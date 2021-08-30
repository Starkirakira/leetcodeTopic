package lc

func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{-1, nil}
	dummy.Next = head

	fast := dummy
	slow := dummy

	for n > 0 && fast.Next != nil {
		fast = fast.Next
		n--
	}

	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}
	slow.Next = slow.Next.Next
	return dummy.Next
}
