package main

import "fmt"

func dfs(grid [][]int, r, c int, visit map[[2]int]bool) int {
	rows, cols := len(grid), len(grid[0])

	// Base cases: out of bounds, obstacle, or already visited
	if r < 0 || c < 0 || r >= rows || c >= cols || visit[[2]int{r, c}] || grid[r][c] == 1 {
		return 0
	}

	// If reached bottom-right corner
	if r == rows-1 && c == cols-1 {
		return 1
	}

	// Mark as visited
	visit[[2]int{r, c}] = true

	// Explore all 4 directions
	count := 0
	count += dfs(grid, r+1, c, visit)
	count += dfs(grid, r-1, c, visit)
	count += dfs(grid, r, c+1, visit)
	count += dfs(grid, r, c-1, visit)
	fmt.Println("visit : ",visit)
	// Backtrack (unmark)
	delete(visit, [2]int{r, c})

	return count
}

func main() {
	grid := [][]int{
		{0, 0, 0, 0},
		{1, 1, 0, 0},
		{0, 0, 0, 1},
		{0, 1, 0, 0},
	}

	visit := make(map[[2]int]bool)
	fmt.Println(dfs(grid, 0, 0, visit)) // prints number of valid paths
}
