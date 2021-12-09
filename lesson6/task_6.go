package main

import (
	"fmt"
)

type LinkNode struct {
	next  *LinkNode
	value int
}

func (n *LinkNode) Reverse() *LinkNode {
	var cur = n
	var prev *LinkNode
	for cur != nil {
		next := cur.next
		cur.next = prev
		prev = cur
		cur = next

	}
	return prev
}

func (n *LinkNode) Print() {
	cur := n
	for cur != nil {
		splt := " - > "
		fmt.Printf("%s%d", splt, cur.value)
		cur = cur.next
	}
}

func main() {
	n1 := &LinkNode{value: 1}
	n2 := &LinkNode{value: 2}
	n1.next = n2
	n3 := &LinkNode{value: 3}
	n2.next = n3
	n1.Reverse()
	n1.Print()

}
