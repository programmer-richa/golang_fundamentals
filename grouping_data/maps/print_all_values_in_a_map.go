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
	var marks = map[string]float64{"Maths": 41, "English": 32, "Science": 34, "Computers": 44, "SST": 50}
	for i, v := range marks {
		fmt.Printf("Marks obtained in %s is %f\n", i, v)
	}
}
