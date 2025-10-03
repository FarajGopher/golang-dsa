package main

import "fmt"

//we follow approach where in its every ith round the largest among all will be at their right position
// time complexity ****** best case using swaped var ******* O(n) ******* worst cast O(n^2)
//space complexity O(1)
//it is stable algorithm order maintained for duplicates
//inplace sorting algo


func bubbleSort(arr []int) {
	swapped := false
	for i := 1; i < len(arr); i++ {
		for j := 0; j < len(arr)-i-1; j++ { //(len(arr)-1 ------->>>>> len(arr)-i)
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				swapped = true
			}
		}
		if(!swapped){
			break
		}
	}
}

func main() {
	arr := []int{6, 1, 5, 7, 8, 2, 3, 7, 9}
	// arr := []int{1,2,3,4,5,6,7,8} no swap
	bubbleSort(arr)
	fmt.Println("array after bubble sort: ", arr)
}