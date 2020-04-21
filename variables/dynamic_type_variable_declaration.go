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
	int1 := 1
	float1 := 1.1
	complex1 := complex(1, 1)
	string1 := "Richa"
	var age = 29 // type will be inferred
	bool1 := true
	// Calling Printf function available in fmt package for printing the output on console
	fmt.Printf("Value stored in variable int1 of type %T : %v\n ", int1, int1)
	fmt.Printf("Value stored in variable float1 of type %T : %v\n ", float1, float1)
	fmt.Printf("Value stored in variable complex1 of type %T : %v\n ", complex1, complex1)
	fmt.Printf("Value stored in variable string1 of type %T : %v\n ", string1, string1)
	fmt.Printf("Value stored in variable bool1 of type %T : %v\n ", bool1, bool1)
	fmt.Printf("Value stored in variable age of type%T : %v\n ", age, age)
}
