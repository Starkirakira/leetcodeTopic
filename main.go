package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Codec struct {
	l []string
}

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func main()  {
	//a := []string{"1","2","3","null","null","4","5"}
	b := Constructor()

	res := b.deserialize("1,2,3,null,null,4,null,null,5,null,null")
	fmt.Printf("%+v",res)
}

func Constructor() Codec {
	return Codec{}
}

func (this *Codec) serialize(root *TreeNode) string  {
	return rserialize(root,"")
}

func rserialize(root *TreeNode, str string) string {
	if root == nil{
		str += "null,"
	}else{
		str += strconv.Itoa(root.Val) + ","
		str = rserialize(root.Left, str)
		str = rserialize(root.Right, str)
	}
	return str
}

func (this *Codec) deserialize(data string)	*TreeNode  {
	l := strings.Split(data, ",")
	for i :=0; i < len(l); i++ {
		if l[i] != ""{
			this.l = append(this.l , l[i])
		}
	}
	//fmt.Println(this.l)
	return this.rdeserialize()
}

func (this *Codec) rdeserialize() *TreeNode {
	if this.l[0] == "null"{
		this.l = this.l[1:]
		return nil
	}

	val,_ := strconv.Atoi(this.l[0])
	root := &TreeNode{Val: val}
	this.l = this.l[1:]
	root.Left = this.rdeserialize()
	root.Right = this.rdeserialize()
	return root
}

