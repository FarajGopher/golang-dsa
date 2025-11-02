package main

import (
	"fmt"
)

// Node represents a single node in the binary tree
type Node struct {
	Data  int
	Left  *Node
	Right *Node
}

// BuildTree builds the binary tree recursively
func BuildTree() *Node {
	var data int
	fmt.Print("Enter the data: ")
	fmt.Scan(&data)

	if data == -1 {
		return nil
	}

	root := &Node{Data: data}

	fmt.Printf("Enter data for inserting in left of %d\n", data)
	root.Left = BuildTree()

	fmt.Printf("Enter data for inserting in right of %d\n", data)
	root.Right = BuildTree()

	return root
}

// LevelOrderTraversal prints the binary tree level by level
func LevelOrderTraversal(root *Node) {
	if root == nil{
		return 
	}

	q := []*Node{root}
	arr := []int{}
	for len(q) != 0 {
		t := q[0]
		arr = append(arr, t.Data)
		q = q[1:]

		if t.Left != nil{
			q = append(q,t.Left)
		}
		if t.Right != nil {
			q = append(q, t.Right)
		}
	}

	fmt.Println("arr: ",arr)
}

// Inorder traversal (Left, Root, Right)
func Inorder(root *Node) {
	if root == nil {
		return
	}
	Inorder(root.Left)
	fmt.Print(root.Data, " ")
	Inorder(root.Right)
}

// Preorder traversal (Root, Left, Right)
func Preorder(root *Node) {
	if root == nil {
		return
	}
	fmt.Print(root.Data, " ")
	Preorder(root.Left)
	Preorder(root.Right)
}

// Postorder traversal (Left, Right, Root)
func Postorder(root *Node) {
	if root == nil {
		return
	}
	Postorder(root.Left)
	Postorder(root.Right)
	fmt.Print(root.Data, " ")
}

func main() {
	root := BuildTree()
	fmt.Println("\nLevel Order Traversal:")
	LevelOrderTraversal(root)

	fmt.Println("\nInorder Traversal:")
	Inorder(root)

	fmt.Println("\n\nPreorder Traversal:")
	Preorder(root)

	fmt.Println("\n\nPostorder Traversal:")
	Postorder(root)
	fmt.Println()
}
