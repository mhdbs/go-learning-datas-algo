package main

import "fmt"

type Tree struct {
	node *Node
}

type Node struct {
	value int
	left  *Node
	right *Node
}

func (t *Tree) insert(value int) *Tree {
	if t.node == nil {
		t.node = &Node{value: value}
	} else {
		t.node.insert(value)
	}
	return t
}

func (n *Node) exists(value int) bool {
	if n == nil {
		return false
	}
	if n.value == value {
		return true
	}
	println("Value: ", n.value)
	if value <= n.value {
		return n.left.exists(value)
	} else {
		return n.right.exists(value)
	}
}

func printNode(n *Node) {
	if n == nil {
		return
	}
	fmt.Println(n.value)
	printNode(n.left)
	printNode(n.right)
}
func (n *Node) insert(value int) {
	if value <= n.value {
		if n.left == nil {
			n.left = &Node{value: value}
		} else {
			n.left.insert(value)
		}
	} else {
		if n.right == nil {
			n.right = &Node{value: value}
		} else {
			n.right.insert(value)
		}
	}
}

func main() {
	t := &Tree{}
	t.insert(10).
		insert(8).
		insert(20).
		insert(9).
		insert(0).
		insert(15).
		insert(25)
	printNode(t.node)

	println(t.node.exists(25))
}
