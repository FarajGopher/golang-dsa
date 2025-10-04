package main

import "fmt"

// QuickSort Algorithm
// -------------------------------------------
// Approach:
// We use the "divide and conquer" strategy.
// In each partition step, we choose a pivot element
// and rearrange the array such that all elements
// smaller than the pivot go to its left
// and all elements greater than the pivot go to its right.
//
// Time Complexity:
// -> Best Case:     O(n log n)     (balanced partitions)
// -> Average Case:  O(n log n)
// -> Worst Case:    O(n^2)         (when pivot always chosen as smallest/largest element)
//
// Space Complexity: O(1) — In-place sorting (ignoring recursion stack)
//
// Stability: Not stable (relative order of duplicates not guaranteed)
//
// It is an in-place sorting algorithm.


func partition(arr []int, s, e int) int {
	pivot := arr[s]
	count := 0

	// Count elements smaller than or equal to pivot
	for i := s + 1; i <= e; i++ {
		if arr[i] <= pivot {
			count++
		}
	}

	pivotIndex := s + count
	arr[s], arr[pivotIndex] = arr[pivotIndex], arr[s]

	i, j := s, e
	for i < pivotIndex && j > pivotIndex {
		for arr[i] <= pivot && i < pivotIndex {
			i++
		}
		for arr[j] > pivot && j > pivotIndex {
			j--
		}

		if i < pivotIndex && j > pivotIndex {
			arr[i], arr[j] = arr[j], arr[i]
			i++
			j--
		}
	}

	return pivotIndex
}

func quickSort(arr []int, s, e int) {
	if s >= e {
		return
	}
	p := partition(arr, s, e)
	quickSort(arr, s, p-1)
	quickSort(arr, p+1, e)
}

func main() {
	arr := []int{6, 5, 4, 3, 2}
	quickSort(arr, 0, len(arr)-1)
	fmt.Println("Sorted:", arr)
}
