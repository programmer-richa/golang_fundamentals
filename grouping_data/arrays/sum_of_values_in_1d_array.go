/*
Package main prints the sum of marks obtained by a student in 5 subjects.
*/

package main

import (
	"fmt"
)

// Entry point of the program
func main() {
	// Array to store marks in 5 subjects
	// here ...(ellipsis) operator is used to determine the size of array based on the values passed
	// the following is an example of shorthand declaration
	marks := [...]float64{50, 70, 90, 100, 85}
	// variable to store total marks
	total := 0.0
	for i, v := range marks {
		fmt.Printf("Marks obtained in subject %d : %f\n", i, v)
		total += v
	}
	fmt.Printf("Total Marks : %f", total)
}
