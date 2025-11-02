package main

import (
	"container/list"
	"fmt"
)

type MatrixBFS struct {
	Grid [][]int
}

func NewMatrixBFS() *MatrixBFS {
	return &MatrixBFS{
		Grid: [][]int{
			{0, 0, 0, 0},
			{1, 1, 0, 0},
			{0, 0, 0, 1},
			{0, 1, 0, 0},
		},
	}
}

func (m *MatrixBFS) BFS(grid [][]int) int {
	ROWS := len(grid)
	COLS := len(grid[0])
	visit := make([][]int, 4)
	for i := range visit {
		visit[i] = make([]int, 4)
	}
	queue := list.New()

	queue.PushBack([]int{0, 0}) // Add {0, 0}
	visit[0][0] = 1

	length := 0
	for queue.Len() > 0 {
		queueLength := queue.Len()
		for i := 0; i < queueLength; i++ {
			pair := queue.Remove(queue.Front()).([]int)
			r, c := pair[0], pair[1]
			if r == ROWS-1 && c == COLS-1 {
				return length
			}
			// We can directly build the four neighbors
			neighbors := [][]int{{r, c + 1}, {r, c - 1}, {r + 1, c}, {r - 1, c}}
			for j := 0; j < 4; j++ {
				newR, newC := neighbors[j][0], neighbors[j][1]
				if min(newR, newC) < 0 || newR == ROWS || newC == COLS || visit[newR][newC] == 1 || grid[newR][newC] == 1 {
					continue
				}
				queue.PushBack(neighbors[j])
				visit[newR][newC] = 1
			}
		}
		length++
	}
	return length // This should never be called
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func main() {
	m := NewMatrixBFS()
	fmt.Println(m.BFS(m.Grid))
}
