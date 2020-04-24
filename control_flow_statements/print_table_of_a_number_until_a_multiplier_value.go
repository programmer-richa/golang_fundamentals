/*
Package main prints the table of a number until the multiplier entered by the user is reached.
*/
package main

import (
	"fmt"
)

// Entry point of the program
func main() {
	var number, multiplier int
	fmt.Print("Enter number: ")
	fmt.Scan(&number)
	fmt.Print("Enter max multiplier: ")
	fmt.Scan(&multiplier)
	fmt.Printf("Table of %d\n", number)
	for i := 1; ; i++ {
		fmt.Printf("%d * %d = %d\n", number, i, number*i)
		if i == multiplier {
			break
		}
	}
}
