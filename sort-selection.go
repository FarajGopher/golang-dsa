package main

import "fmt"

//time complexity ***** best case O(n^2) ***** worst case O(n^2)
//space complexity O(1)

//best to use when array size is small
//unstable algorithm :- does not preserve the order of same element

func selectionSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		minIndex := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		arr[minIndex], arr[i] = arr[i], arr[minIndex]
	}
}

func main() {
	arr := []int{6, 1, 5, 7, 8, 2, 3, 7, 9}
	selectionSort(arr)
	fmt.Println("after selection sort array : ", arr)
}
