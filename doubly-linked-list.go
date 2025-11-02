package main

import "fmt"

type Node struct {
	data int
	prev *Node
	next *Node
}

type linkedListNew struct {
	head *Node
}

func (l *linkedListNew) insertHead(v int) {
	newNode := &Node{data: v, prev: nil, next: nil}

	if l.head == nil {
		l.head = newNode
	} else {
		newNode.next = l.head
		l.head.prev = newNode
		l.head = newNode
	}
}


func (l *linkedListNew) PrintLinkedList() {
	if l.head == nil {
		fmt.Println("empty list")
		return
	}

	curr := l.head
	for curr != nil {
		fmt.Print(curr.data)
		if curr.next != nil {
			fmt.Print(" --> ")
		}
		curr = curr.next
	}
	fmt.Println()
}

func main() {
	l := linkedListNew{}
	l.insertHead(10)
	l.insertHead(20)
	l.PrintLinkedList()

}
