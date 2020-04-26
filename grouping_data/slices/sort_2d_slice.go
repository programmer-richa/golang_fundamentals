/*
Package main prints the sum of marks obtained by 5 students in 5 subjects.
*/

package main

import "fmt"

// Entry point of the program
func main() {

	// Array to store marks of 5 students in 5 subjects
	// the following is an example of shorthand declaration
	marks := [][]float64{
		{50, 70, 90, 100, 85, 0},
		{55, 75, 79, 90, 80, 0},
		{55, 75, 79, 90, 80, 0},
		{10, 20, 30, 40, 50, 0},
		{20, 30, 40, 50, 60, 0},
	}
	// loop to store total marks at last index
	students := len(marks)
	for i, student := range marks {
		subjects := len(marks)

		total := 0.0
		for j := 0; j < subjects; j++ {
			mark := student[j]
			total += mark
		}
		marks[i][subjects] = total
	}
	// sort array
	for i := 0; i < students-1; i++ {
		subjects := len(marks)
		for j := i + 1; j < students; j++ {
			if marks[i][subjects] < marks[j][subjects] {
				temp := marks[i]
				marks[i] = marks[j]
				marks[j] = temp
			}
		}
	}

	fmt.Print(marks)

}
