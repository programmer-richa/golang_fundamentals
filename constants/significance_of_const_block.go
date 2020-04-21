// Package main specifies that the variables, constants and functions
// declared in this folder belongs to the main namespace
package main

import "fmt"

// Declaring constants using a const block that automatically assigns value of constant declared above to the uninitialized constant in the block.
// A constant block groups the constants.
const (
	// It is mandatory to assign a value to the first declared constant in a block.
	first = 1
	// value of 'first' is assigned to 'repeatFirst'
	repeatFirst
	second = first * 2
	// value of 'second' is assigned to 'repeatSecond'
	repeatSecond
)

// Entry point of a program
func main() {
	fmt.Printf("Type and value of package level untyped constant 'first' is '%T' and '%v' respectively.\n", first, first)
	fmt.Printf("Type and value of package level untyped constant 'repeatFirst' is '%T' and '%v' respectively.\n", repeatFirst, repeatFirst)
	fmt.Printf("Type and value of package level untyped constant 'second' is '%T' and '%v' respectively.\n", second, second)
	fmt.Printf("Type and value of package level untyped constant 'repeatSecond' is '%T' and '%v' respectively.\n", repeatSecond, repeatSecond)

}
