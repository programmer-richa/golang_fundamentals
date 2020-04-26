/*
Package main demonstrates slice creation from an array.
*/

package main

import (
	"fmt"
)

// Entry point of the program
func main() {
	// array variable to store 5 values
	var arr [5]float64 = [5]float64{1, 2, 3, 4, 5}
	fmt.Printf("Array value : %v\n", arr) // Prints [1,2,3,4,5]
	sliceAllValues := arr[:]
	fmt.Printf("Slice with all value : %v\n", sliceAllValues) // Prints [1,2,3,4,5]
	sliceFirst3Values := arr[:3]
	fmt.Printf("Slice with first 3 value : %v\n", sliceFirst3Values) // Prints [1,2,3]
	sliceSelectedValues := arr[1:4]
	fmt.Printf("Slice with value from index 1 to 3 value : %v\n", sliceSelectedValues) // Prints [2,3,4]
}
