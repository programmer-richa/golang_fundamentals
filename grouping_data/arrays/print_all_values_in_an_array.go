/*
Package main prints the marks obtained by a student in 5 subjects.
*/

package main

import (
	"fmt"
)

// Entry point of the program
func main() {
	// Array to store marks in 5 subjects
	var marks [5]float64 = [5]float64{1, 2, 3, 4, 5}
	for i, v := range marks {
		fmt.Printf("value stored at index %d is %f\n", i, v)
	}
}
