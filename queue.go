package main

import "fmt"

type Queue struct {
	arr   []int
	front int
	rear  int
	size  int
}

// Constructor
func NewQueue(size int) *Queue {
	return &Queue{
		size:  size,
		arr:   make([]int, size),
		front: 0,
		rear:  0,
	}
}

// Enqueue operation
func (q *Queue) Enqueue(x int) {
	if q.rear == q.size {
		fmt.Println("cann't insert queue is full")
		return
	}
	q.arr[q.rear] = x
	q.rear++
}

// Dequeue 
func (q *Queue) Dequeue() int{
	if q.front == q.rear {
		fmt.Println("Queue is empty cannot dequeue")
		return -1
	}
	val:= q.arr[q.front]
	q.front++
	if q.front == q.rear{
		q.front = 0
		q.rear = 0
	}
	return val
}


//peek
func (q *Queue) Peek() int{
	if q.front == q.rear {
		fmt.Println("Queue is empty")
		return -1
	}
	return q.arr[q.front]
}

//check is empty
func (q *Queue) IsEmpty() bool{
	return q.front == q.rear 
}

func (q *Queue) PrintQueueElements() []int{
	return q.arr[:q.rear]
}

func main() {
	//1st method to initialize a queue
	// size := 10000
	// q := Queue{
	// 	arr: make([]int, size),
	// 	front: 0,
	// 	rear: 0,
	// 	size: size,
	// }

	//2nd method to initialize a queue using constructor function
	q := NewQueue(10000)
	q.Enqueue(10)
	q.Enqueue(20)
	q.Enqueue(30)

	fmt.Println("queue elements:",q.PrintQueueElements())
	fmt.Println("Front element:", q.Peek()) 
	fmt.Println("Dequeued:", q.Dequeue())   
	fmt.Println("Dequeued:", q.Dequeue())  
	fmt.Println("Is Empty:", q.IsEmpty())  

}
