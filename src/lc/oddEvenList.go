package lc

func OddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	newHead := head.Next
	odd := head
	even := newHead

	for even != nil && even.Next != nil {
		odd.Next = even.Next
		odd = odd.Next
		even.Next = odd.Next
		even = even.Next
	}
	odd.Next = newHead
	return head
}
