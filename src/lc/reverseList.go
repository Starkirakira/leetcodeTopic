package lc

func ReverseList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	var left *ListNode
	var right *ListNode
	cur := head
	for cur != nil {
		right = cur.Next
		cur.Next = left
		left = cur
		cur = right
	}
	return left
}
