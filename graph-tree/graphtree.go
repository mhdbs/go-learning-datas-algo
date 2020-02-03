package main

import "fmt"

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func main() {
	var N int
	fmt.Scanf("%d", &N)
	var nodes []Node = make([]Node, N)
	for i := 0; i < N; i++ {
		var val, indexLeft, indexRight int
		fmt.Scanf("%d %d %d", &val, &indexLeft, &indexRight)
		nodes[i].Value = val
		if indexLeft >= 0 {
			nodes[i].Left = &nodes[indexLeft]
		}
		if indexRight >= 0 {
			nodes[i].Right = &nodes[indexRight]
		}
	}
	for _, node := range nodes {
		NodeOut(&node)
	}
}

func NodeOut(n *Node) {
	fmt.Println("Node value: ", n.Value)
	if n.Left != nil {
		fmt.Println("Left: ", n.Left.Value)
	}
	if n.Right != nil {
		fmt.Println("Right: ", n.Right.Value)
	}
}
