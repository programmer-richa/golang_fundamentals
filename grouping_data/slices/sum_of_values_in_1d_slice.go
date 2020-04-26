/*
Package main prints the sum of marks obtained by a student in 5 subjects.
*/

package main

import (
	"fmt"
)

// Entry point of the program
func main() {
	// Slice to store marks in 5 subjects
	// here the size of the slice is determined based on the values passed
	// the following is an example of shorthand declaration
	marks := []float64{50, 70, 90, 100, 85}
	// variable to store total marks
	total := 0.0
	for i, v := range marks {
		fmt.Printf("Marks obtained in subject %d : %.2f\n", i, v)
		total += v
	}
	fmt.Printf("Total Marks : %.2f", total)
}
