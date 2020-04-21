// Package main specifies that the variables, constants and functions
// declared in this folder belongs to the main namespace
package main

import "fmt"

// Importing required modules

// PI is untyped constants that are assigned a float value.
// Since a constant name starts with a capital letter, it is treated as an exported constant.
// An exported constant is a constant that is accessible from outside of the package.
// For an exported constant, it is recommended to write a comment starting with the identifier used for the constant, as done in comment line 1.
const PI = 3.14

// publish is a typed constant which stores value true
// Since the constant name starts with a lower letter, it is treated as a package level constant.
// A package level constant is not accessible from outside of the package
const publish bool = true

// Declaring constants using a const block that provides information regarding a project.
// A constant block groups the constants.
const (
	author = "Richa"
	title  = "Golang Fundamentals"
	task   = 1
)

// Entry point of a program
func main() {
	const local = "This is a local constant as it is declared inside the main() block."
	fmt.Printf("Unlike variables, it is optional to use all declared constants in the program.\n")
	fmt.Printf("Type and value of exported untyped constant 'PI' is '%T' and '%v' respectively.\n", PI, PI)
	fmt.Printf("Type and value of package level typed constant 'publish' is '%T' and '%v' respectively.\n", publish, publish)
	fmt.Printf("Type and value of local untyped constant 'local' is '%T' and '%v' respectively.\n", local, local)

}
