package main

import (
	"fmt"
)

//bit manipulation have 2-3 things 
//1. & ,| ,^ ,not - ^before (^n), bit shifting << >>
//bit shifting left is also called multiple of 2 and right shift we are dividing it by 2
//log base 2 representation 2^0,2^1,2^3





func operations() {
	// AND
	n := 1 & 1

	// OR 
	n = 1 | 0

	// XOR
	n = 0 ^ 1

	// NOT (negation)
	n = ^n

	// Bit Shifting
	n = 1
	n = n << 1
	n = n >> 1
}

// Count the number of 1 bits in an int
func countBits(n int) int {
	count := 0
	for n > 0 {
		if (n & 1) == 1 {
			count++
		}
		n = n >> 1 // same as n / 2
	}
	return count
}

func main() {
	operations() // Perform the operations
	fmt.Println(countBits(10)) // For example, count the number of bits in 10
}
