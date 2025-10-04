package main

import "fmt"

type Deque struct {
	arr        []int
	front, rear int
	size        int
}

// Constructor
func NewDeque(n int) *Deque {
	return &Deque{
		arr:   make([]int, n),
		front: -1,
		rear:  -1,
		size:  n,
	}
}

// Pushes 'x' at the front
func (d *Deque) PushFront(x int) bool {
	if d.IsFull() {
		return false
	} else if d.IsEmpty() {
		d.front, d.rear = 0, 0
	} else if d.front == 0 && d.rear != d.size-1 {
		d.front = d.size - 1
	} else {
		d.front--
	}
	d.arr[d.front] = x
	return true
}

// Pushes 'x' at the rear
func (d *Deque) PushRear(x int) bool {
	if d.IsFull() {
		return false
	} else if d.IsEmpty() {
		d.front, d.rear = 0, 0
	} else if d.rear == d.size-1 && d.front != 0 {
		d.rear = 0
	} else {
		d.rear++
	}
	d.arr[d.rear] = x
	return true
}

// Pops from the front
func (d *Deque) PopFront() int {
	if d.IsEmpty() {
		return -1
	}

	ans := d.arr[d.front]
	d.arr[d.front] = -1

	if d.front == d.rear { // only one element
		d.front, d.rear = -1, -1
	} else if d.front == d.size-1 {
		d.front = 0 // cyclic nature
	} else {
		d.front++
	}
	return ans
}

// Pops from the rear
func (d *Deque) PopRear() int {
	if d.IsEmpty() {
		return -1
	}

	ans := d.arr[d.rear]
	d.arr[d.rear] = -1

	if d.front == d.rear { // single element
		d.front, d.rear = -1, -1
	} else if d.rear == 0 {
		d.rear = d.size - 1 // cyclic nature
	} else {
		d.rear--
	}
	return ans
}

// Returns front element
func (d *Deque) GetFront() int {
	if d.IsEmpty() {
		return -1
	}
	return d.arr[d.front]
}

// Returns rear element
func (d *Deque) GetRear() int {
	if d.IsEmpty() {
		return -1
	}
	return d.arr[d.rear]
}

// Check if empty
func (d *Deque) IsEmpty() bool {
	return d.front == -1
}

// Check if full
func (d *Deque) IsFull() bool {
	return (d.front == 0 && d.rear == d.size-1) ||
		(d.rear == (d.front-1+d.size)%d.size)
}

// Demo
func main() {
	d := NewDeque(5)

	fmt.Println(d.PushRear(10))  // true
	fmt.Println(d.PushRear(20))  // true
	fmt.Println(d.PushFront(5))  // true
	fmt.Println(d.PushFront(1))  // true
	fmt.Println(d.GetFront())    // 1
	fmt.Println(d.GetRear())     // 20
	fmt.Println(d.PopRear())     // 20
	fmt.Println(d.PopFront())    // 1
	fmt.Println(d.GetFront())    // 5
	fmt.Println(d.GetRear())     // 10
}
