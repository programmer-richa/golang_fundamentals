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
	case "a", "A", "e", "E", "i", "I", "o", "O", "u", "U":
		fmt.Printf("%s is a vowel", char)
	default:
		fmt.Printf("%s is not a vowel", char)
	}

}
