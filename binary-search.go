package main

import "fmt"

//**** time complexity O(log n)


func main() {
	arr := []int{1, 2, 3, 4, 5}
	find := 4
	min := 0
	max := len(arr) - 1
	found := -1
	for min <= max{
		// mid := (min + max) / 2
		mid := min + ((max-min)/2)
		if arr[mid] == find {
			found = mid 
			break
		} else if arr[mid] < find {
			min = mid + 1
		} else {
			max = mid - 1
		}
	}
	if found > 0 {
	fmt.Println("found at index : ",found)
	}else{
		fmt.Println("not found")
	}
}