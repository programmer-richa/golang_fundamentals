/*
Package main prints the marks obtained by a student in 5 subjects.
*/

package main

import (
	"fmt"
)

// Entry point of the program
func main() {
	// map to store marks in 5 subjects
	// shorthand declaration
	marks := map[string]float64{"Maths": 41, "English": 32, "Science": 34, "Computers": 44, "SST": 50}
	total := 0.0
	for i, v := range marks {
		fmt.Printf("Marks obtained in %s is %.2f\n", i, v)
		total += v
	}
	fmt.Printf("Total Marks : %.2f", total)
}
