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
	//[Computers:44 English:32 Maths:41 SST:50 Science:34 ]
	fmt.Printf("Map : %v\n", marks)
	total := 0.0
	for _, v := range marks {
		total += v

	}
	marks["Total"] = total
	//[Computers:44 English:32 Maths:41 SST:50 Science:34 Total:201]
	fmt.Printf("Map after adding total key : %v\n", marks)
}
