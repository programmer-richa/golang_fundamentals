/*
Package main tests whether the char entered by the user is a vowel.
*/
package main

import (
	"fmt"
)

// Entry point of the program
func main() {
	var char string
	fmt.Print("Enter a char: ")
	fmt.Scan(&char)
	switch char {
	case "a":
		fallthrough
	case "A":
		fallthrough
	case "e":
		fallthrough
	case "E":
		fallthrough
	case "i":
		fallthrough
	case "I":
		fallthrough
	case "o":
		fallthrough
	case "O":
		fallthrough
	case "u":
		fallthrough
	case "U":
		fmt.Printf("%s is a vowel", char)
	default:
		fmt.Printf("%s is not a vowel", char)
	}

}
