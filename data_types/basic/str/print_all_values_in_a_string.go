/*
Package main prints the alphabets in a string.
*/

package main

import (
	"fmt"
)

// Entry point of the program
func main() {
	// String to store name
	var name = "Richa"
	for i, v := range name {
		fmt.Printf("value stored at index %d is rune having value %d and alphabet equivalence %v\n", i, v, string(v))
	}
}
