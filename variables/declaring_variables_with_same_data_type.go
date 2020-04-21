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
	var a, b int = 10, 20
	sum := a + b
	// Calling Printf function available in fmt package for printing the output on console
	fmt.Printf("Sum of %v and %v is %v", a, b, sum)
}
