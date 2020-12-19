package lc

type NodeRandomList struct {
	Val int
	Next *NodeRandomList
	Random *NodeRandomList

}

func CopyRandomList(head *NodeRandomList) *NodeRandomList {
	if head == nil {
		return nil
	}

	newHead := NodeRandomList{
		Val: head.Val,
		Next: nil,
		Random: nil,
	}
	p := head.Next

	pre := &newHead
	connection := make(map[*NodeRandomList]*NodeRandomList)
	connection[head] = pre

	for p != nil {
		newNode := &NodeRandomList{
			Val: p.Val,
			Next: nil,
			Random: nil,
		}
		pre.Next = newNode
		connection[p] = newNode
		p = p.Next
		pre = pre.Next
	}

	p = head
	newP := &newHead
	for p != nil {
		if p.Random != nil {
			newP.Random = connection[p.Random]
		}
		p = p.Next
		newP = newP.Next
	}
	return &newHead

}
