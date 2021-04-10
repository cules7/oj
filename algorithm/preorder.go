package main

import (
	"fmt"
)

type Node struct {
	data int
	l    *Node
	r    *Node
}
type Stack struct {
}

func (stack *Stack) Push(n *Node) {

}
func (stack *Stack) Pop() *Node {
	return nil
}
func preorder(root *Node) {
	var stack = new(Stack)
	var p = root
	for p != nil {
		fmt.Println(p.data)
		if p.l != nil {
			stack.Push(p)
			p = p.l
			continue
		}
		if p.r != nil {
			p = p.r
			continue
		}
		p = stack.Pop()
		if p == nil {
			return
		}
		p = p.r
	}
}
