package main

import (
	"fmt"
	"strings"
)

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

type BST struct {
	Root *Node
}

// Insert adds a value to the BST
func (b *BST) Insert(value int) {
	newNode := &Node{Value: value}
	if b.Root == nil {
		b.Root = newNode
		return
	}

	current := b.Root
	for {
		if value == current.Value {
			return // Duplicate, ignore
		}
		if value < current.Value {
			if current.Left == nil {
				current.Left = newNode
				return
			}
			current = current.Left
		} else {
			if current.Right == nil {
				current.Right = newNode
				return
			}
			current = current.Right
		}
	}
}

// Include checks if a value exists in the BST
func (b *BST) Include(value int) bool {
	current := b.Root
	for current != nil {
		if value < current.Value {
			current = current.Left
		} else if value > current.Value {
			current = current.Right
		} else {
			return true
		}
	}
	return false
}

// BFS returns level-order traversal
func (b *BST) BFS() []int {
	if b.Root == nil {
		return nil
	}
	queue := []*Node{b.Root}
	result := []int{}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		result = append(result, current.Value)

		if current.Left != nil {
			queue = append(queue, current.Left)
		}
		if current.Right != nil {
			queue = append(queue, current.Right)
		}
	}
	return result
}

// DFS Preorder (Root -> Left -> Right)
func (b *BST) DFSPreorder(node *Node, result *[]int) {
	if node == nil {
		return
	}
	*result = append(*result, node.Value)
	b.DFSPreorder(node.Left, result)
	b.DFSPreorder(node.Right, result)
}

// DFS Inorder (Left -> Root -> Right)
func (b *BST) DFSInorder(node *Node, result *[]int) {
	if node == nil {
		return
	}
	b.DFSInorder(node.Left, result)
	*result = append(*result, node.Value)
	b.DFSInorder(node.Right, result)
}

// DFS Postorder (Left -> Right -> Root)
func (b *BST) DFSPostorder(node *Node, result *[]int) {
	if node == nil {
		return
	}
	b.DFSPostorder(node.Left, result)
	b.DFSPostorder(node.Right, result)
	*result = append(*result, node.Value)
}

// FindMin returns smallest value
func (b *BST) FindMin() *int {
	if b.Root == nil {
		return nil
	}
	current := b.Root
	for current.Left != nil {
		current = current.Left
	}
	return &current.Value
}

// FindMax returns largest value
func (b *BST) FindMax() *int {
	if b.Root == nil {
		return nil
	}
	current := b.Root
	for current.Right != nil {
		current = current.Right
	}
	return &current.Value
}

// Height returns height of tree
func (b *BST) Height(node *Node) int {
	if node == nil {
		return 0
	}
	left := b.Height(node.Left)
	right := b.Height(node.Right)
	if left > right {
		return left + 1
	}
	return right + 1
}

// Pretty Print Tree (visual representation)
func (b *BST) Print(node *Node, prefix string, isLeft bool) {
	if node == nil {
		return
	}
	if node.Right != nil {
		b.Print(node.Right, prefix+func() string {
			if isLeft {
				return "│   "
			}
			return "    "
		}(), false)
	}
	fmt.Println(prefix + func() string {
		if isLeft {
			return "└── "
		}
		return "┌── "
	}() + fmt.Sprintf("%d", node.Value))
	if node.Left != nil {
		b.Print(node.Left, prefix+func() string {
			if isLeft {
				return "    "
			}
			return "│   "
		}(), true)
	}
}

func main() {
	bst := &BST{}
	bst.Insert(7)
	bst.Insert(2)
	bst.Insert(3)
	bst.Insert(1)
	bst.Insert(9)
	bst.Insert(4)

	fmt.Println("Include 3:", bst.Include(3))
	fmt.Println("BFS:", bst.BFS())

	pre := []int{}
	bst.DFSPreorder(bst.Root, &pre)
	fmt.Println("DFS Preorder:", pre)

	in := []int{}
	bst.DFSInorder(bst.Root, &in)
	fmt.Println("DFS Inorder:", in)

	post := []int{}
	bst.DFSPostorder(bst.Root, &post)
	fmt.Println("DFS Postorder:", post)

	fmt.Println("Min:", *bst.FindMin())
	fmt.Println("Max:", *bst.FindMax())
	fmt.Println("Height:", bst.Height(bst.Root))

	fmt.Println("\nVisual Tree:")
	bst.Print(bst.Root, "", true)
}
