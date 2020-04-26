/*
Package main tests membership of a key in a map and delete it.
delete() function is used to delete an entry from a map. It requires the map and the corresponding key which is to be deleted.
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
	_, ok := marks["Maths"]
	if ok {
		delete(marks, "Maths")
		fmt.Printf("Map after deleting 'Maths' entry : %v\n", marks)
	} else {
		fmt.Print("Math key is not present")
	}
}
