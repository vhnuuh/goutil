//http://c.biancheng.net/view/5568.html
package main

import "fmt"

type Node struct {
	data int
	next *Node
}

func showNode(p *Node) {
	for p != nil {
		fmt.Println(*p)
		p = p.next
	}
}

func main() {
	head := new(Node)
	head.data = 0
	tail := head
	for i := 1; i < 10; i++ {
		node := Node{data: i}
		tail.next = &node
		tail = &node
	}

	showNode(head)
}
