package lc

func flattenList(root *ListDoublePathList) *ListDoublePathList {
	curChild := flattenList(root.Child)
	curNext := flattenList(root.Next)
	if curChild != nil {
		root.Next = curChild
		curChild.Prev = root
	}
	tmp := curChild

	for tmp != nil && tmp.Next != nil {
		tmp = tmp.Next
	}
	if tmp != nil && curNext != nil {
		tmp.Next = curNext
		curNext.Prev = tmp
	}
	root.Child = nil
	return root

}
