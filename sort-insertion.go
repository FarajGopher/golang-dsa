package main

import "fmt"

func insertionSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		temp := arr[i]
		j := i - 1
		for ; j >= 0; j-- {
			if arr[j] > temp {
				arr[j+1] = arr[j]
			} else {
				break
			}
		}
		arr[j+1] = temp
	}
}

func main() {
	arr := []int{6, 5, 4, 3, 2}
	insertionSort(arr)
	fmt.Println("array after insertion sort : ",arr)
}
