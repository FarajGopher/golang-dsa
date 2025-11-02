package main

import "fmt"

type node struct {
	data int
	next *node
}

type linkedList struct {
	head *node
}

func (l *linkedList) Insert(value int) {
	newNode := &node{data: value, next: nil}
	if l.head == nil {
		l.head = newNode
		return
	} else {
		curr := l.head
		for curr.next != nil {
			curr = curr.next
		}
		curr.next = newNode
	}
	return
}

func (l *linkedList) PrintLinkedList() {
	if l.head == nil {
		fmt.Println("empty list")
		return
	}else{
		curr := l.head
		for curr != nil {
			fmt.Print(curr.data," --> ")
			curr = curr.next
		}
	}
	return
}

func main() {
	l := linkedList{}
	l.Insert(10)
	l.Insert(20)
	l.Insert(30)

	l.PrintLinkedList()
}