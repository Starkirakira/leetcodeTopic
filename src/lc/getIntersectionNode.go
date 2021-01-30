package lc

func GetIntersectionNode(headA, headB *ListNode) *ListNode {
	t1 := headA
	t2 := headB
	for t1 != t2 {
		if t1 == nil {
			t1 = headB
		} else {
			t1 = t1.Next
		}

		if t2 == nil {
			t2 = headA
		} else {
			t2 = t2.Next
		}
	}
	return t1
}
