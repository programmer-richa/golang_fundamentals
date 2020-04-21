// Package main specifies that the variables, constants and functions
// declared in this folder belongs to the main namespace
package main

// Importing required modules
import "fmt"

// Entry point of a program
func main() {
	// Declaring variable of int type that stores value 10 and is identified by identifier a
	var a int = 10
	// Calling Println function available in fmt package for printing the output on console
	fmt.Println("Value stored in variable a : ", a)
}
