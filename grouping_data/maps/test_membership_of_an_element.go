/*
Package main tests membership of a key in map.
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
	// test if entry is present in the map or not
	maths, ok := marks["Maths"]
	if ok {
		fmt.Printf("Math key is present and its value is %v", maths)
	} else {
		fmt.Print("Math key is not present")
	}
}
