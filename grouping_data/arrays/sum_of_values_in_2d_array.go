/*
Package main prints the sum of marks obtained by 2 students in 5 subjects.
*/

package main

import (
	"fmt"
)

// Entry point of the program
func main() {
	// Array to store marks of 2 students in 5 subjects
	// the following is an example of shorthand declaration
	marks := [2][5]float64{
		{50, 70, 90, 100, 85},
		{55, 75, 79, 90, 80},
	}
	// variable to store total marks

	for i, student := range marks {
		total := 0.0
		fmt.Printf("Marks obtained by student %d : \n", i+1)
		for j, mark := range student {
			fmt.Printf("Subject %d : %.2f\n", j+1, mark)
			total += mark
		}
		fmt.Printf("Total Marks : %.2f\n\n", total)
	}

}
