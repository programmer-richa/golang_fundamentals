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
	// Declaring total variable to store total marks obtained and percentage of student
	total, percentage := 0.0, 0.0 // here 0.0 specifies the type of total variable as float
	for i := 0; i < 5; i++ {
		fmt.Printf("Enter marks in subject %d: ", i+1)
		fmt.Scan(&marks[i])
		total += marks[i]
	}
	percentage = (total / 500) * 100
	fmt.Printf("Total marks: %f\n", total)
	fmt.Printf("Percentage: %f\n", percentage)
	if percentage > 80 && percentage <= 100 {
		fmt.Println("Grade : A\n")
	} else if percentage > 60 && percentage <= 80 {
		fmt.Println("Grade : B\n")
	} else if percentage > 40 && percentage <= 60 {
		fmt.Println("Grade : C\n")
	} else {
		fmt.Printf("Work hard.")
	}
}
