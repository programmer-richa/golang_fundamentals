// Package main specifies that the variables, constants and functions
// declared in this folder belongs to the main namespace
package main

// Importing required modules
import (
	"fmt"
)

// Entry point of a program
func main() {
	// Declaring variable without specifying data type
	name, age, height := "Ram", 20, 5.9
	// Calling Printf function available in fmt package for printing the output on console
	fmt.Printf("Name: %v \t Age: %v \t Height %v", name, age, height)
}
