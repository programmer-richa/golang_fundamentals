/*
Package main prints the marks obtained by 2 students in 5 subjects.
*/

package main

import (
	"fmt"
)

// Entry point of the program
func main() {
	// map to store marks in 5 subjects
	// shorthand declaration
	student1 := map[string]float64{"Maths": 41, "English": 32, "Science": 34, "Computers": 44, "SST": 50}
	student2 := map[string]float64{"Maths": 34, "English": 35, "Science": 43, "Computers": 24, "SST": 48}
	// Declaring nested map
	marks := map[string]map[string]float64{}
	// adding values to the map
	marks["Student 1"] = student1
	marks["Student 2"] = student2

	for i, student := range marks {
		fmt.Printf("Marks of %s:", i)
		total := 0.0
		for j, v := range student {
			fmt.Printf("Marks obtained in %s is %.2f\n", j, v)
			total += v
		}
		fmt.Printf("Total Marks: %.2f\n\n", total)
	}
}
