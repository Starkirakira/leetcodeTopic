package lc

import (
	"strconv"
	"strings"
)



type Codecii struct {

}

func Constructorii() Codecii {
	return Codecii{}
}

// Serializes a tree to a single string.
func (this *Codecii) Serializeii(root *TreeNodeii) string {
	if root == nil {
		return ""
	}
	var res []string
	var queue = []*TreeNodeii{root}
	for 0 < len(queue) {
		if queue[0] != nil {
			res = append(res, strconv.Itoa(queue[0].Val))
			queue = append(queue, (queue[0].Left), (queue[0].Right))
		} else {
			res = append(res, "#")
		}
		queue = queue[1:]
	}

	return strings.Join(res, ",")

}

// Deserializes your encoded data to tree.
func (this *Codecii) Deserializeii(data string) *TreeNodeii {
	if len(data) == 0 {
		return nil
	}

	res := strings.Split(data, ",")
	root := &TreeNodeii{}
	root.Val,_ =strconv.Atoi(res[0])
	var queue = []*TreeNodeii{root}
	res = res[1:]
	for 0 < len(queue) {
		if res[0] != "#" {
			l, _ := strconv.Atoi(res[0])
			queue[0].Left = &TreeNodeii{Val: l}
			queue = append(queue, (queue[0].Left))
		}

		if res[1] != "#" {
			r, _ := strconv.Atoi(res[1])
			queue[0].Right = &TreeNodeii{Val: r}
			queue = append(queue, (queue[0].Right))
		}
		queue = queue[1:]
		res = res[2:]
	}
	return root


}


