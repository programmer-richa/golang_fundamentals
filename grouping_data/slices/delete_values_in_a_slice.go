/*
Package main demonstrates the working of append function to delete values in a slice.
*/

package main

import (
	"fmt"
)

// Entry point of the program
func main() {
	// slice variable to store 5 values
	var s []float64 = []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Printf("Slice : %v\n", s) // Prints [1, 2, 3, 4, 5,6,7,8,9,10}]
	s = append(s[:5], s[6:]...)
	fmt.Printf("Slice value after deleting 6 from the slice : %v\n", s) // Prints [1,2,3,4,5,4,5,7,8,9]
}
