/*
Package main prints the table of a number until the multiplier entered by the user is reached.
*/
package main

import (
	"fmt"
)

// Entry point of the program
func main() {
	var number, multiplier, skip int
	fmt.Print("Enter number: ")
	fmt.Scan(&number)
	fmt.Print("Enter max multiplier: ")
	fmt.Scan(&multiplier)
	fmt.Print("Enter multiplier to be skipped: ")
	fmt.Scan(&skip)
	fmt.Printf("Table of %d\n", number)
	for i := 1; i <= multiplier; i++ {
		if i == skip {
			continue
		}
		fmt.Printf("%d * %d = %d\n", number, i, number*i)

	}
}
