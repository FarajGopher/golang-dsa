package main

import "fmt"

// simple implementation

// var stack []int

// func Push(v int) {
// 	stack = append(stack, v)
// }

// func Pop() {
// 	stack = stack[:len(stack)-1]
// }


// using structure implementation

type Stack struct {
	item []int
}

//append single element in a stack
func (s *Stack)Push(v int){
	s.item = append(s.item,v)
}

//remove last element form the stack
func (s *Stack)Pop(){
	s.item = s.item[:len(s.item)-1]
}

//Remove first element from the stack
func (s *Stack)RemoveFirstElem(){
	s.item = s.item[1:]
}

//Remove Middle Element from the stack
func (s *Stack)RemoveMiddleElem(){
	s.item = append(s.item[:len(s.item)/2],s.item[len(s.item)/2+1:]...)
}

//Remove element at specific index from the stack
func (s *Stack)RamoveAt(i int){
	s.item = append(s.item[:i],s.item[i+1:]... )
}

//Remove element from RemoveValue from the stack
func  (s *Stack)RemoveValue(val int){
	for i,v := range s.item {
		if v == val {
			s.RamoveAt(i)
		}
	}
}

func (s *Stack)InsertAtBottom(n int){
	if len(s.item) == 0 {
		s.item = append(s.item, n)
		return
	}

	top := s.item[len(s.item)-1]//top element
	s.Pop()

	//recursive call
	s.InsertAtBottom(n)
	s.Push(top)
}


//Reverse the stack
func (s *Stack)Reverse(){
	if len(s.item) == 0 {
		return
	}

	top := s.item[len(s.item)-1]//take top element
	s.Pop() //equivalent to stack.pop()
	s.Reverse() 
	s.InsertAtBottom(top)
}

func (s *Stack)SortedInserted(n int) {
	if len(s.item)== 0 || n > s.item[len(s.item)-1] {
		s.Push(n)
		return
	}

	top := s.item[len(s.item)-1]
	s.Pop()
	s.SortedInserted(n)
	s.Push(top)
}

//sort a stack
func (s *Stack)Sort(){
	if len(s.item) == 0{
		return 
	}
	top := s.item[len(s.item)-1]
	s.Pop()
	s.Sort()
	s.SortedInserted(top)
}

func main() {
	stack := Stack{
		item: nil,
	}
	stack.Push(5)
	stack.Push(3)
	stack.Push(10)
	stack.Push(4)
	stack.Push(8)
	stack.Push(6)
	stack.Push(8)
	stack.Push(7)
	stack.Push(2)
	stack.Push(1)
	stack.Pop()
	stack.RemoveFirstElem()
	stack.RemoveMiddleElem()
	stack.RamoveAt(2)
	stack.RemoveValue(6)
	stack.Reverse()
	stack.Sort()
	fmt.Println("stack: ", stack.item)
}
