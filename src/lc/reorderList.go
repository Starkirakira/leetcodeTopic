package lc

func ReorderList(head *ListNode) {
	dummy := &ListNode{0, nil}
	dummy.Next = head
	slow := dummy
	fast := dummy

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	headB := slow.Next
	slow.Next = nil
	p2 := ReverseList(headB)
	p1 := head
	p3 := &ListNode{}
	for p2 != nil {
		p3 = p1.Next
		p1.Next = p2
		p1 = p2
		p2 = p3
	}

}
