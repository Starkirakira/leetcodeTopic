package lc

func InsertListNode(head *ListNode, x int) *ListNode {
	node := &ListNode{x, nil}
	if head == nil {
		head = node
		node.Next = head
	} else if head.Next == head {
		head.Next = node
		node.Next = head
	} else {
		cur := head
		next := head.Next
		biggest := head
		for !(cur.Val <= node.Val && next.Val >= node.Val) && next != head {
			cur = next
			next = next.Next
			if cur.Val >= biggest.Val {
				biggest = cur
			}
		}
		if cur.Val <= node.Val && next.Val >= node.Val {
			cur.Next = node
			node.Next = next
		} else {
			node.Next = biggest.Next
			biggest.Next = node
		}
	}
	return head
}
