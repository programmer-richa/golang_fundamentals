/*
Package main demonstrates the working of append function to add values in a slice.
*/

package main

import (
	"fmt"
)

// Entry point of the program
func main() {
	// slice variable to store 5 values
	var s []float64 = []float64{1, 2, 3, 4, 5}
	fmt.Printf("Slice value : %v\n", s) // Prints [1,2,3,4,5]
	s = append(s, 4, 5, 6)
	fmt.Printf("Slice value after appending '4,5,6' : %v\n", s) // Prints [1,2,3,4,5,4,5,6]
	arr := [...]float64{7, 8, 9}
	fmt.Printf("Array value : %v\n", arr) // Prints [7,8,9]
	s = append(s, arr[:]...)
	fmt.Printf("Slice value after appending all values of array : %v\n", s) // Prints [1,2,3,4,5,4,5,6,7,8,9]
}
