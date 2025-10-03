package main

import "fmt"

func merge(arr []int, s, e int) {
	m := s + (e-s)/2
	len1 := m - s + 1
	len2 := e - m

	// create slices
	first := make([]int, len1)
	second := make([]int, len2)

	// copy values
	for i := 0; i < len1; i++ {
		first[i] = arr[s+i]
	}
	for i := 0; i < len2; i++ {
		second[i] = arr[m+1+i]
	}

	// merge two sorted halves
	i, j, k := 0, 0, s
	for i < len1 && j < len2 {
		if first[i] <= second[j] {
			arr[k] = first[i]
			i++
		} else {
			arr[k] = second[j]
			j++
		}
		k++
	}
	
	// copy remaining
	for i < len1 {
		arr[k] = first[i]
		i++
		k++
	}
	for j < len2 {
		arr[k] = second[j]
		j++
		k++
	}
}

func mergeSort(arr []int, s, e int) {
	if s >= e {
		return
	}
	mid := s + (e-s)/2
	mergeSort(arr, s, mid)   // left half
	mergeSort(arr, mid+1, e) // right half
	merge(arr, s, e)
}

func main() {
	arr := []int{6, 5, 4, 3, 2}
	mergeSort(arr, 0, len(arr)-1)
	fmt.Println("Sorted:", arr)
}
