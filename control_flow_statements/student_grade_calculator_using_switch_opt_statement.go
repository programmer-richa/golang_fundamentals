/*
Package main calculates the grade of a student based on his marks obtained in 5 subjects.
*/

package main

import (
	"fmt"
)

// Entry point of the program
func main() {
	// Array to store marks in 5 subjects
	var marks [5]float64
	// Declaring total variable to store total marks obtained
	total := 0.0 // here 0.0 specifies the type of total variable as float
	for i := 0; i < 5; i++ {
		fmt.Printf("Enter marks in subject %d: ", i+1)
		fmt.Scan(&marks[i])
		total += marks[i]
	}

	fmt.Printf("Total marks: %f\n", total)

	switch percentage := (total / 500) * 100; {
	case percentage > 80 && percentage <= 100:
		fmt.Printf("Percentage: %f\n", percentage)
		fmt.Println("Grade : A\n")
	case percentage > 60 && percentage <= 80:
		fmt.Printf("Percentage: %f\n", percentage)
		fmt.Println("Grade : B\n")
	case percentage > 40 && percentage <= 60:
		fmt.Printf("Percentage: %f\n", percentage)
		fmt.Println("Grade : C\n")
	default:
		fmt.Printf("Percentage: %f\n", percentage)
		fmt.Printf("Work hard.")

	}

}
