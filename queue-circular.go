package main

import "fmt"

type CircularQueue struct {
	front int
	rear  int
	size  int
	arr   []int
}

func (cq *CircularQueue) enqueue(item int) {
	if (cq.front == 0 && cq.rear == cq.size-1) || (cq.rear+1)%cq.size == cq.front {
		fmt.Println("Queue is full")
		return
	} else if cq.front == -1 { // first element to push
		cq.front ,cq.rear = 0,0
		cq.arr[cq.rear] = item
	} else if cq.rear == cq.size-1 && cq.front != 0 { // rear is at last and first is not at index 0 this means start of queue is free
		cq.rear = 0
		cq.arr[cq.rear] = item
	} else {						//else normal rear insertion
		cq.rear++ 						
		cq.arr[cq.rear] = item
	}
	return
}

func (cq *CircularQueue) dequeue() int{
	if cq.front == -1 {
		fmt.Println("Queue is empty")
		return -1
	}
	ans := cq.arr[cq.front]
	cq.arr[cq.front] = -1
	if cq.front == cq.rear{
		cq.front ,cq.rear = -1,-1
	}else if cq.front == cq.size-1{
		cq.front = 0
	}else{
		cq.front++
	}
	return ans
}

func main() {

	size := 10
	q := CircularQueue{
		size:  10,
		front: -1,
		rear:  -1,
		arr:   make([]int, size),
	}

	q.enqueue(1)
	q.enqueue(2)
	q.enqueue(3)
	q.enqueue(4)
	q.enqueue(5)
	q.enqueue(6)
	q.enqueue(7)
	q.enqueue(8)
	q.enqueue(9)
	q.enqueue(10)
	fmt.Println("queue after insertion: ", q.arr)

	q.dequeue()
	q.dequeue()
	q.dequeue()
	fmt.Println("queue after dequeue: ",q.arr)

	q.enqueue(11)
	q.enqueue(12)
	fmt.Println("queue after insertion: ", q.arr)
}