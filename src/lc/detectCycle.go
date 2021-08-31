package lc

//判断是否有环，若有则快慢指针再次出发得到环入口,超时
//func DetectCycle(head *ListNode) *ListNode {
//	inLoop := getNodeInLoop(head)
//	if inLoop == nil {
//		return nil
//	}
//	slow := head
//	fast := inLoop
//	for fast != slow {
//		slow = slow.Next
//		fast = fast.Next
//	}
//	return slow
//}
//
//func getNodeInLoop(head *ListNode) *ListNode {
//	if head == nil || head.Next == nil {
//		return nil
//	}
//
//	slow := head.Next
//	fast := head.Next
//	for slow != nil && fast != nil  {
//		if slow == fast{
//			return slow
//		}
//		if fast.Next == nil {
//			return nil
//		}
//		slow = slow.Next
//		fast = fast.Next.Next
//	}
//	return nil
//}

func DetectCycle(head *ListNode) *ListNode {
	slow := head
	fast := head
	f := false
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if fast == slow {
			f = true
			break
		}
	}
	if !f {
		return nil
	}
	slow = head
	for slow != fast {
		slow = slow.Next
		fast = fast.Next
	}
	return slow
}
